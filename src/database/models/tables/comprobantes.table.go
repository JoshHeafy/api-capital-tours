package tables

import (
	"api-capital-tours/src/database/models"

	"github.com/google/uuid"
)

func Comprobante_GetSchema() ([]models.Base, string) {
	var comprobante []models.Base
	tableName := "comprobante_" + "pago"
	comprobante = append(comprobante, models.Base{ //id_comprobante_comprobante
		Name:        "id_comprobante_pago",
		Description: "ID del comprobante de pago",
		Important:   true,
		Required:    true,
		Default:     uuid.New().String(),
		Type:        "string",
		Strings:     models.Strings{},
	})
	comprobante = append(comprobante, models.Base{ //numero_documento
		Name:        "numero_documento",
		Description: "Número de documento",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Min:       1,
			Max:       11,
			LowerCase: true,
		},
	})
	comprobante = append(comprobante, models.Base{ //tipo
		Name:        "tipo",
		Description: "Tipo de comprobante",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min:       0,
			Max:       2,
			LowerCase: true,
		},
	})
	comprobante = append(comprobante, models.Base{ //numero_serie
		Name:        "numero_serie",
		Description: "Numero de serie",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Min:       0,
			Max:       4,
			LowerCase: true,
		},
	})
	comprobante = append(comprobante, models.Base{ //numero_comprobante
		Name:        "numero_comprobante",
		Description: "Numero de comprobante",
		Required:    true,
		Type:        "string",
		Strings:     models.Strings{},
	})
	comprobante = append(comprobante, models.Base{ //fecha_pago
		Name:        "fecha_pago",
		Description: "fecha_pago",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Date: true,
		},
	})
	comprobante = append(comprobante, models.Base{ //importe
		Name:        "importe",
		Description: "Importe",
		Required:    true,
		Type:        "float64",
		Float:       models.Floats{},
	})
	comprobante = append(comprobante, models.Base{ //igv
		Name:        "igv",
		Description: "IGV",
		Required:    true,
		Update:      true,
		Type:        "float64",
		Float:       models.Floats{},
	})
	comprobante = append(comprobante, models.Base{ //descuento
		Name:        "descuento",
		Description: "Descuento",
		Type:        "float64",
		Float:       models.Floats{},
	})
	comprobante = append(comprobante, models.Base{ //total
		Name:        "total",
		Description: "Total",
		Required:    true,
		Type:        "float64",
		Float:       models.Floats{},
	})
	comprobante = append(comprobante, models.Base{ //observaciones
		Name:        "observaciones",
		Description: "Observaciones",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Max:       100,
			LowerCase: true,
		},
	})
	comprobante = append(comprobante, models.Base{ //estado
		Name:        "estado",
		Description: "Estado",
		Default:     1,
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	comprobante = append(comprobante, models.Base{ //id_inscripcion
		Name:        "id_inscripcion",
		Description: "ID de la inscripcion",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings:     models.Strings{},
	})

	return comprobante, tableName
}
