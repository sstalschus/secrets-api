package main

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/sstalschus/secrets-api/cmd/secrets-api/middlewares"
	"github.com/sstalschus/secrets-api/infra/log"
	"github.com/sstalschus/secrets-api/internal"
)

var routes []route

func initializeRoutes() {
	routes = []route{
		newRoute("GET", "/ping", ping),
		newRoute("POST", "/users",
			middlewares.HandleRequestID(
				middlewares.RequestLogger(userController.SignUp, logger))),

		newRoute("GET", "/users",
			middlewares.HandleRequestID(
				middlewares.RequestLogger(
					middlewares.EnsureAuth(
						userController.GetUser, authProvider, logger), logger))),

		newRoute("POST", "/token",
			middlewares.HandleRequestID(
				middlewares.RequestLogger(authController.Authenticate, logger))),

		newRoute("POST", "/secrets",
			middlewares.HandleRequestID(
				middlewares.RequestLogger(
					middlewares.EnsureAuth(
						secretController.CreateSecret, authProvider, logger), logger))),

		newRoute("GET", "/secrets",
			middlewares.HandleRequestID(
				middlewares.RequestLogger(
					middlewares.EnsureAuth(
						secretController.GetSecrets, authProvider, logger), logger))),

		newRoute("GET", "/secrets/([^/]+)",
			middlewares.HandleRequestID(
				middlewares.RequestLogger(
					middlewares.EnsureAuth(
						secretController.GetSecret, authProvider, logger), logger))),
	}
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

func Server(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	initializeRoutes()
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := internal.CtxWithValues(r.Context(), log.Body{
				"CtxKey": matches[1:],
			})
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pong"))
}

func Handler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization")
			w.Header().Set("Access-Control-Expose-Credentials", "true")
			w.Header().Set("Access-Control-Max-Age", "30000")
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		} else {
			Server(w, r)
		}
	}
}

func enableCors(w *http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
		(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
		(*w).Header().Set("Access-Control-Expose-Headers", "Authorization")
		(*w).Header().Set("Access-Control-Expose-Credentials", "true")
		(*w).Header().Set("Access-Control-Max-Age", "30000")
		(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
		return
	}
}
