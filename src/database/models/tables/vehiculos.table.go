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
			Min:       0,
			Max:       7,
			UpperCase: true,
		},
	})
	vehiculos = append(vehiculos, go_basic_orm.Model{ //numero_asientos
		Name:        "numero_asientos",
		Description: "numero_asientos",
		Required:    true,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 10,
		},
	})
	vehiculos = append(vehiculos, go_basic_orm.Model{ //anio
		Name:        "anio",
		Description: "anio",
		Required:    true,
		Type:        "uint64",
		Uint:        go_basic_orm.Uints{},
	})
	vehiculos = append(vehiculos, go_basic_orm.Model{ //numero_pasajeros
		Name:        "numero_pasajeros",
		Description: "numero_pasajeros",
		Required:    true,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 10,
		},
	})
	vehiculos = append(vehiculos, go_basic_orm.Model{ //color
		Name:        "color",
		Description: "color",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       0,
			Max:       7,
			UpperCase: true,
		},
	})
	vehiculos = append(vehiculos, go_basic_orm.Model{ //numero_serie
		Name:        "numero_serie",
		Description: "numero_serie",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       2,
			Max:       20,
			UpperCase: true,
		},
	})
	vehiculos = append(vehiculos, go_basic_orm.Model{ //observaciones
		Name:        "observaciones",
		Description: "observaciones",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       5,
			Max:       100,
			LowerCase: true,
		},
	})
	vehiculos = append(vehiculos, go_basic_orm.Model{ //numero_documento
		Name:        "numero_documento",
		Description: "numero_documento",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Max: 11,
		},
	})
	vehiculos = append(vehiculos, go_basic_orm.Model{ //marca
		Name:        "marca",
		Description: "marca",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       1,
			Max:       20,
			LowerCase: true,
		},
	})
	vehiculos = append(vehiculos, go_basic_orm.Model{ //modelo
		Name:        "modelo",
		Description: "modelo",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Max:       20,
			LowerCase: true,
		},
	})
	return vehiculos, tableName
}
