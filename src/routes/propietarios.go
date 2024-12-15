package routes

import (
	"api-capital-tours/src/controller"
	"api-capital-tours/src/database/models/tables"
	"api-capital-tours/src/database/orm"
	"api-capital-tours/src/libraries/date"
	"api-capital-tours/src/libraries/library"
	"api-capital-tours/src/middleware"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/gorilla/mux"
)

func RutasPropietarios(r *mux.Router) {
	s := r.PathPrefix("/propietarios").Subrouter()

	s.Handle("/list", middleware.Autentication(http.HandlerFunc(getAllPropietarios))).Methods("GET")
	s.Handle("/create", middleware.Autentication(http.HandlerFunc(insertPropietarios))).Methods("POST")
	s.Handle("/info-prop/{numero_documento}", middleware.Autentication(http.HandlerFunc(getOnePropietarioByDocument))).Methods("GET")
	s.Handle("/update/{numero_documento}", middleware.Autentication(http.HandlerFunc(updatePropietario))).Methods("PUT")
	s.Handle("/filter/{filtro}", middleware.Autentication(http.HandlerFunc(propietariosFilter))).Methods("GET")
	s.HandleFunc("/consulta-web/{numero_placa}", consultaWeb).Methods("GET")
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
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("error al leer el cuerpo de la solicitud"))
		return
	}

	data_insert := append([]map[string]interface{}{}, data_request)

	for i := range data_insert {
		data_insert[i]["numero_documento"] = formatString(data_insert[i]["numero_documento"].(string))
	}

	schema, table := tables.Propietarios_GetSchema()
	propietarios := orm.SqlExec{}

	if err := propietarios.New(data_insert, table).Insert(schema); err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	if err := propietarios.Exec(); err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("hubo un error al crear el propietario, por favor intente nuevamente o comuniquese con el administrador"))
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

	if err := propietarios.New(data_update, table).Update(schema); err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	if err := propietarios.Exec(); err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("hubo un error al actualizar el propietario, por favor intente nuevamente o comuniquese con el administrador"))
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

	_data_propietarios := orm.NewQuerys("propietarios").Select().Ilike("numero_documento", "%"+filtro+"%").OrLike("nombre_propietario", "%"+filtro+"%").OrderBy("nombre_propietario").Exec(orm.Config_Query{Cloud: true}).All()

	response.Data["propietarios"] = _data_propietarios
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func consultaWeb(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_placa := params["numero_placa"]

	_data_inscripcion := orm.NewQuerys("inscripciones").Select("id_inscripcion,fecha_fin,fecha_pago,importe,estado,numero_flota,numero_placa,numero_documento").Where("numero_placa", "=", numero_placa).Exec(orm.Config_Query{Cloud: true}).One()
	if len(_data_inscripcion) <= 0 {
		controller.ErrorsWaning(w, errors.New("número de placa no válido o sin suscripción"))
		return
	}
	if _data_inscripcion["fecha_fin"] != nil {
		controller.ErrorsSuccess(w, errors.New("su suscripción esta inactiva"))
		return
	}

	_data_propietario := orm.NewQuerys("propietarios").Select().Where("numero_documento", "=", _data_inscripcion["numero_documento"].(string)).Exec(orm.Config_Query{Cloud: true}).One()

	if len(_data_propietario) <= 0 {
		controller.ErrorsWaning(w, errors.New("o oh!, no pudimos obtener su información"))
		return
	}

	_data_detalle_comprobante := orm.NewQuerys("detalle_comprobantes dc").Select("months || '/' || years as periodo").InnerJoin("comprobante_pago cp", "cp.id_comprobante_pago = dc.id_comprobante_pago").Where("cp.id_inscripcion", "=", _data_inscripcion["id_inscripcion"]).Exec(orm.Config_Query{Cloud: true}).All()

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

	delete(_data_inscripcion, "id_inscripcion")
	delete(_data_inscripcion, "numero_documento")
	delete(_data_inscripcion, "fecha_pago")
	delete(_data_inscripcion, "fecha_fin")

	response.Data["periodo_inscripcion"] = dataPagos
	response.Data["inscripcion"] = _data_inscripcion
	response.Data["propietario"] = _data_propietario
	response.Data["status_pago"] = map[string]interface{}{
		"proximo_pago": new_date_fact,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func formatString(input string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Println("Error al compilar la expresión regular:", err)
		return input
	}

	formatted := reg.ReplaceAllString(input, "")
	return formatted
}
