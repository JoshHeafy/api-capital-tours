package tables

import (
	"github.com/deybin/go_basic_orm"
	"github.com/google/uuid"
)

func Solicitudes_GetSchema() ([]go_basic_orm.Model, string) {
	var solicitudes []go_basic_orm.Model
	tableName := "solicitudes"

	solicitudes = append(solicitudes, go_basic_orm.Model{ //id_solicitudes
		Name:        "id_solicitudes",
		Description: "id_solicitudes",
		Default:     uuid.New().String(),
		Important:   true,
		Required:    true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	solicitudes = append(solicitudes, go_basic_orm.Model{ //nombre
		Name:        "nombre",
		Description: "nombre",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       10,
			Max:       200,
			LowerCase: true,
		},
	})
	solicitudes = append(solicitudes, go_basic_orm.Model{ //email
		Name:        "email",
		Description: "email",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min: 12,
			Max: 100,
		},
	})
	solicitudes = append(solicitudes, go_basic_orm.Model{ //telefono
		Name:        "telefono",
		Description: "telefono",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min: 6,
			Max: 9,
		},
	})
	solicitudes = append(solicitudes, go_basic_orm.Model{ //asunto
		Name:        "asunto",
		Description: "asunto",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       5,
			Max:       150,
			LowerCase: true,
		},
	})
	solicitudes = append(solicitudes, go_basic_orm.Model{ //mensaje
		Name:        "mensaje",
		Description: "mensaje",
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       10,
			Max:       500,
			LowerCase: true,
		},
	})
	return solicitudes, tableName
}
