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

func RutasSolicitudes(r *mux.Router) {
	s := r.PathPrefix("/solicitudes").Subrouter()
	s.HandleFunc("/create", insertSolicitud).Methods("POST")
	s.Handle("/mark-as-read/{id_solicitud}", middleware.Autentication(http.HandlerFunc(markAsRead))).Methods("POST")
	s.Handle("/list", middleware.Autentication(http.HandlerFunc(getSolicitudes))).Methods("GET")
}

func getSolicitudes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	_data_solicitudes := orm.NewQuerys("solicitudes").Select().OrderBy("fecha_solicitud").Exec(orm.Config_Query{Cloud: true}).All()

	if len(_data_solicitudes) <= 0 {
		controller.ErrorsWaning(w, errors.New("error al obtener solicitudes"))
		return
	}

	response.Data["solicitudes"] = _data_solicitudes
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertSolicitud(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("error al leer el cuerpo de la solicitud"))
		return
	}

	data_insert := append([]map[string]interface{}{}, data_request)

	schema, table := tables.Solicitudes_GetSchema()
	solicitudes := orm.SqlExec{}

	if err := solicitudes.New(data_insert, table).Insert(schema); err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	if err := solicitudes.Exec(); err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("hubo un error al crear la solicitud, por favor intente nuevamente o comuniquese con el administrador"))
		return
	}

	response.Data = solicitudes.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func markAsRead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	id_solicitud := params["id_solicitud"]

	var data_update []map[string]interface{}
	data_update = append(data_update, map[string]interface{}{
		"where": map[string]interface{}{"id_solicitud": id_solicitud},
		"leido": uint64(1),
	})

	schema, table := tables.Solicitudes_GetSchema()
	solicitudes := orm.SqlExec{}

	if err := solicitudes.New(data_update, table).Update(schema); err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	if err := solicitudes.Exec(); err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("hubo un error al actualizar la solicitud, por favor intente nuevamente o comuniquese con el administrador"))
		return
	}

	response.Data = solicitudes.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
