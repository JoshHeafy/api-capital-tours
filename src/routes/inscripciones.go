package routes

import (
	"api-capital-tours/src/controller"
	"api-capital-tours/src/database/models/tables"
	"api-capital-tours/src/libraries/date"
	"api-capital-tours/src/middleware"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/deybin/go_basic_orm"
	"github.com/gorilla/mux"
)

func RutasInscripciones(r *mux.Router) {
	s := r.PathPrefix("/inscripciones").Subrouter()

	s.Handle("/list", middleware.Autentication(http.HandlerFunc(getInscripciones))).Methods("GET")
	s.Handle("/info/{numero_documento}", middleware.Autentication(http.HandlerFunc(getOneInscripcion))).Methods("GET")
	s.Handle("/create-ins", middleware.Autentication(http.HandlerFunc(insertInscripciones))).Methods("POST")
	s.Handle("/service-alta/{numero_placa}", middleware.Autentication(http.HandlerFunc(darAlta))).Methods("PUT")
	s.Handle("/service-baja/{numero_placa}", middleware.Autentication(http.HandlerFunc(darBaja))).Methods("PUT")
}

func getInscripciones(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	//get allData from database
	data_inscripciones, _ := new(go_basic_orm.Querys).NewQuerys("inscripciones").Select().Exec(go_basic_orm.Config_Query{Cloud: true}).All()
	response.Data["inscripciones"] = data_inscripciones
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertInscripciones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	data_insert := make(map[string]interface{})

	for key, value := range data_request {
		data_insert[key] = value
	}

	data_insert["fecha_inicio"] = date.GetFechaLocationString()
	data_insert["years"] = date.GetYear()
	data_insert["months"] = date.GetMonth()
	fechaPago := time.Now().AddDate(0, 1, 0)
	data_insert["fecha_pago"] = string(fechaPago.Format("02/01/2006"))

	schema, table := tables.Inscripciones_GetSchema()
	inscripciones := go_basic_orm.SqlExec{}
	err = inscripciones.New([]map[string]interface{}{data_insert}, table).Insert(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = inscripciones.Exec("capital_tours")
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = inscripciones.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func darAlta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	numero_placa := params["numero_placa"]
	if numero_placa == "" {
		response.Msg = "Error al escribir el servicio"
		response.StatusCode = 400
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	data_body := make(map[string]interface{})
	data_body["estado"] = int64(1)

	data_body["where"] = map[string]interface{}{"numero_placa": numero_placa}

	var data_update []map[string]interface{}
	data_update = append(data_update, data_body)

	schema, table := tables.Inscripciones_GetSchema()
	servicio := go_basic_orm.SqlExec{}
	err := servicio.New(data_update, table).Update(schema)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	err = servicio.Exec("capital_tours")
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}
	response.Data = servicio.Data[0]

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func darBaja(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	numero_placa := params["numero_placa"]
	if numero_placa == "" {
		response.Msg = "Error al escribir el servicio"
		response.StatusCode = 400
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	data_body := make(map[string]interface{})
	data_body["estado"] = int64(0)

	data_body["where"] = map[string]interface{}{"numero_placa": numero_placa}

	var data_update []map[string]interface{}
	data_update = append(data_update, data_body)

	schema, table := tables.Inscripciones_GetSchema()
	servicio := go_basic_orm.SqlExec{}
	err := servicio.New(data_update, table).Update(schema)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = servicio.Exec("capital_tours")
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = servicio.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getOneInscripcion(w http.ResponseWriter, r *http.Request) {
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

	data_inscripciones, _ := new(go_basic_orm.Querys).NewQuerys("inscripciones").Select().Where("numero_documento", "=", numero_documento).Exec(go_basic_orm.Config_Query{Cloud: true}).All()

	if len(data_inscripciones) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data["inscripciones-info"] = data_inscripciones

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
