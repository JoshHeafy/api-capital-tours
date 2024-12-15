package tables

import "api-capital-tours/src/database/models"

func DetalleComprobantes_GetSchema() ([]models.Base, string) {
	var detalleComprobantes []models.Base
	tableName := "detalle_" + "comprobantes"
	detalleComprobantes = append(detalleComprobantes, models.Base{ //id_comprobante_pago
		Name:        "id_comprobante_pago",
		Description: "ID del detalle de comprobante de pago",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings:     models.Strings{},
	})
	detalleComprobantes = append(detalleComprobantes, models.Base{ //importe
		Name:        "importe",
		Description: "Importe",
		Required:    true,
		Type:        "float64",
		Float:       models.Floats{},
	})
	detalleComprobantes = append(detalleComprobantes, models.Base{ //descuento
		Name:        "descuento",
		Description: "Descuento",
		Type:        "float64",
		Required:    true,
		Float:       models.Floats{},
	})
	detalleComprobantes = append(detalleComprobantes, models.Base{ //igv
		Name:        "igv",
		Description: "IGV",
		Required:    true,
		Type:        "float64",
		Float:       models.Floats{},
	})
	detalleComprobantes = append(detalleComprobantes, models.Base{ //total
		Name:        "total",
		Description: "Total",
		Required:    true,
		Type:        "float64",
		Float:       models.Floats{},
	})
	detalleComprobantes = append(detalleComprobantes, models.Base{ //years
		Name:        "years",
		Description: "Año",
		Required:    true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 3000,
		},
	})
	detalleComprobantes = append(detalleComprobantes, models.Base{ //months
		Name:        "months",
		Description: "Mes",
		Required:    true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 12,
		},
	})
	return detalleComprobantes, tableName
}
