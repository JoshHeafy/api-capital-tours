package main

import (
	"api-capital-tours/src/middleware"
	"api-capital-tours/src/routes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	middleware.EnableCORS(r)
	r.HandleFunc("/", HomeHandler)

	routes.RutasAuth(r)
	routes.RutasSolicitudes(r)
	routes.RutasPropietarios(r)
	routes.RutasVehiculos(r)
	routes.RutasInscripciones(r)
	routes.RutasComprobantes(r)

	fmt.Println("Server on port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := map[string]interface{}{"api": "Api CapitalTours", "version": "3.0.0"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
