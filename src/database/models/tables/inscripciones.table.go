package tables

import (
	"api-capital-tours/src/database/models"

	"github.com/google/uuid"
)

func Inscripciones_GetSchema() ([]models.Base, string) {
	var inscripciones []models.Base
	tableName := "inscripciones"
	id_inscripcion := uuid.New().String()
	inscripciones = append(inscripciones, models.Base{ //id_inscripcion
		Name:        "id_inscripcion",
		Description: "id_inscripcion",
		Required:    true,
		Default:     id_inscripcion,
		Type:        "string",
		Strings: models.Strings{
			Max: 36,
		},
	})
	inscripciones = append(inscripciones, models.Base{ //numero_documento
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
	inscripciones = append(inscripciones, models.Base{ //fecha_inicio
		Name:        "fecha_inicio",
		Description: "fecha_inicio",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min:  10,
			Date: true,
		},
	})
	inscripciones = append(inscripciones, models.Base{ //importe
		Name:        "importe",
		Description: "importe",
		Update:      true,
		Type:        "float64",
		Float: models.Floats{
			Menor: 25,
		},
	})
	inscripciones = append(inscripciones, models.Base{ //fecha_pago
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
	inscripciones = append(inscripciones, models.Base{ //years
		Name:        "years",
		Description: "years",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 3000,
		},
	})
	inscripciones = append(inscripciones, models.Base{ //months
		Name:        "months",
		Description: "months",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 12,
		},
	})
	inscripciones = append(inscripciones, models.Base{ //estado
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
	inscripciones = append(inscripciones, models.Base{ //fecha_fin
		Name:        "fecha_fin",
		Description: "fecha_fin",
		Update:      true,
		Empty:       true,
		Type:        "string",
		Strings: models.Strings{
			Max: 10,
		},
	})
	inscripciones = append(inscripciones, models.Base{ //numero_flota
		Name:        "numero_flota",
		Description: "numero_flota",
		Required:    true,
		Update:      true,
		Type:        "int64",
		Int: models.Ints{
			Min: 1,
		},
	})
	inscripciones = append(inscripciones, models.Base{ //numero_placa
		Name:        "numero_placa",
		Description: "numero_placa",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings: models.Strings{
			UpperCase: true,
		},
	})
	return inscripciones, tableName
}
