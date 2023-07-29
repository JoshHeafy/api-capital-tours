package tables

import (
	"github.com/deybin/go_basic_orm"
	"github.com/google/uuid"
)

func Permisos_GetSchema() ([]go_basic_orm.Model, string) {
	var permisos []go_basic_orm.Model
	tableName := "permisos"
	id_permiso := uuid.New().String()

	permisos = append(permisos, go_basic_orm.Model{ //id_permiso
		Name:        "id_permiso",
		Description: "id_permiso",
		Required:    true,
		Default:     id_permiso,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	permisos = append(permisos, go_basic_orm.Model{ //fecha_inicio
		Name:        "fecha_inicio",
		Description: "fecha_inicio",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Date: true,
		},
	})
	permisos = append(permisos, go_basic_orm.Model{ //fecha_fin
		Name:        "fecha_fin",
		Description: "fecha_fin",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Date: true,
		},
	})
	permisos = append(permisos, go_basic_orm.Model{ //id_inscripcion
		Name:        "id_inscripcion",
		Description: "id_inscripcion",
		Required:    true,
		Important:   true,
		Update:      true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	return permisos, tableName
}
