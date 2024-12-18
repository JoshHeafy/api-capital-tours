package tables

import (
	"api-capital-tours/src/database/models"

	"github.com/google/uuid"
)

func UsersAdmin_GetSchema() ([]models.Base, string) {
	var users []models.Base
	tableName := "users_" + "admin"
	users = append(users, models.Base{ //id_user_admin
		Name:        "id_user_admin",
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
		Required:    true,
		Update:      false,
		Type:        "int64",
		Int: models.Ints{
			Max: 5,
		},
	})
	users = append(users, models.Base{ //username
		Name:        "username",
		Description: "Nombre de usuario",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min: 5,
			Max: 100,
		},
	})
	users = append(users, models.Base{ //nombre
		Name:        "nombre",
		Description: "Nombre",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min: 3,
			Max: 100,
		},
	})
	users = append(users, models.Base{ //apellidos
		Name:        "apellidos",
		Description: "Apellidos",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min: 5,
			Max: 100,
		},
	})
	users = append(users, models.Base{ //id_img
		Name:        "id_img",
		Description: "ID de la imagen",
		Update:      true,
		Type:        "string",
		Strings:     models.Strings{},
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
		Description: "Contraseña",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Max:       3,
			Encriptar: true,
		},
	})

	return users, tableName
}
