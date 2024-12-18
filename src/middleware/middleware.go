package middleware

import (
	"api-capital-tours/src/auth"
	"api-capital-tours/src/controller"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Autentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := controller.NewResponseManager()
		token := r.Header.Get("Access-Token")
		_, err := auth.ValidateToken(token)

		if token == "" || err != nil {
			response.Msg = "Inicio de Sesión Inválido"
			response.Status = "error"
			response.StatusCode = 401
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func MiddlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Lista de dominios permitidos
			allowedOrigins := map[string]bool{
				"http://localhost:3000":                  true,
				"http://localhost:3001":                  true,
				"https://capital-tours.vercel.app":       true,
				"https://capital-tours-admin.vercel.app": true,
			}

			origin := req.Header.Get("Origin")

			if allowedOrigins[origin] {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}

			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Auth-Date, Auth-Periodo, Access-Token")

			if req.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, req)
		},
	)
}

func EnableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Auth-Date, Auth-Periodo, Access-Token")
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodOptions)

	router.Use(MiddlewareCors)
}
