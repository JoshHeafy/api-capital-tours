package routes

import (
	"api-capital-tours/src/controller"
	"api-capital-tours/src/database/models/tables"
	"api-capital-tours/src/database/orm"
	"api-capital-tours/src/libraries/date"
	"api-capital-tours/src/middleware"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func RutasComprobantes(r *mux.Router) {
	s := r.PathPrefix("/comprobantes").Subrouter()

	s.Handle("/info/{numero_documento}", middleware.Autentication(http.HandlerFunc(getOneComprobante))).Methods("GET")
	s.Handle("/info-detail/{id_comprobante_pago}", middleware.Autentication(http.HandlerFunc(getOneComprobanteDetail))).Methods("GET")
	s.Handle("/info-to-admin/{numero_flota}", middleware.Autentication(http.HandlerFunc(getComprobanteAdmin))).Methods("GET")
	s.Handle("/create/{id_inscripcion}", middleware.Autentication(http.HandlerFunc(insertComprobante))).Methods("POST")
}

func getOneComprobante(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_documento := params["numero_documento"]

	_data_comprobantes := orm.NewQuerys("comprobante_pago").Select().Where("numero_documento", "=", numero_documento).Exec(orm.Config_Query{Cloud: true}).All()

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

	_data_comprobante_detail := orm.NewQuerys("detalle_comprobantes").Select().Where("id_comprobante_pago", "=", id_comprobante_pago).Exec(orm.Config_Query{Cloud: true}).One()

	if len(_data_comprobante_detail) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data["comprobante_detail"] = _data_comprobante_detail

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getComprobanteAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	numero_flota := params["numero_flota"]

	_data_inscripcion := orm.NewQuerys("inscripciones").Select().Where("numero_flota", "=", numero_flota).Exec(orm.Config_Query{Cloud: true}).One()

	if len(_data_inscripcion) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	_data_comprobantes := orm.NewQuerys("comprobante_pago").Select().Where("numero_documento", " ilike", _data_inscripcion["numero_documento"]).Exec(orm.Config_Query{Cloud: true}).All()

	_data_propietario := orm.NewQuerys("propietarios").Select("nombre_propietario").Where("numero_documento", "=", _data_inscripcion["numero_documento"]).Exec(orm.Config_Query{Cloud: true}).One()

	if len(_data_comprobantes) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	var detalleComprobante []map[string]interface{}

	for _, comp := range _data_comprobantes {
		_detalle_comprobante := orm.NewQuerys("detalle_comprobantes").Select("months, years").Where("id_comprobante_pago", "=", comp["id_comprobante_pago"]).Exec(orm.Config_Query{Cloud: true}).One()

		periodo := fmt.Sprintf("%d-%d", _detalle_comprobante["months"], _detalle_comprobante["years"])

		comprobante := map[string]interface{}{
			"numero_serie":       "0001", //hardcode
			"numero_comprobante": comp["numero_comprobante"],
			"fecha":              comp["fecha_pago"],
			"numero_documento":   _data_inscripcion["numero_documento"],
			"cliente":            _data_propietario["nombre_propietario"],
			"numero_placa":       _data_inscripcion["numero_placa"],
			"numero_flota":       _data_inscripcion["numero_flota"],
			"periodo":            periodo,
			"importe":            comp["importe"],
			"descuento":          comp["descuento"],
			"total":              comp["total"],
		}
		detalleComprobante = append(detalleComprobante, comprobante)
	}

	response.Data["comprobantes_detail"] = detalleComprobante

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertComprobante(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	id_inscripcion := params["id_inscripcion"]

	_data_inscripciones := orm.NewQuerys("inscripciones").Select("id_inscripcion, numero_documento").Where("id_inscripcion", "=", id_inscripcion).Exec(orm.Config_Query{Cloud: true}).One()

	if len(_data_inscripciones) <= 0 {
		controller.ErrorsWaning(w, errors.New("id de suscripcion inválido"))
		return
	}

	_comprobantes_number := orm.NewQuerys("comprobante_pago").Select("numero_comprobante").Exec(orm.Config_Query{Cloud: true}).All()

	var numeroComprobante string

	if len(_comprobantes_number) <= 0 {
		numeroComprobante = "0000000001"
	} else {
		var comprobantesList []int
		for _, val := range _comprobantes_number {
			newComprobante, err := strconv.Atoi(val["numero_comprobante"].(string))
			if err != nil {
				log.Println("Error al convertir el string a int:", err)
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
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("error al leer el cuerpo de la solicitud"))
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
	comprobantes := orm.SqlExec{}
	err = comprobantes.New([]map[string]interface{}{data_insert}, table).Insert(schema_comprobante)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	if err := comprobantes.Exec(); err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("hubo un error al crear el comprobante, por favor intente nuevamente o comuniquese con el administrador"))
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
	comprobantes_detalle := orm.SqlExec{}
	err = comprobantes_detalle.New([]map[string]interface{}{data_insert_comprobante_detail}, table).Insert(schema_comprobante_detail)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	if err := comprobantes_detalle.Exec(); err != nil {
		log.Println(err)
		controller.ErrorsWaning(w, errors.New("hubo un error al crear el detalle del comprobante, por favor intente nuevamente o comuniquese con el administrador"))
		return
	}

	response.Data = comprobantes.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
