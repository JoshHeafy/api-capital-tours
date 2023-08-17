package routes

import (
	"api-capital-tours/src/controller"
	"api-capital-tours/src/database/models/tables"
	"api-capital-tours/src/libraries/date"
	"api-capital-tours/src/middleware"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/deybin/go_basic_orm"
	"github.com/gorilla/mux"
)

func RutasComprobantes(r *mux.Router) {
	s := r.PathPrefix("/comprobantes").Subrouter()

	s.Handle("/info/{numero_documento}", middleware.Autentication(http.HandlerFunc(getOneComprobante))).Methods("GET")
	s.Handle("/info-detail/{id_comprobante_pago}", middleware.Autentication(http.HandlerFunc(getOneComprobanteDetail))).Methods("GET")
	s.Handle("/create/{id_inscripcion}", middleware.Autentication(http.HandlerFunc(insertComprobante))).Methods("POST")
}

func getOneComprobante(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_documento := params["numero_documento"]

	_data_comprobantes, _ := new(go_basic_orm.Querys).NewQuerys("comprobante_pago").Select().Where("numero_documento", "=", numero_documento).Exec(go_basic_orm.Config_Query{Cloud: true}).All()

	if len(_data_comprobantes) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data["comprobantes_info"] = _data_comprobantes

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getOneComprobanteDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	id_comprobante_pago := params["id_comprobante_pago"]
	if id_comprobante_pago == "" {
		controller.ErrorsError(w, errors.New("comprobante no encontrado"))
		return
	}

	_data_comprobante_detail, _ := new(go_basic_orm.Querys).NewQuerys("detalle_comprobantes").Select().Where("id_comprobante_pago", "=", id_comprobante_pago).Exec(go_basic_orm.Config_Query{Cloud: true}).One()

	if len(_data_comprobante_detail) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data["comprobante_detail"] = _data_comprobante_detail

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertComprobante(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	id_inscripcion := params["id_inscripcion"]

	_data_inscripciones, _ := new(go_basic_orm.Querys).NewQuerys("inscripciones").Select("id_inscripcion, numero_documento").Where("id_inscripcion", "=", id_inscripcion).Exec(go_basic_orm.Config_Query{Cloud: true}).One()

	if len(_data_inscripciones) <= 0 {
		controller.ErrorsWaning(w, errors.New("id de suscripcion inválido"))
		return
	}

	_comprobantes_number, _ := new(go_basic_orm.Querys).NewQuerys("comprobante_pago").Select("numero_comprobante").Exec(go_basic_orm.Config_Query{Cloud: true}).All()

	var numeroComprobante string

	if len(_comprobantes_number) <= 0 {
		numeroComprobante = "0000000001"
	} else {
		var comprobantesList []int
		for _, val := range _comprobantes_number {
			newComprobante, err := strconv.Atoi(val["numero_comprobante"].(string))
			if err != nil {
				fmt.Println("Error al convertir el string a int:", err)
				return
			}
			comprobantesList = append(comprobantesList, newComprobante)
		}

		if len(comprobantesList) != 0 {
			lastNumberString := strconv.Itoa(len(comprobantesList) + 1)
			var zeros string
			for i := 1; i <= 10-len(lastNumberString); i++ {
				zeros = strings.Repeat("0", i)
			}
			numeroComprobante = zeros + lastNumberString
		}
	}

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	data_insert := make(map[string]interface{})
	data_insert_comprobante_detail := make(map[string]interface{})

	for key, value := range data_request {
		data_insert[key] = value
	}

	data_insert["numero_documento"] = _data_inscripciones["numero_documento"]
	data_insert["numero_comprobante"] = numeroComprobante
	data_insert["id_inscripcion"] = _data_inscripciones["id_inscripcion"]
	data_insert["numero_serie"] = string("0001") //por ahora solo una sucursal
	data_insert["fecha_pago"] = date.GetFechaLocationString()
	data_insert["igv"] = float64(0.18 * data_request["importe"].(float64))
	data_insert["total"] = data_request["importe"].(float64) - data_insert["igv"].(float64) - data_request["descuento"].(float64)

	//insert a comprobante_pago
	schema_comprobante, table := tables.Comprobante_GetSchema()
	comprobantes := go_basic_orm.SqlExec{}
	err = comprobantes.New([]map[string]interface{}{data_insert}, table).Insert(schema_comprobante)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = comprobantes.Exec("capital_tours")
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	data_insert_comprobante_detail["id_comprobante_pago"] = comprobantes.Data[0]["id_comprobante_pago"]
	data_insert_comprobante_detail["importe"] = comprobantes.Data[0]["importe"]
	data_insert_comprobante_detail["descuento"] = comprobantes.Data[0]["descuento"]
	data_insert_comprobante_detail["igv"] = comprobantes.Data[0]["igv"]
	data_insert_comprobante_detail["total"] = comprobantes.Data[0]["total"]
	data_insert_comprobante_detail["years"] = data_request["years"]
	data_insert_comprobante_detail["months"] = data_request["months"]

	//insert a detalle_comprobantes
	schema_comprobante_detail, table := tables.DetalleComprobantes_GetSchema()
	comprobantes_detalle := go_basic_orm.SqlExec{}
	err = comprobantes_detalle.New([]map[string]interface{}{data_insert_comprobante_detail}, table).Insert(schema_comprobante_detail)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = comprobantes_detalle.Exec("capital_tours")
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = comprobantes.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
