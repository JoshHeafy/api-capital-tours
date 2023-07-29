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

func RutasPropietarios(r *mux.Router) {
	s := r.PathPrefix("/propietarios").Subrouter()

	s.Handle("/list", middleware.Autentication(http.HandlerFunc(getAllPropietarios))).Methods("GET")
	s.Handle("/create", middleware.Autentication(http.HandlerFunc(insertPropietarios))).Methods("POST")
	s.Handle("/info-prop/{numero_documento}", middleware.Autentication(http.HandlerFunc(getOnePropietarioByDocument))).Methods("GET")
	s.Handle("/update/{numero_documento}", middleware.Autentication(http.HandlerFunc(updatePropietario))).Methods("PUT")
}

func getAllPropietarios(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	data_propietarios, _ := new(go_basic_orm.Querys).NewQuerys("propietarios").Select().Exec(go_basic_orm.Config_Query{Cloud: true}).All()
	response.Data["propietarios"] = data_propietarios
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertPropietarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	data_insert := append([]map[string]interface{}{}, data_request)

	schema, table := tables.Propietarios_GetSchema()
	propietarios := go_basic_orm.SqlExec{}
	err = propietarios.New(data_insert, table).Insert(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = propietarios.Exec("capital_tours")
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = propietarios.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getOnePropietarioByDocument(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_documento := params["numero_documento"]
	if numero_documento == "" {
		controller.ErrorsError(w, errors.New("documento no encontrado"))
		return
	}

	data_propietario, _ := new(go_basic_orm.Querys).NewQuerys("propietarios").Select().Where("numero_documento", "=", numero_documento).Exec(go_basic_orm.Config_Query{Cloud: true}).One()

	if len(data_propietario) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data["propietario-info"] = data_propietario

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func updatePropietario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_documento := params["numero_documento"]

	if numero_documento == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}
	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}
	data_request["where"] = map[string]interface{}{"numero_documento": numero_documento}

	var data_update []map[string]interface{}
	data_update = append(data_update, data_request)

	schema, table := tables.Propietarios_GetSchema()
	propietarios := go_basic_orm.SqlExec{}
	err = propietarios.New(data_update, table).Update(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = propietarios.Exec("capital_tours")
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = propietarios.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
