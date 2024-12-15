package tables

import (
	"api-capital-tours/src/database/models"

	"github.com/google/uuid"
)

func UsersMovil_GetSchema() ([]models.Base, string) {
	var users []models.Base
	tableName := "users_" + "mobile"
	users = append(users, models.Base{ //id_user
		Name:        "id_user",
		Description: "ID del usuario",
		Default:     uuid.New().String(),
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings:     models.Strings{},
	})
	users = append(users, models.Base{ //cargo
		Name:        "cargo",
		Description: "Cargo",
		Default:     0,
		Required:    true,
		Update:      false,
		Type:        "int64",
		Int: models.Ints{
			Max: 5,
		},
	})
	users = append(users, models.Base{ //email
		Name:        "email",
		Description: "Email",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min: 10,
			Max: 150,
		},
	})
	users = append(users, models.Base{ //password
		Name:        "password",
		Description: "Contraseña",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Max:       3,
			Encriptar: true,
		},
	})
	users = append(users, models.Base{ //numero_placa
		Name:        "numero_placa",
		Description: "Placa del vehículo",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min:       5,
			Max:       100,
			UpperCase: true,
		},
	})

	return users, tableName
}
