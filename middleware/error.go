package middleware

import (
	"net/http"
	"vermak2/constants"
	u "vermak2/utils"
)

var NotFoundHandler = func(next http.Handler)http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(constants.FalseStatus, "This resources not found"))
		next.ServeHTTP(w, r)
	})
}
