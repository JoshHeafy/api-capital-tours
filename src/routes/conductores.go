package routes

import (
	"api-capital-tours/src/controller"
	"api-capital-tours/src/database/models/tables"
	"api-capital-tours/src/database/orm"
	"api-capital-tours/src/middleware"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RutasConductores(r *mux.Router) {
	s := r.PathPrefix("/conductores").Subrouter()

	s.Handle("/list", middleware.Autentication(http.HandlerFunc(getAllConductores))).Methods("GET")
	s.Handle("/create", middleware.Autentication(http.HandlerFunc(insertConductores))).Methods("POST")
	s.Handle("/info/{numero_licencia}", middleware.Autentication(http.HandlerFunc(getOneConductorByDocument))).Methods("GET")
	s.Handle("/update/{numero_licencia}", middleware.Autentication(http.HandlerFunc(updateConductor))).Methods("PUT")
	s.Handle("/filter/{filtro}", middleware.Autentication(http.HandlerFunc(conductoresFilter))).Methods("GET")
}

func getAllConductores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	_data_conductores := orm.NewQuerys("conductores").Select().OrderBy("nombre_conductor").Exec(orm.Config_Query{Cloud: true}).All()
	if len(_data_conductores) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontro conductores"))
		return
	}

	response.Data["conductores"] = _data_conductores
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertConductores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("error al leer el cuerpo de la solicitud"))
		return
	}

	var numero_licencia string
	if _, ok := data_request["numero_licencia"]; ok {
		numero_licencia = data_request["numero_licencia"].(string)
	}

	if numero_licencia != "" {
		_data_conductor := orm.NewQuerys("conductores").Select().Where("numero_licencia", "=", numero_licencia).Exec(orm.Config_Query{Cloud: true}).One()
		if len(_data_conductor) > 0 {
			controller.ErrorsWaning(w, errors.New("ya existe un conductor con el número de licencia: "+numero_licencia))
			return
		}
	}

	data_insert := append([]map[string]interface{}{}, data_request)

	schema, table := tables.Conductores_GetSchema()
	conductores := orm.SqlExec{}

	if err := conductores.New(data_insert, table).Insert(schema); err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	if err := conductores.Exec(); err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("hubo un error al crear el conductor, por favor intente nuevamente o comuniquese con el administrador"))
		return
	}

	response.Data = conductores.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getOneConductorByDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_licencia := params["numero_licencia"]

	_data_conductor := orm.NewQuerys("conductores").Select().Where("numero_licencia", "=", numero_licencia).Exec(orm.Config_Query{Cloud: true}).One()

	if len(_data_conductor) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data["conductor_info"] = _data_conductor

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func updateConductor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_licencia := params["numero_licencia"]

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}
	data_request["where"] = map[string]interface{}{"numero_licencia": numero_licencia}

	var data_update []map[string]interface{}
	data_update = append(data_update, data_request)

	schema, table := tables.Conductores_GetSchema()
	conductores := orm.SqlExec{}

	if err := conductores.New(data_update, table).Update(schema); err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	if err := conductores.Exec(); err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("hubo un error al actualizar el conductor, por favor intente nuevamente o comuniquese con el administrador"))
		return
	}

	response.Data = conductores.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func conductoresFilter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	filtro := params["filtro"]

	_data_conductores := orm.NewQuerys("conductores").Select().Ilike("numero_licencia", "%"+filtro+"%").OrIlike("nombre_conductor", "%"+filtro+"%").OrderBy("nombre_conductor").Exec(orm.Config_Query{Cloud: true}).All()

	response.Data["conductores"] = _data_conductores
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
