package filters

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/emicklei/go-restful"
)

type jwtFilter struct {
	publicKey []byte
}

func NewJWTFilter(publicKey []byte) *jwtFilter {
	return &jwtFilter{
		publicKey: publicKey,
	}
}

func (f *jwtFilter) Filter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	// if the user is logging in for the first time then the
	// path will be /login. If that is the case then we just
	// go to the next filter because there is no token to
	// authenticate against.
	if req.Request.URL.Path != "/users/login" {

		token, err := jwt.ParseFromRequest(req.Request, func(token *jwt.Token) (interface{}, error) {
			return f.publicKey, nil
		})
		if err != nil || !token.Valid {
			fmt.Printf("invalid token for url %s: %s\n ", req.Request.URL.Path, err)
			resp.WriteErrorString(http.StatusUnauthorized, "Not authorized")
			return
		}
	}
	chain.ProcessFilter(req, resp)
}

func (f *jwtFilter) getToken(token *jwt.Token) ([]byte, error) {
	return f.publicKey, nil
}
