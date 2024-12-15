package tables

import (
	"api-capital-tours/src/database/models"

	"github.com/google/uuid"
)

func Solicitudes_GetSchema() ([]models.Base, string) {
	var solicitudes []models.Base
	tableName := "solicitudes"

	solicitudes = append(solicitudes, models.Base{ //id_solicitud
		Name:        "id_solicitud",
		Description: "ID de la solicitud",
		Default:     uuid.New().String(),
		Important:   true,
		Required:    true,
		Type:        "string",
		Strings:     models.Strings{},
	})
	solicitudes = append(solicitudes, models.Base{ //nombre
		Name:        "nombre",
		Description: "Nombre",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Min:       10,
			Max:       200,
			LowerCase: true,
		},
	})
	solicitudes = append(solicitudes, models.Base{ //email
		Name:        "email",
		Description: "Email",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Min: 12,
			Max: 100,
		},
	})
	solicitudes = append(solicitudes, models.Base{ //telefono
		Name:        "telefono",
		Description: "Telefono",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Min: 6,
			Max: 9,
		},
	})
	solicitudes = append(solicitudes, models.Base{ //asunto
		Name:        "asunto",
		Description: "Asunto",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Min:       4,
			Max:       150,
			LowerCase: true,
		},
	})
	solicitudes = append(solicitudes, models.Base{ //mensaje
		Name:        "mensaje",
		Description: "Mensaje",
		Type:        "string",
		Strings: models.Strings{
			Min:       10,
			Max:       500,
			LowerCase: true,
		},
	})
	solicitudes = append(solicitudes, models.Base{ //leido
		Name:        "leido",
		Description: "Leido",
		Type:        "uint64",
		Default:     0,
		Update:      true,
	})
	return solicitudes, tableName
}
