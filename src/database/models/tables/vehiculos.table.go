package tables

import (
	"github.com/deybin/go_basic_orm"
)

func Vehiculos_GetSchema() ([]go_basic_orm.Model, string) {
	var vehiculos []go_basic_orm.Model
	tableName := "vehiculos"
	vehiculos = append(vehiculos, go_basic_orm.Model{ //numero_placa
		Name:        "numero_placa",
		Description: "numero_placa",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       7,
			Max:       7,
			UpperCase: true,
		},
	})
	vehiculos = append(vehiculos, go_basic_orm.Model{ //marca
		Name:        "marca",
		Description: "marca",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       3,
			Max:       15,
			LowerCase: true,
		},
	})
	vehiculos = append(vehiculos, go_basic_orm.Model{ //modelo
		Name:        "modelo",
		Description: "modelo",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       3,
			Max:       50,
			LowerCase: true,
		},
	})
	vehiculos = append(vehiculos, go_basic_orm.Model{ //anio
		Name:        "anio",
		Description: "anio",
		Required:    true,
		Update:      true,
		Type:        "int64",
		Int: go_basic_orm.Ints{
			Min: 1800,
			Max: 3000,
		},
	})
	vehiculos = append(vehiculos, go_basic_orm.Model{ //color
		Name:        "color",
		Description: "color",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Max:       7,
			LowerCase: true,
		},
	})
	vehiculos = append(vehiculos, go_basic_orm.Model{ //numero_serie
		Name:        "numero_serie",
		Description: "numero_serie",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       17,
			UpperCase: true,
		},
	})
	vehiculos = append(vehiculos, go_basic_orm.Model{ //numero_pasajeros
		Name:        "numero_pasajeros",
		Description: "numero_pasajeros",
		Required:    true,
		Type:        "int64",
		Int: go_basic_orm.Ints{
			Max: 9,
		},
	})
	vehiculos = append(vehiculos, go_basic_orm.Model{ //numero_asientos
		Name:        "numero_asientos",
		Description: "numero_asientos",
		Required:    true,
		Update:      true,
		Type:        "int64",
		Int: go_basic_orm.Ints{
			Max: 9,
		},
	})
	vehiculos = append(vehiculos, go_basic_orm.Model{ //observaciones
		Name:        "observaciones",
		Description: "observaciones",
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Max:       100,
			LowerCase: true,
		},
	})
	vehiculos = append(vehiculos, go_basic_orm.Model{ //numero_documento
		Name:        "numero_documento",
		Description: "numero_documento",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Max: 11,
		},
	})

	return vehiculos, tableName
}
