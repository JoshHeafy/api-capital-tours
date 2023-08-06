package routes

import (
	"api-capital-tours/src/controller"
	"api-capital-tours/src/database/models/tables"
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
	s.Handle("/info/{numero_documento}", middleware.Autentication(http.HandlerFunc(getOneInscripcion))).Methods("GET")
	s.Handle("/create-ins", middleware.Autentication(http.HandlerFunc(insertInscripciones))).Methods("POST")
	s.Handle("/service-alta/{numero_placa}", middleware.Autentication(http.HandlerFunc(darAlta))).Methods("PUT")
	s.Handle("/service-baja/{numero_placa}", middleware.Autentication(http.HandlerFunc(darBaja))).Methods("PUT")
	s.HandleFunc("/periodo/{numero_placa}", consultaPeriodo).Methods("GET")
}

func getInscripciones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	_data_inscripciones, _ := new(go_basic_orm.Querys).NewQuerys("inscripciones").Select().OrderBy("numero_documento").Exec(go_basic_orm.Config_Query{Cloud: true}).All()

	if len(_data_inscripciones) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontro suscripciónes"))
		return
	}

	response.Data["inscripciones"] = _data_inscripciones
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getOneInscripcion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_documento := params["numero_documento"]

	_data_inscripciones, _ := new(go_basic_orm.Querys).NewQuerys("inscripciones").Select().Where("numero_documento", "=", numero_documento).Exec(go_basic_orm.Config_Query{Cloud: true}).All()
	if len(_data_inscripciones) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data["inscripciones-info"] = _data_inscripciones

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

	err = inscripciones.Exec("capital_tours")
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

	err = inscripciones_detail.Exec("capital_tours")
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = inscripciones.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// - Opcion1: Agregar el valor de fecha_fin y estado 0 y generar una nueva inscripcion para dar de alta
// - Opcion2: Validar las fechas de los periodos para asi seleccionar solo los ordenados pero mas actuales
// - Opcion3: Borrar comprobantes antiguos al dar de alta, reiniciar inscripcion desde la fecha de alta

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

	err = servicio.Exec("capital_tours")
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

	_data_inscripcion, err := new(go_basic_orm.Querys).NewQuerys("inscripciones").Select("id_inscripcion,fecha_fin,fecha_pago,importe,estado,numero_flota,numero_placa,numero_documento").Where("numero_placa", "=", numero_placa).Exec(go_basic_orm.Config_Query{Cloud: true}).One()
	if err != nil {
		controller.ErrorsWaning(w, errors.New("no se encontró inscripción de este vehículo"))
		return
	}
	if _data_inscripcion["fecha_fin"] != nil {
		controller.ErrorsSuccess(w, errors.New("vehículo dado de baja"))
		return
	}

	_data_detalle_comprobante, err_detail_comp := new(go_basic_orm.Querys).NewQuerys("detalle_comprobantes dc").Select("months || '/' || years as periodo").InnerJoin("comprobante_pago cp", "cp.id_comprobante_pago = dc.id_comprobante_pago").Where("cp.id_inscripcion", "=", _data_inscripcion["id_inscripcion"]).Exec(go_basic_orm.Config_Query{Cloud: true}).All()
	if err_detail_comp != nil {
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

	for i := yearInit; i <= yearNow; i++ {
		for e := monthInit; e <= month; e++ {
			var estado uint64 = 0
			if e == int64(monthNow) && yearInit == yearNow {
				estado = 1
			}
			if e < int64(monthNow) || yearInit < yearNow {
				estado = 2
			}
			year := fmt.Sprintf("%v", i)
			month := fmt.Sprintf("%d", e)
			if library.IndexOf_String(newFact, month+"/"+year) == -1 {
				dataPagos = append(dataPagos, map[string]interface{}{
					"years":   year,
					"months":  month,
					"importe": importe,
					"estado":  estado, // 0: Ok, 1: Cerca, 2: Mora
				})
			}
		}

		monthInit = 1
	}

	response.Data["periodo-inscripcion"] = dataPagos
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
