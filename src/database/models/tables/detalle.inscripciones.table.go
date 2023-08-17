package tables

import (
	"api-capital-tours/src/database/models"

	"github.com/google/uuid"
)

func Detalleinscripciones_GetSchema() ([]models.Base, string) {
	var detalleinscripciones []models.Base
	tableName := "detalle_" + "inscripciones"
	detalleinscripciones = append(detalleinscripciones, models.Base{ //id_detalle_inscripcion
		Name:        "id_detalle_inscripcion",
		Description: "id_detalle_inscripcion",
		Default:     uuid.New().String(),
		Required:    true,
		Type:        "string",
		Strings:     models.Strings{},
	})
	detalleinscripciones = append(detalleinscripciones, models.Base{ //fecha_pago
		Name:        "fecha_pago",
		Description: "fecha_pago",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min:  10,
			Date: true,
		},
	})
	detalleinscripciones = append(detalleinscripciones, models.Base{ //years
		Name:        "years",
		Description: "years",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 3000,
		},
	})
	detalleinscripciones = append(detalleinscripciones, models.Base{ //months
		Name:        "months",
		Description: "months",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 12,
		},
	})
	detalleinscripciones = append(detalleinscripciones, models.Base{ //importe
		Name:        "importe",
		Description: "importe",
		Update:      true,
		Type:        "float64",
		Float:       models.Floats{},
	})
	detalleinscripciones = append(detalleinscripciones, models.Base{ //numero_documento
		Name:        "numero_documento",
		Description: "numero_documento",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Max: 11,
			Min: 8,
		},
	})
	detalleinscripciones = append(detalleinscripciones, models.Base{ //estado
		Name:        "estado",
		Description: "estado",
		Default:     1,
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 1,
		},
	})
	detalleinscripciones = append(detalleinscripciones, models.Base{ //id_inscripcion
		Name:        "id_inscripcion",
		Description: "id_inscripcion",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings:     models.Strings{},
	})
	return detalleinscripciones, tableName
}
