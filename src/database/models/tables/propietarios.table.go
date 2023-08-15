package tables

import (
	"github.com/deybin/go_basic_orm"
)

func Propietarios_GetSchema() ([]go_basic_orm.Model, string) {
	var propietarios []go_basic_orm.Model
	tableName := "propietarios"
	propietarios = append(propietarios, go_basic_orm.Model{ //numero_documento
		Name:        "numero_documento",
		Description: "numero_documento",
		Important:   true,
		Required:    true,
		Update:      false,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min: 8,
			Max: 11,
		},
	})
	propietarios = append(propietarios, go_basic_orm.Model{ //tipo_documento
		Name:        "tipo_documento",
		Description: "tipo_documento",
		Required:    true,
		Update:      false,
		Type:        "uint64",
		Uint:        go_basic_orm.Uints{},
	})
	propietarios = append(propietarios, go_basic_orm.Model{ //nombre_propietario
		Name:        "nombre_propietario",
		Description: "nombre_propietario",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       5,
			Max:       150,
			LowerCase: true,
		},
	})
	propietarios = append(propietarios, go_basic_orm.Model{ //direccion
		Name:        "direccion",
		Description: "direccion",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       5,
			Max:       150,
			LowerCase: true,
		},
	})
	propietarios = append(propietarios, go_basic_orm.Model{ //telefono
		Name:        "telefono",
		Description: "telefono",
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       6,
			Max:       9,
			LowerCase: true,
		},
	})
	propietarios = append(propietarios, go_basic_orm.Model{ //email
		Name:        "email",
		Description: "email",
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       12,
			Max:       50,
			LowerCase: true,
		},
	})
	propietarios = append(propietarios, go_basic_orm.Model{ //referencia
		Name:        "referencia",
		Description: "referencia",
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       5,
			Max:       150,
			LowerCase: true,
		},
	})
	return propietarios, tableName
}
