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

func RutasVehiculos(r *mux.Router) {
	s := r.PathPrefix("/vehiculos").Subrouter()

	s.Handle("/list", middleware.Autentication(http.HandlerFunc(getVehiculos))).Methods("GET")
	s.Handle("/info/{numero_documento}", middleware.Autentication(http.HandlerFunc(getClientVehiculos))).Methods("GET")
	s.Handle("/info-placa/{numero_placa}", middleware.Autentication(http.HandlerFunc(getClientVehiculoByPlaca))).Methods("GET")
	s.Handle("/create", middleware.Autentication(http.HandlerFunc(insertVehiculos))).Methods("POST")
	s.Handle("/update/{numero_placa}", middleware.Autentication(http.HandlerFunc(updateVehiculo))).Methods("PUT")
	s.Handle("/re-assign/{numero_placa}", middleware.Autentication(http.HandlerFunc(reAssignVehiculo))).Methods("PUT")
}

func getVehiculos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	_data_vehiculos := orm.NewQuerys("vehiculos").Select().OrderBy("numero_documento").Exec(orm.Config_Query{Cloud: true}).All()
	if len(_data_vehiculos) <= 0 {
		controller.ErrorsWaning(w, errors.New("error al obtener vehículos"))
		return
	}

	response.Data["vehiculos"] = _data_vehiculos
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertVehiculos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("error al leer el cuerpo de la solicitud"))
		return
	}

	data_insert := append([]map[string]interface{}{}, data_request)

	schema, table := tables.Vehiculos_GetSchema()
	vehiculos := orm.SqlExec{}

	if err := vehiculos.New(data_insert, table).Insert(schema); err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	if err := vehiculos.Exec(); err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("hubo un error al crear el vehículo, por favor intente nuevamente o comuniquese con el administrador"))
		return
	}

	response.Data = vehiculos.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getClientVehiculos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_documento := params["numero_documento"]

	_data_client_car := orm.NewQuerys("vehiculos").Select().Where("numero_documento", "=", numero_documento).Exec(orm.Config_Query{Cloud: true}).All()
	if len(_data_client_car) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data["vehiculos_info"] = _data_client_car

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getClientVehiculoByPlaca(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_placa := params["numero_placa"]

	_data_client_car := orm.NewQuerys("vehiculos").Select().Where("numero_placa", "=", numero_placa).Exec(orm.Config_Query{Cloud: true}).One()
	if len(_data_client_car) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontró este vehículo"))
		return
	}

	response.Data["vehiculo_info"] = _data_client_car

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func updateVehiculo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_placa := params["numero_placa"]

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("error al leer el cuerpo de la solicitud"))
		return
	}
	data_request["where"] = map[string]interface{}{"numero_placa": numero_placa}

	var data_update []map[string]interface{}
	data_update = append(data_update, data_request)

	schema, table := tables.Vehiculos_GetSchema()
	vehiculos := orm.SqlExec{}

	if err := vehiculos.New(data_update, table).Update(schema); err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	if err := vehiculos.Exec(); err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("hubo un error al actualizar el vehículo, por favor intente nuevamente o comuniquese con el administrador"))
		return
	}

	response.Data = vehiculos.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func reAssignVehiculo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_placa := params["numero_placa"]

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("error al leer el cuerpo de la solicitud"))
		return
	}

	var data_update_vehiculo []map[string]interface{}
	var data_update_inscripcion []map[string]interface{}
	var data_update_inscripcion_detail []map[string]interface{}
	data_update_vehiculo = append(data_update_vehiculo, map[string]interface{}{
		"numero_documento": data_request["numero_documento"],
		"where": map[string]interface{}{
			"numero_placa": numero_placa,
		},
	})

	_data_inscripcion := orm.NewQuerys("inscripciones").Select().Where("numero_placa", "=", numero_placa).Exec(orm.Config_Query{Cloud: true}).All()

	var inscripciones_data map[string]interface{}
	var inscripciones_detail_data map[string]interface{}

	if len(_data_inscripcion) > 0 {
		data_update_inscripcion = append(data_update_inscripcion, map[string]interface{}{
			"numero_documento": data_request["numero_documento"],
			"where": map[string]interface{}{
				"numero_placa": numero_placa,
			},
		})

		data_update_inscripcion_detail = append(data_update_inscripcion_detail, map[string]interface{}{
			"numero_documento": data_request["numero_documento"],
			"where": map[string]interface{}{
				"id_inscripcion": _data_inscripcion[0]["id_inscripcion"],
			},
		})

		inscripciones_data = reAssingInscripcionInsert(w, data_update_inscripcion)
		inscripciones_detail_data = reAssingInscripcionDetailInsert(w, data_update_inscripcion_detail)
	}

	schema, table := tables.Vehiculos_GetSchema()
	vehiculos := orm.SqlExec{}

	if err := vehiculos.New(data_update_vehiculo, table).Update(schema); err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	if err = vehiculos.Exec(); err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("hubo un error al re asignar el vehiculo, por favor intente nuevamente o comuniquese con el administrador"))
		return
	}

	response.Data["update_vehiculo"] = vehiculos.Data[0]
	response.Data["update_ins"] = inscripciones_data
	response.Data["update_ins_det"] = inscripciones_detail_data
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func reAssingInscripcionInsert(w http.ResponseWriter, data []map[string]interface{}) map[string]interface{} {
	schemaIns, tableIns := tables.Inscripciones_GetSchema()
	inscripciones := orm.SqlExec{}

	if errIns := inscripciones.New(data, tableIns).Update(schemaIns); errIns != nil {
		controller.ErrorsWaning(w, errIns)
		return map[string]interface{}{}
	}

	if errIns := inscripciones.Exec(); errIns != nil {
		log.Println(errIns)
		controller.ErrorsWaning(w, errors.New("hubo un error al re asignar la inscripcion, por favor intente nuevamente o comuniquese con el administrador"))
		return map[string]interface{}{}
	}

	return inscripciones.Data[0]
}

func reAssingInscripcionDetailInsert(w http.ResponseWriter, data []map[string]interface{}) map[string]interface{} {
	schema, table := tables.Detalleinscripciones_GetSchema()
	inscripciones_detail := orm.SqlExec{}

	if err := inscripciones_detail.New(data, table).Update(schema); err != nil {
		controller.ErrorsWaning(w, err)
		return map[string]interface{}{}
	}

	if err := inscripciones_detail.Exec(); err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("hubo un error al re asignar el detalle de la inscripcion, por favor intente nuevamente o comuniquese con el administrador"))
		return map[string]interface{}{}
	}

	return inscripciones_detail.Data[0]
}
