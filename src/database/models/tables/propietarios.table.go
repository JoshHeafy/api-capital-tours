package tables

import (
	"api-capital-tours/src/database/models"
)

func Propietarios_GetSchema() ([]models.Base, string) {
	var propietarios []models.Base
	tableName := "propietarios"
	propietarios = append(propietarios, models.Base{ //numero_documento
		Name:        "numero_documento",
		Description: "Número de documento",
		Important:   true,
		Required:    true,
		Update:      false,
		Type:        "string",
		Strings: models.Strings{
			Min: 7,
			Max: 20,
		},
	})
	propietarios = append(propietarios, models.Base{ //tipo_documento
		Name:        "tipo_documento",
		Description: "Tipo de documento",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint:        models.Uints{},
	})
	propietarios = append(propietarios, models.Base{ //nombre_propietario
		Name:        "nombre_propietario",
		Description: "Nombre del propietario",
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
		Description: "Dirección",
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
		Description: "Teléfono",
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
		Description: "Email",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min:       10,
			Max:       50,
			LowerCase: true,
		},
	})
	propietarios = append(propietarios, models.Base{ //referencia
		Name:        "referencia",
		Description: "Referencia",
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
