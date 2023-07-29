package routes

import (
	"api-capital-tours/src/controller"
	"api-capital-tours/src/database/models/tables"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/deybin/go_basic_orm"
	"github.com/gorilla/mux"
)

func RutasVehiculos(r *mux.Router) {
	s := r.PathPrefix("/vehiculos").Subrouter()

	s.HandleFunc("/list", getVehiculos).Methods("GET")
	s.HandleFunc("/info/{numero_documento}", getClientCar).Methods("GET")
	s.HandleFunc("/create", insertVehiculos).Methods("POST")
	s.HandleFunc("/update/{numero_placa}", updateVehiculo).Methods("PUT")

}

func getVehiculos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	//get allData from database
	dataVehiculos, _ := new(go_basic_orm.Querys).NewQuerys("vehiculos").Select().Exec(go_basic_orm.Config_Query{Cloud: true}).All()
	response.Data["vehiculos"] = dataVehiculos
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertVehiculos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	data_insert := append([]map[string]interface{}{}, data_request)

	schema, table := tables.Vehiculos_GetSchema()
	vehiculos := go_basic_orm.SqlExec{}
	err = vehiculos.New(data_insert, table).Insert(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = vehiculos.Exec("capital_tours")
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = vehiculos.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getClientCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_documento := params["numero_documento"]
	if numero_documento == "" {
		response.Msg = "Documento no encintrado"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	//get allData from database
	dataCientCar, _ := new(go_basic_orm.Querys).NewQuerys("vehiculos").Select().Where("numero_documento", "=", numero_documento).Exec(go_basic_orm.Config_Query{Cloud: true}).All()

	if len(dataCientCar) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data["vehiculos-info"] = dataCientCar
	// response.Data["cookie_token"] = sessionID

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func updateVehiculo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_placa := params["numero_placa"]

	if numero_placa == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}
	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}
	data_request["where"] = map[string]interface{}{"numero_placa": numero_placa}

	var data_update []map[string]interface{}
	data_update = append(data_update, data_request)

	schema, table := tables.Vehiculos_GetSchema()
	vehiculos := go_basic_orm.SqlExec{}
	err = vehiculos.New(data_update, table).Update(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = vehiculos.Exec("capital_tours")
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = vehiculos.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
