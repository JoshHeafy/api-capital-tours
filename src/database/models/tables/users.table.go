package tables

import (
	"github.com/deybin/go_basic_orm"
	"github.com/google/uuid"
)

func Users_GetSchema() ([]go_basic_orm.Model, string) {
	var users []go_basic_orm.Model
	tableName := "users"
	id_user := uuid.New().String()
	users = append(users, go_basic_orm.Model{ //id_user
		Name:        "id_user",
		Description: "id_user",
		Default:     id_user,
		Required:    true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	users = append(users, go_basic_orm.Model{ //cargo
		Name:        "cargo",
		Description: "cargo",
		Update:      false,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 10,
		},
	})
	users = append(users, go_basic_orm.Model{ //username
		Name:        "username",
		Description: "username",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       5,
			Max:       100,
			LowerCase: true,
		},
	})
	users = append(users, go_basic_orm.Model{ //email
		Name:        "email",
		Description: "email",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       10,
			Max:       100,
			LowerCase: true,
		},
	})
	users = append(users, go_basic_orm.Model{ //password
		Name:        "password",
		Description: "password",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       6,
			Max:       200,
			LowerCase: true,
			Encriptar: true,
		},
	})
	users = append(users, go_basic_orm.Model{ //numero_placa
		Name:        "numero_placa",
		Description: "numero_placa",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	return users, tableName
}
