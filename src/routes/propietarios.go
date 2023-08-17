package routes

import (
	"api-capital-tours/src/controller"
	"api-capital-tours/src/database/models/tables"
	"api-capital-tours/src/database/orm"
	"api-capital-tours/src/middleware"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func RutasPropietarios(r *mux.Router) {
	s := r.PathPrefix("/propietarios").Subrouter()

	s.Handle("/list", middleware.Autentication(http.HandlerFunc(getAllPropietarios))).Methods("GET")
	s.Handle("/create", middleware.Autentication(http.HandlerFunc(insertPropietarios))).Methods("POST")
	s.Handle("/info-prop/{numero_documento}", middleware.Autentication(http.HandlerFunc(getOnePropietarioByDocument))).Methods("GET")
	s.Handle("/update/{numero_documento}", middleware.Autentication(http.HandlerFunc(updatePropietario))).Methods("PUT")
	s.Handle("/filter/{filtro}", middleware.Autentication(http.HandlerFunc(propietariosFilter))).Methods("GET")
}

func getAllPropietarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	_data_propietarios := orm.NewQuerys("propietarios").Select().OrderBy("nombre_propietario").Exec(orm.Config_Query{Cloud: true}).All()
	if len(_data_propietarios) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontro propietarios"))
		return
	}

	response.Data["propietarios"] = _data_propietarios
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
	propietarios := orm.SqlExec{}
	err = propietarios.New(data_insert, table).Insert(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = propietarios.Exec()
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = propietarios.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getOnePropietarioByDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_documento := params["numero_documento"]

	_data_propietario := orm.NewQuerys("propietarios").Select().Where("numero_documento", "=", numero_documento).Exec(orm.Config_Query{Cloud: true}).One()

	if len(_data_propietario) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data["propietario_info"] = _data_propietario

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func updatePropietario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_documento := params["numero_documento"]

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}
	data_request["where"] = map[string]interface{}{"numero_documento": numero_documento}

	var data_update []map[string]interface{}
	data_update = append(data_update, data_request)

	schema, table := tables.Propietarios_GetSchema()
	propietarios := orm.SqlExec{}
	err = propietarios.New(data_update, table).Update(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = propietarios.Exec()
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = propietarios.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func propietariosFilter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	filtro := params["filtro"]

	_data_propietarios := orm.NewQuerys("propietarios").Select().Like("numero_documento", "%"+filtro+"%").OrLike("nombre_propietario", "%"+filtro+"%").OrderBy("nombre_propietario").Exec(orm.Config_Query{Cloud: true}).All()

	response.Data["propietarios"] = _data_propietarios
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
