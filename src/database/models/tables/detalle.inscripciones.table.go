package tables

import (
	"github.com/deybin/go_basic_orm"
)

func Detalleinscripciones_GetSchema() ([]go_basic_orm.Model, string) {
	var detalleinscripciones []go_basic_orm.Model
	tableName := "detalle_" + "inscripciones"
	detalleinscripciones = append(detalleinscripciones, go_basic_orm.Model{ //iddetalleinscripcion
		Name:        "iddetalleinscripcion",
		Description: "iddetalleinscripcion",
		Required:    true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	detalleinscripciones = append(detalleinscripciones, go_basic_orm.Model{ //idinscripcion
		Name:        "idinscripcion",
		Description: "idinscripcion",
		Important:   true,
		Required:    true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	detalleinscripciones = append(detalleinscripciones, go_basic_orm.Model{ //importe
		Name:        "importe",
		Description: "importe",
		Required:    true,
		Update:      true,
	})
	detalleinscripciones = append(detalleinscripciones, go_basic_orm.Model{ //estado
		Name:        "estado",
		Description: "estado",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 10,
		},
	})
	detalleinscripciones = append(detalleinscripciones, go_basic_orm.Model{ //years
		Name:        "years",
		Description: "years",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       2,
			Max:       2,
			LowerCase: true,
		},
	})
	detalleinscripciones = append(detalleinscripciones, go_basic_orm.Model{ //months
		Name:        "months",
		Description: "months",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       2,
			Max:       2,
			LowerCase: true,
		},
	})
	detalleinscripciones = append(detalleinscripciones, go_basic_orm.Model{ //fechapago
		Name:        "fechapago",
		Description: "fechapago",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Date: true,
		},
	})
	return detalleinscripciones, tableName
}
