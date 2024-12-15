package tables

import (
	"api-capital-tours/src/database/models"
)

func Vehiculos_GetSchema() ([]models.Base, string) {
	var vehiculos []models.Base
	tableName := "vehiculos"
	vehiculos = append(vehiculos, models.Base{ //numero_placa
		Name:        "numero_placa",
		Description: "Numero de placa",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings: models.Strings{
			Min:       7,
			Max:       7,
			UpperCase: true,
		},
	})
	vehiculos = append(vehiculos, models.Base{ //marca
		Name:        "marca",
		Description: "Marca",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min:       3,
			Max:       15,
			LowerCase: true,
		},
	})
	vehiculos = append(vehiculos, models.Base{ //modelo
		Name:        "modelo",
		Description: "Modelo",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Min:       3,
			Max:       50,
			LowerCase: true,
		},
	})
	vehiculos = append(vehiculos, models.Base{ //anio
		Name:        "anio",
		Description: "Año",
		Required:    true,
		Update:      true,
		Type:        "int64",
		Int: models.Ints{
			Min: 1800,
			Max: 3000,
		},
	})
	vehiculos = append(vehiculos, models.Base{ //color
		Name:        "color",
		Description: "Color",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Max:       7,
			LowerCase: true,
		},
	})
	vehiculos = append(vehiculos, models.Base{ //numero_serie
		Name:        "numero_serie",
		Description: "Número de serie",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Min:       17,
			UpperCase: true,
		},
	})
	vehiculos = append(vehiculos, models.Base{ //numero_pasajeros
		Name:        "numero_pasajeros",
		Description: "Número de pasajeros",
		Required:    true,
		Type:        "int64",
		Int: models.Ints{
			Max: 9,
		},
	})
	vehiculos = append(vehiculos, models.Base{ //numero_asientos
		Name:        "numero_asientos",
		Description: "Número de asientos",
		Required:    true,
		Update:      true,
		Type:        "int64",
		Int: models.Ints{
			Max: 9,
		},
	})
	vehiculos = append(vehiculos, models.Base{ //observaciones
		Name:        "observaciones",
		Description: "Observaciones",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Max:       100,
			LowerCase: true,
		},
	})
	vehiculos = append(vehiculos, models.Base{ //numero_documento
		Name:        "numero_documento",
		Description: "Número de documento",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Max: 11,
		},
	})

	return vehiculos, tableName
}
