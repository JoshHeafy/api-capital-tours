package tables

import (
	"github.com/deybin/go_basic_orm"
	"github.com/google/uuid"
)

func Comprobante_GetSchema() ([]go_basic_orm.Model, string) {
	var comprobante []go_basic_orm.Model
	tableName := "comprobante_" + "pago"
	comprobante = append(comprobante, go_basic_orm.Model{ //id_comprobante_comprobante
		Name:        "id_comprobante_pago",
		Description: "id_comprobante_pago",
		Important:   true,
		Required:    true,
		Default:     uuid.New().String(),
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	comprobante = append(comprobante, go_basic_orm.Model{ //numero_documento
		Name:        "numero_documento",
		Description: "numero_documento",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       1,
			Max:       11,
			LowerCase: true,
		},
	})
	comprobante = append(comprobante, go_basic_orm.Model{ //tipo
		Name:        "tipo",
		Description: "tipo",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       0,
			Max:       2,
			LowerCase: true,
		},
	})
	comprobante = append(comprobante, go_basic_orm.Model{ //numero_serie
		Name:        "numero_serie",
		Description: "numero_serie",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       0,
			Max:       4,
			LowerCase: true,
		},
	})
	comprobante = append(comprobante, go_basic_orm.Model{ //numero_comprobante
		Name:        "numero_comprobante",
		Description: "numero_comprobante",
		Required:    true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	comprobante = append(comprobante, go_basic_orm.Model{ //fecha_pago
		Name:        "fecha_pago",
		Description: "fecha_pago",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Date: true,
		},
	})
	comprobante = append(comprobante, go_basic_orm.Model{ //importe
		Name:        "importe",
		Description: "importe",
		Required:    true,
		Type:        "float64",

		Float: go_basic_orm.Floats{},
	})
	comprobante = append(comprobante, go_basic_orm.Model{ //igv
		Name:        "igv",
		Description: "igv",
		Required:    true,
		Update:      true,
		Type:        "float64",
		Float:       go_basic_orm.Floats{},
	})
	comprobante = append(comprobante, go_basic_orm.Model{ //descuento
		Name:        "descuento",
		Description: "descuento",
		Type:        "float64",

		Float: go_basic_orm.Floats{},
	})
	comprobante = append(comprobante, go_basic_orm.Model{ //total
		Name:        "total",
		Description: "total",
		Required:    true,
		Type:        "float64",
		Float:       go_basic_orm.Floats{},
	})
	comprobante = append(comprobante, go_basic_orm.Model{ //observaciones
		Name:        "observaciones",
		Description: "observaciones",
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Max:       100,
			LowerCase: true,
		},
	})
	comprobante = append(comprobante, go_basic_orm.Model{ //estado
		Name:        "estado",
		Description: "estado",
		Default:     1,
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 10,
		},
	})
	comprobante = append(comprobante, go_basic_orm.Model{ //id_inscripcion
		Name:        "id_inscripcion",
		Description: "id_inscripcion",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})

	return comprobante, tableName
}
