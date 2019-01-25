package middlewares

import (
	"net/http"

	"github.com/urantiatech/beego"
)

// AuthMiddleware authenticates
type AuthMiddleware struct {
	handler http.Handler
}

func (a *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler.ServeHTTP(w, r)
}
