package tables

import (
	"api-capital-tours/src/database/models"
)

func Conductores_GetSchema() ([]models.Base, string) {
	var conductores []models.Base
	tableName := "conductores"
	conductores = append(conductores, models.Base{
		Name:        "numero_licencia",
		Description: "Número de licencia",
		Important:   true,
		Required:    true,
		Update:      false,
		Type:        "string",
		Strings: models.Strings{
			Min: 7,
			Max: 10,
		},
	})
	conductores = append(conductores, models.Base{
		Name:        "categoria_licencia",
		Description: "Categoría de licencia",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Max: 6,
		},
	})
	conductores = append(conductores, models.Base{
		Name:        "fecha_caducacion_licencia",
		Description: "Fecha de caducación de licencia",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Date: true,
			Max:  10,
		},
	})
	conductores = append(conductores, models.Base{
		Name:        "fecha_nacimiento",
		Description: "Fecha de nacimiento",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Date: true,
			Max:  10,
		},
	})
	conductores = append(conductores, models.Base{
		Name:        "nombre_conductor",
		Description: "Nombre del conductor",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min:       5,
			Max:       150,
			LowerCase: true,
		},
	})
	conductores = append(conductores, models.Base{
		Name:        "genero",
		Description: "Género",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 2,
		},
	})
	conductores = append(conductores, models.Base{
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
	conductores = append(conductores, models.Base{
		Name:        "telefono",
		Description: "Teléfono",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min:       6,
			Max:       15,
			LowerCase: true,
		},
	})
	conductores = append(conductores, models.Base{
		Name:        "email",
		Description: "Email",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min:       10,
			Max:       150,
			LowerCase: true,
		},
	})
	conductores = append(conductores, models.Base{
		Name:        "estado",
		Description: "Estado",
		Default:     1,
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 1,
		},
	})
	conductores = append(conductores, models.Base{
		Name:        "numero_placa",
		Description: "Número de placa",
		Update:      true,
		Where:       true,
		Type:        "string",
		Strings: models.Strings{
			UpperCase: true,
		},
	})
	return conductores, tableName
}
