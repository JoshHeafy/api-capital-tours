package routes

import (
	"api-capital-tours/src/controller"
	"api-capital-tours/src/database/models/tables"
	"api-capital-tours/src/middleware"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/deybin/go_basic_orm"
	"github.com/gorilla/mux"
)

func RutasSolicitudes(r *mux.Router) {
	s := r.PathPrefix("/solicitudes").Subrouter()
	s.HandleFunc("/create", insertSolicitud).Methods("POST")
	s.Handle("/list", middleware.Autentication(http.HandlerFunc(getSolicitudes))).Methods("GET")
}
func getSolicitudes(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	data_solicitudes, err := new(go_basic_orm.Querys).NewQuerys("solicitudes").Select().Exec(go_basic_orm.Config_Query{Cloud: true}).All()

	if err != nil {
		controller.ErrorsWaning(w, errors.New("error al obtener solicitudes"))
		return
	}

	response.Data["solicitudes"] = data_solicitudes
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertSolicitud(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	data_insert := append([]map[string]interface{}{}, data_request)

	schema, table := tables.Solicitudes_GetSchema()
	solicitudes := go_basic_orm.SqlExec{}
	err = solicitudes.New(data_insert, table).Insert(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = solicitudes.Exec("capital_tours")
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = solicitudes.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
