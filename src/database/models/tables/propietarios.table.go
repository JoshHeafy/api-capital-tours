package tables

import (
	"api-capital-tours/src/database/models"
)

func Propietarios_GetSchema() ([]models.Base, string) {
	var propietarios []models.Base
	tableName := "propietarios"
	propietarios = append(propietarios, models.Base{ //numero_documento
		Name:        "numero_documento",
		Description: "numero_documento",
		Important:   true,
		Required:    true,
		Update:      false,
		Type:        "string",
		Strings: models.Strings{
			Min: 8,
			Max: 11,
		},
	})
	propietarios = append(propietarios, models.Base{ //tipo_documento
		Name:        "tipo_documento",
		Description: "tipo_documento",
		Required:    true,
		Update:      false,
		Type:        "uint64",
		Uint:        models.Uints{},
	})
	propietarios = append(propietarios, models.Base{ //nombre_propietario
		Name:        "nombre_propietario",
		Description: "nombre_propietario",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min:       5,
			Max:       150,
			LowerCase: true,
		},
	})
	propietarios = append(propietarios, models.Base{ //direccion
		Name:        "direccion",
		Description: "direccion",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min:       5,
			Max:       150,
			LowerCase: true,
		},
	})
	propietarios = append(propietarios, models.Base{ //telefono
		Name:        "telefono",
		Description: "telefono",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min:       6,
			Max:       9,
			LowerCase: true,
		},
	})
	propietarios = append(propietarios, models.Base{ //email
		Name:        "email",
		Description: "email",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min:       12,
			Max:       50,
			LowerCase: true,
		},
	})
	propietarios = append(propietarios, models.Base{ //referencia
		Name:        "referencia",
		Description: "referencia",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min:       5,
			Max:       150,
			LowerCase: true,
		},
	})
	return propietarios, tableName
}
