package routes

import (
	"api-capital-tours/src/controller"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/deybin/go_basic_orm"
	"github.com/gorilla/mux"
)

func RutasAnalytics(r *mux.Router) {
	s := r.PathPrefix("/analytics").Subrouter()
	s.HandleFunc("/dashboard", generateAnalitycs).Methods("POST")
}

func generateAnalitycs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	req_body, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	matchDate, _ := regexp.MatchString(`^\d{2}/\d{2}/\d{4}$`, req_body["date"].(string))
	if !matchDate {
		controller.ErrorsWaning(w, errors.New("formato de fecha no válida"))
		return
	}

	date_analytic := req_body["date"].(string)

	_data_inscripciones, _ := new(go_basic_orm.Querys).NewQuerys("inscripciones").Select().Where("fecha_inicio", " = ", date_analytic).Exec(go_basic_orm.Config_Query{Cloud: true}).All()

	_data_comprobantes, _ := new(go_basic_orm.Querys).NewQuerys("comprobante_pago").Select().Where("fecha_pago", " = ", date_analytic).Exec(go_basic_orm.Config_Query{Cloud: true}).All()

	tota_suscripciones := len(_data_inscripciones)
	falt_suscripciones := 20 - tota_suscripciones

	tota_ingresos := 0.0
	for _, pago := range _data_comprobantes {
		tota_ingresos += pago["total"].(float64)
	}
	falt_ingresos := 1000 - tota_ingresos

	tota_pagos := len(_data_comprobantes)
	falt_pagos := 150 - tota_pagos

	if falt_suscripciones < 0 {
		falt_suscripciones = 0
	}
	if falt_ingresos < 0 {
		falt_ingresos = 0
	}
	if falt_pagos < 0 {
		falt_pagos = 0
	}

	analyticsSuscripciones := map[string]interface{}{
		"total":    tota_suscripciones,
		"faltante": falt_suscripciones,
	}

	analyticsIngresos := map[string]interface{}{
		"total":    tota_ingresos,
		"faltante": falt_ingresos,
	}

	analyticsPagos := map[string]interface{}{
		"total":    tota_pagos,
		"faltante": falt_pagos,
	}

	response.Data["suscripciones"] = analyticsSuscripciones
	response.Data["ingresos"] = analyticsIngresos
	response.Data["pagos"] = analyticsPagos
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
