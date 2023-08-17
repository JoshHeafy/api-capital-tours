package tables

import (
	"api-capital-tours/src/database/models"

	"github.com/google/uuid"
)

func Permisos_GetSchema() ([]models.Base, string) {
	var permisos []models.Base
	tableName := "permisos"
	id_permiso := uuid.New().String()

	permisos = append(permisos, models.Base{ //id_permiso
		Name:        "id_permiso",
		Description: "id_permiso",
		Required:    true,
		Default:     id_permiso,
		Type:        "string",
		Strings:     models.Strings{},
	})
	permisos = append(permisos, models.Base{ //fecha_inicio
		Name:        "fecha_inicio",
		Description: "fecha_inicio",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Date: true,
		},
	})
	permisos = append(permisos, models.Base{ //fecha_fin
		Name:        "fecha_fin",
		Description: "fecha_fin",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Date: true,
		},
	})
	permisos = append(permisos, models.Base{ //id_inscripcion
		Name:        "id_inscripcion",
		Description: "id_inscripcion",
		Required:    true,
		Important:   true,
		Update:      true,
		Type:        "string",
		Strings:     models.Strings{},
	})
	return permisos, tableName
}
