package routes

import (
	"api-capital-tours/src/controller"
	"api-capital-tours/src/libraries/date"
	"api-capital-tours/src/libraries/library"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deybin/go_basic_orm"
	"github.com/gorilla/mux"
)

func RutasConsultas(r *mux.Router) {
	s := r.PathPrefix("/propietarios").Subrouter()

	s.HandleFunc("/list/", verPropietarios).Methods("GET")
	s.HandleFunc("/info/{numeroDocumento}", consultaDocumento).Methods("GET")
	s.HandleFunc("/info-car/{numeroPlaca}", consultaPlaca).Methods("GET")
	s.HandleFunc("/info-car-servicio/{numeroDocumento}", consultaPeriodo).Methods("GET")
}

func verPropietarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := controller.NewResponseManager()

	dataPropietarios, _ := new(go_basic_orm.Querys).NewQuerys("inscripciones").Select().Exec(go_basic_orm.Config_Query{Cloud: true}).All()
	if len(dataPropietarios) <= 0 {
		response.Msg = "No se encontro ningun propietario"
		response.StatusCode = 300
		response.Status = "Error"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data["propietarios"] = dataPropietarios
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func consultaPeriodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numeroDocumento := params["numeroDocumento"]

	if numeroDocumento == "" {
		response.Msg = "Error: no se encontró el documento del cliente"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	dataInscripciones, _ := new(go_basic_orm.Querys).NewQuerys("inscripciones").Select("years || '-' || months as periodo").Where("numero_documento", "=", numeroDocumento).Exec(go_basic_orm.Config_Query{Cloud: true}).All()

	var newdataDetalleInscripcion []string
	for _, v := range dataInscripciones {
		if v["periodo"] != nil {
			newdataDetalleInscripcion = append(newdataDetalleInscripcion, v["periodo"].(string))
		}
	}

	dataDetalleInscripcion, _ := new(go_basic_orm.Querys).NewQuerys("detalle_inscripciones").Select("fecha_pago", "importe").Where("numero_documento", "=", numeroDocumento).And("estado", "=", 0).Exec(go_basic_orm.Config_Query{Cloud: true}).One()

	datePago := date.GetDate(dataDetalleInscripcion["fecha_pago"].(string))

	dateNow := date.GetDateLocation()
	monthInit := int64(datePago.Month())
	yearInit := datePago.Year()
	monthNow := int64(dateNow.Month())
	yearNow := dateNow.Year()

	var dataPagos []map[string]interface{}
	var month = int64(12)

	importe := dataDetalleInscripcion["importe"]

	for i := yearInit; i <= yearNow; i++ {
		if i == yearNow {
			month = monthNow
		}
		for e := monthInit; e <= month; e++ {
			year := fmt.Sprintf("%v", i)
			month := fmt.Sprintf("%02d", e)
			if library.IndexOf_String(newdataDetalleInscripcion, year+"-"+month) == -1 {
				dataPagos = append(dataPagos, map[string]interface{}{
					"years":   year,
					"months":  month,
					"importe": importe,
				})
			}
		}

		monthInit = 1
	}

	response.Data["detalleInscripcion"] = dataPagos
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func consultaDocumento(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numeroDocumento := params["numeroDocumento"]

	if numeroDocumento == "" {
		response.Msg = "Error no se encontro el documento del cliente"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	data_inscripcion, _ := new(go_basic_orm.Querys).NewQuerys("inscripciones").Select().Where("numero_Documento", "=", numeroDocumento).Exec(go_basic_orm.Config_Query{Cloud: true}).All()

	response.Data["inscripciones"] = data_inscripcion
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func consultaPlaca(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numeroPlaca := params["numeroPlaca"]

	if numeroPlaca == "" {
		response.Msg = "Error: No se proporcionó un número de placa"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	data_inscripciones, _ := new(go_basic_orm.Querys).NewQuerys("inscripciones").Select().Where("numero_placa", "=", numeroPlaca).Exec(go_basic_orm.Config_Query{Cloud: true}).One()

	if len(data_inscripciones) == 0 {
		response.Msg = "No se encontraron inscripciones relacionadas al número de placa proporcionado"
		response.StatusCode = 404
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data["inscripciones"] = data_inscripciones
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
