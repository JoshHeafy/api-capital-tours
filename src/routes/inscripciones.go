package routes

import (
	"api-capital-tours/src/auth"
	"api-capital-tours/src/controller"
	"api-capital-tours/src/database/models/tables"
	"api-capital-tours/src/database/orm"
	"api-capital-tours/src/libraries/date"
	"api-capital-tours/src/libraries/library"
	"api-capital-tours/src/middleware"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/deybin/go_basic_orm"
	"github.com/gorilla/mux"
)

func RutasInscripciones(r *mux.Router) {
	s := r.PathPrefix("/inscripciones").Subrouter()

	s.Handle("/list", middleware.Autentication(http.HandlerFunc(getInscripciones))).Methods("GET")
	s.Handle("/list-last", middleware.Autentication(http.HandlerFunc(getLastInscripciones))).Methods("GET")
	s.Handle("/info/{numero_documento}", middleware.Autentication(http.HandlerFunc(getInscripcionByClient))).Methods("GET")
	s.Handle("/create-ins", middleware.Autentication(http.HandlerFunc(insertInscripciones))).Methods("POST")
	s.Handle("/service-alta/{numero_placa}", middleware.Autentication(http.HandlerFunc(darAlta))).Methods("PUT")
	s.Handle("/service-baja/{numero_placa}", middleware.Autentication(http.HandlerFunc(darBaja))).Methods("PUT")
	s.HandleFunc("/periodo/{numero_placa}", consultaPeriodo).Methods("GET")
}

func getInscripciones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	_data_inscripciones := orm.NewQuerys("inscripciones").Select().OrderBy("numero_flota").Exec(orm.Config_Query{Cloud: true}).All()

	if len(_data_inscripciones) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontro suscripciónes"))
		return
	}

	response.Data["inscripciones"] = _data_inscripciones
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getLastInscripciones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	_data_inscripciones := orm.NewQuerys("inscripciones").Select().OrderBy("TO_DATE(fecha_inicio, 'DD/MM/YYYY') DESC").Exec(orm.Config_Query{Cloud: true}).All()

	if len(_data_inscripciones) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontro suscripciónes"))
		return
	}

	var inscripciones []map[string]interface{}
	for i, ins := range _data_inscripciones {
		if i < 5 {
			inscripciones = append(inscripciones, ins)
		}
	}

	response.Data["inscripciones"] = inscripciones
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getInscripcionByClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_documento := params["numero_documento"]

	_data_inscripciones := orm.NewQuerys("inscripciones").Select().Where("numero_documento", "=", numero_documento).OrderBy("numero_flota").Exec(orm.Config_Query{Cloud: true}).All()
	if len(_data_inscripciones) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data["inscripciones_info"] = _data_inscripciones

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

	_data_inscripcion := orm.NewQuerys("inscripciones").Select().Where("numero_placa", "=", data_request["numero_placa"]).Exec(orm.Config_Query{Cloud: true}).All()

	if len(_data_inscripcion) >= 1 {
		controller.ErrorsWaning(w, errors.New("este vehiculo ya tiene una suscripción"))
		return
	}

	data_insert := make(map[string]interface{})
	data_insert_detalle_inscripcion := make(map[string]interface{})

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

	err = inscripciones.Exec(auth.GetDBName())
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	data_insert_detalle_inscripcion["fecha_pago"] = inscripciones.Data[0]["fecha_pago"]
	data_insert_detalle_inscripcion["years"] = inscripciones.Data[0]["years"]
	data_insert_detalle_inscripcion["months"] = inscripciones.Data[0]["months"]
	data_insert_detalle_inscripcion["importe"] = inscripciones.Data[0]["importe"]
	data_insert_detalle_inscripcion["numero_documento"] = inscripciones.Data[0]["numero_documento"]
	data_insert_detalle_inscripcion["id_inscripcion"] = inscripciones.Data[0]["id_inscripcion"]

	schema_ins_detail, table := tables.Detalleinscripciones_GetSchema()
	inscripciones_detail := go_basic_orm.SqlExec{}
	err = inscripciones_detail.New([]map[string]interface{}{data_insert_detalle_inscripcion}, table).Insert(schema_ins_detail)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = inscripciones_detail.Exec(auth.GetDBName())
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

	data_request, _ := controller.CheckBody(w, r)

	params := mux.Vars(r)
	numero_placa := params["numero_placa"]

	fechaPago := time.Now().AddDate(0, 1, 0)

	var data_update []map[string]interface{}
	updateMap := map[string]interface{}{
		"estado":       uint64(1),
		"importe":      data_request["importe"].(float64),
		"fecha_fin":    "",
		"years":        date.GetYear(),
		"months":       date.GetMonth(),
		"fecha_inicio": date.GetFechaLocationString(),
		"fecha_pago":   string(fechaPago.Format("02/01/2006")),
		"where": map[string]interface{}{
			"numero_placa": numero_placa,
		},
	}

	data_update = append(data_update, updateMap)

	schema, table := tables.Inscripciones_GetSchema()
	servicio := go_basic_orm.SqlExec{}
	err := servicio.New(data_update, table).Update(schema)
	if err != nil {
		controller.ErrorsWaning(w, errors.New(err.Error()))
		return
	}
	err = servicio.Exec(auth.GetDBName())
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

	var data_update []map[string]interface{}
	data_update = append(data_update, map[string]interface{}{
		"estado":    int64(0),
		"fecha_fin": date.GetFechaLocationString(),
		"where": map[string]interface{}{
			"numero_placa": numero_placa,
		},
	})

	schema, table := tables.Inscripciones_GetSchema()
	servicio := go_basic_orm.SqlExec{}
	err := servicio.New(data_update, table).Update(schema)
	if err != nil {
		controller.ErrorsWaning(w, errors.New(err.Error()))
		return
	}

	err = servicio.Exec(auth.GetDBName())
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = servicio.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func consultaPeriodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_placa := params["numero_placa"]

	_data_inscripcion := orm.NewQuerys("inscripciones").Select("id_inscripcion,fecha_fin,fecha_pago,importe,estado,numero_flota,numero_placa,numero_documento").Where("numero_placa", "=", numero_placa).Exec(orm.Config_Query{Cloud: true}).One()
	if len(_data_inscripcion) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontró inscripción de este vehículo"))
		return
	}
	if _data_inscripcion["fecha_fin"] != nil {
		controller.ErrorsSuccess(w, errors.New("vehículo dado de baja"))
		return
	}

	_data_detalle_comprobante := orm.NewQuerys("detalle_comprobantes dc").Select("months || '/' || years as periodo").InnerJoin("comprobante_pago cp", "cp.id_comprobante_pago = dc.id_comprobante_pago").Where("cp.id_inscripcion", "=", _data_inscripcion["id_inscripcion"]).Exec(orm.Config_Query{Cloud: true}).All()
	if len(_data_detalle_comprobante) <= 0 {
		controller.ErrorsWaning(w, errors.New("error al obtener comprobantes de pago"))
		return
	}

	var newFact []string
	for _, v := range _data_detalle_comprobante {
		if v["periodo"] != nil {
			newFact = append(newFact, v["periodo"].(string))
		}
	}

	datePago := date.GetDate(_data_inscripcion["fecha_pago"].(string))
	dateNow := date.GetDateLocation()

	monthInit := int64(datePago.Month())
	yearInit := datePago.Year()
	yearNow := dateNow.Year()
	monthNow := dateNow.Month()

	var dataPagos []map[string]interface{}
	var month = int64(12)

	importe := _data_inscripcion["importe"]

	var new_date_fact string
	loc, _ := time.LoadLocation("America/Bogota")
	dia := datePago.Day()
	if dia > 28 {
		d := time.Date(yearNow, time.Month(monthNow), 1, 0, 0, 0, 0, loc)
		f := date.GetLastDateOfMonth(d)
		if f.Day() <= dia {
			dia = f.Day()
		}
	}

	new_date_fact = fmt.Sprintf("%02d/%02d/%d", dia, monthNow+1, yearNow)

	for i := yearInit; i <= yearNow; i++ {
		for e := monthInit; e <= month; e++ {
			var estado uint64 = 1

			if e <= int64(dateNow.Month()) || i < yearNow {
				var date_difference_list []string

				date_difference_list = append(date_difference_list, fmt.Sprintf("%02d/%02d/%d",
					dia, e, i), fmt.Sprintf("%02d/%02d/%d", dateNow.Day(), dateNow.Month(), dateNow.Year()))

				diff := date.DiferenciaDate(date_difference_list...)

				if diff < 0 {
					estado = 2
				}
			}

			if e == int64(monthNow) && i == yearNow {
				dia := datePago.Day()
				if dia > 28 {
					d := time.Date(i, time.Month(e), 1, 0, 0, 0, 0, loc)
					f := date.GetLastDateOfMonth(d)
					if f.Day() <= dia {
						dia = f.Day()
					}
				}
				new_date_fact = fmt.Sprintf("%02d/%02d/%d", dia, e, i)
				estado = 0
			}

			periodo := fmt.Sprintf("%d/%d", e, i)
			if library.IndexOf_String(newFact, periodo) == -1 {
				dataPagos = append(dataPagos, map[string]interface{}{
					"years":   i,
					"months":  e,
					"importe": importe,
					"estado":  estado, // 0: Cerca, 1: Ok, 2: Mora

				})
			}
		}

		monthInit = 1
	}

	response.Data["periodo_inscripcion"] = dataPagos
	response.Data["status_pago"] = map[string]interface{}{
		"proximo_pago": new_date_fact,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
