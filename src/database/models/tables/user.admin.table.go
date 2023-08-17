package tables

import (
	"github.com/deybin/go_basic_orm"
	"github.com/google/uuid"
)

func UserAdmin_GetSchema() ([]go_basic_orm.Model, string) {
	var users []go_basic_orm.Model
	tableName := "users_" + "admin"
	users = append(users, go_basic_orm.Model{ //id_user_admin
		Name:        "id_user_admin",
		Description: "id_user_admin",
		Default:     uuid.New().String(),
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	users = append(users, go_basic_orm.Model{ //cargo
		Name:        "cargo",
		Description: "cargo",
		Required:    true,
		Update:      false,
		Type:        "int64",
		Int: go_basic_orm.Ints{
			Max: 5,
		},
	})
	users = append(users, go_basic_orm.Model{ //username
		Name:        "username",
		Description: "username",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min: 5,
			Max: 100,
		},
	})
	users = append(users, go_basic_orm.Model{ //nombre
		Name:        "nombre",
		Description: "nombre",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min: 3,
			Max: 100,
		},
	})
	users = append(users, go_basic_orm.Model{ //apellidos
		Name:        "apellidos",
		Description: "apellidos",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min: 5,
			Max: 100,
		},
	})
	users = append(users, go_basic_orm.Model{ //id_img
		Name:        "id_img",
		Description: "id_img",
		Update:      true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	users = append(users, go_basic_orm.Model{ //email
		Name:        "email",
		Description: "email",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min: 10,
			Max: 150,
		},
	})
	users = append(users, go_basic_orm.Model{ //password
		Name:        "password",
		Description: "password",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Max:       3,
			Encriptar: true,
		},
	})

	return users, tableName
}
