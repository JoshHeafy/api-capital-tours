package tables

import "api-capital-tours/src/database/models"

func DetalleComprobantes_GetSchema() ([]models.Base, string) {
	var detalleComprobantes []models.Base
	tableName := "detalle_" + "comprobantes"
	detalleComprobantes = append(detalleComprobantes, models.Base{ //id_comprobante_pago
		Name:        "id_comprobante_pago",
		Description: "id_comprobante_pago",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings:     models.Strings{},
	})
	detalleComprobantes = append(detalleComprobantes, models.Base{ //importe
		Name:        "importe",
		Description: "importe",
		Required:    true,
		Type:        "float64",
		Float:       models.Floats{},
	})
	detalleComprobantes = append(detalleComprobantes, models.Base{ //descuento
		Name:        "descuento",
		Description: "descuento",
		Type:        "float64",
		Required:    true,
		Float:       models.Floats{},
	})
	detalleComprobantes = append(detalleComprobantes, models.Base{ //igv
		Name:        "igv",
		Description: "igv",
		Required:    true,
		Type:        "float64",
		Float:       models.Floats{},
	})
	detalleComprobantes = append(detalleComprobantes, models.Base{ //total
		Name:        "total",
		Description: "total",
		Required:    true,
		Type:        "float64",
		Float:       models.Floats{},
	})
	detalleComprobantes = append(detalleComprobantes, models.Base{ //years
		Name:        "years",
		Description: "years",
		Required:    true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 3000,
		},
	})
	detalleComprobantes = append(detalleComprobantes, models.Base{ //months
		Name:        "months",
		Description: "months",
		Required:    true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 12,
		},
	})
	return detalleComprobantes, tableName
}
