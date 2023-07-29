package tables

import (
	"github.com/deybin/go_basic_orm"
	"github.com/google/uuid"
)

func Inscripciones_GetSchema() ([]go_basic_orm.Model, string) {
	var inscripciones []go_basic_orm.Model
	tableName := "inscripciones"
	id_inscripcion := uuid.New().String()
	inscripciones = append(inscripciones, go_basic_orm.Model{ //id_inscripcion
		Name:        "id_inscripcion",
		Description: "id_inscripcion",
		Required:    true,
		Default:     id_inscripcion,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Max: 36,
		},
	})
	inscripciones = append(inscripciones, go_basic_orm.Model{ //numero_documento
		Name:        "numero_documento",
		Description: "numero_documento",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Max: 11,
			Min: 8,
		},
	})
	inscripciones = append(inscripciones, go_basic_orm.Model{ //fecha_inicio
		Name:        "fecha_inicio",
		Description: "fecha_inicio",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:  10,
			Date: true,
		},
	})
	inscripciones = append(inscripciones, go_basic_orm.Model{ //importe
		Name:        "importe",
		Description: "importe",
		Update:      true,
		Type:        "float64",
		Float:       go_basic_orm.Floats{},
	})
	inscripciones = append(inscripciones, go_basic_orm.Model{ //fecha_pago
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
	inscripciones = append(inscripciones, go_basic_orm.Model{ //years
		Name:        "years",
		Description: "years",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 3000,
		},
	})
	inscripciones = append(inscripciones, go_basic_orm.Model{ //months
		Name:        "months",
		Description: "months",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 12,
		},
	})
	inscripciones = append(inscripciones, go_basic_orm.Model{ //estado
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
	inscripciones = append(inscripciones, go_basic_orm.Model{ //fecha_fin
		Name:        "fecha_fin",
		Description: "fecha_fin",
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Max:  10,
			Date: true,
		},
	})
	inscripciones = append(inscripciones, go_basic_orm.Model{ //numero_flota
		Name:        "numero_flota",
		Description: "numero_flota",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint:        go_basic_orm.Uints{},
	})
	inscripciones = append(inscripciones, go_basic_orm.Model{ //numero_placa
		Name:        "numero_placa",
		Description: "numero_placa",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			UpperCase: true,
		},
	})
	return inscripciones, tableName
}
