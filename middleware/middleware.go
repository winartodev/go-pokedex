package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/winartodev/go-pokedex/enum"
	"github.com/winartodev/go-pokedex/helper"
	"github.com/winartodev/go-pokedex/middleware/auth"
)

// Auth will validate url path, and role of user obtained from jwt token
func Auth(handle httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				helper.FailedResponse(w, http.StatusUnauthorized, fmt.Errorf("user is not logged in"))
				return
			}
			helper.FailedResponse(w, http.StatusBadRequest, err)
			return
		}

		sessionToken := c.Value
		claims, err := auth.ValidateToken(sessionToken)
		if err != nil {
			helper.FailedResponse(w, http.StatusUnauthorized, err)
			return
		}

		urlPath := strings.Split(r.URL.Path, "/")[1]
		switch urlPath {
		case "internal":
			if claims.Role != enum.Admin {
				helper.FailedResponse(w, http.StatusUnauthorized, fmt.Errorf("user not %s", enum.Admin.String()))
				return
			}
		case "user":
			if claims.Role != enum.User {
				helper.FailedResponse(w, http.StatusUnauthorized, fmt.Errorf("user is not logged in"))
				return
			}
		}

		handle(w, r, p)
	})
}
