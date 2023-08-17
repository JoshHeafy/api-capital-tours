package tables

import (
	"github.com/deybin/go_basic_orm"
	"github.com/google/uuid"
)

func Detalleinscripciones_GetSchema() ([]go_basic_orm.Model, string) {
	var detalleinscripciones []go_basic_orm.Model
	tableName := "detalle_" + "inscripciones"
	detalleinscripciones = append(detalleinscripciones, go_basic_orm.Model{ //id_detalle_inscripcion
		Name:        "id_detalle_inscripcion",
		Description: "id_detalle_inscripcion",
		Default:     uuid.New().String(),
		Required:    true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	detalleinscripciones = append(detalleinscripciones, go_basic_orm.Model{ //fecha_pago
		Name:        "fecha_pago",
		Description: "fecha_pago",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:  10,
			Date: true,
		},
	})
	detalleinscripciones = append(detalleinscripciones, go_basic_orm.Model{ //years
		Name:        "years",
		Description: "years",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 3000,
		},
	})
	detalleinscripciones = append(detalleinscripciones, go_basic_orm.Model{ //months
		Name:        "months",
		Description: "months",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 12,
		},
	})
	detalleinscripciones = append(detalleinscripciones, go_basic_orm.Model{ //importe
		Name:        "importe",
		Description: "importe",
		Update:      true,
		Type:        "float64",
		Float:       go_basic_orm.Floats{},
	})
	detalleinscripciones = append(detalleinscripciones, go_basic_orm.Model{ //numero_documento
		Name:        "numero_documento",
		Description: "numero_documento",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Max: 11,
			Min: 8,
		},
	})
	detalleinscripciones = append(detalleinscripciones, go_basic_orm.Model{ //estado
		Name:        "estado",
		Description: "estado",
		Default:     1,
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 1,
		},
	})
	detalleinscripciones = append(detalleinscripciones, go_basic_orm.Model{ //id_inscripcion
		Name:        "id_inscripcion",
		Description: "id_inscripcion",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	return detalleinscripciones, tableName
}
