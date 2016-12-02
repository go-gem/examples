package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-gem/gem"
	"github.com/go-gem/gem/middleware"
	"github.com/valyala/fasthttp"
)

type UserClaims struct {
	jwt.StandardClaims
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	// users, key => value equal that name => id.
	users = map[string]int{
		"foo": 1,
		"bar": 2,
	}
)

// jwt configuration.
var (
	signingMethod = jwt.SigningMethodHS256
	signKey       = []byte("123456789")
	keyFunc       = func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	}

	jwtMiddleware = middleware.NewJWT(signingMethod, keyFunc)
)

func init() {
	jwtMiddleware.NewClaims = func() jwt.Claims {
		return &UserClaims{}
	}
}

func generateHandler(ctx *gem.Context) {
	name, ok := ctx.UserValue("name").(string)
	if !ok || name == "" {
		ctx.Error("Empty name", 200)
		return
	}

	var id int
	if id, ok = users[name]; !ok {
		ctx.Error("The user does not exist.", 200)
		return
	}

	token := jwt.NewWithClaims(signingMethod, UserClaims{
		ID:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 600,
		},
	})
	s, err := token.SignedString(signKey)
	if err != nil {
		ctx.Logger().Errorln(err)
		return
	}

	ctx.HTML(200, "token: "+s)
}

func homeHandler(ctx *gem.Context) {
	claims, ok := ctx.UserValue("jwt_claims").(*UserClaims)
	if !ok {
		ctx.Error("Invalid jwt token", fasthttp.StatusBadRequest)
		return
	}

	/*claims, ok := token.Claims.(UserClaims)
	if !ok {
		ctx.Error("Invalid claims type", fasthttp.StatusBadRequest)
		return
	}*/

	ctx.HTML(200, fmt.Sprintf("ID: %d, name: %s.", claims.ID, claims.Name))
}

func main() {
	router := gem.NewRouter()

	router.GET("/token/:name", generateHandler)

	router.Use(jwtMiddleware)

	router.GET("/home", homeHandler)

	log.Fatal(gem.ListenAndServe(":1234", router.Handler))
}
