package tables

import (
	"github.com/deybin/go_basic_orm"
)

func DetalleComprobantes_GetSchema() ([]go_basic_orm.Model, string) {
	var detalleComprobantes []go_basic_orm.Model
	tableName := "detalle_" + "comprobantes"
	detalleComprobantes = append(detalleComprobantes, go_basic_orm.Model{ //id_comprobante_pago
		Name:        "id_comprobante_pago",
		Description: "id_comprobante_pago",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	detalleComprobantes = append(detalleComprobantes, go_basic_orm.Model{ //importe
		Name:        "importe",
		Description: "importe",
		Required:    true,
		Type:        "float64",
		Float:       go_basic_orm.Floats{},
	})
	detalleComprobantes = append(detalleComprobantes, go_basic_orm.Model{ //descuento
		Name:        "descuento",
		Description: "descuento",
		Type:        "float64",
		Required:    true,
		Float:       go_basic_orm.Floats{},
	})
	detalleComprobantes = append(detalleComprobantes, go_basic_orm.Model{ //igv
		Name:        "igv",
		Description: "igv",
		Required:    true,
		Type:        "float64",
		Float:       go_basic_orm.Floats{},
	})
	detalleComprobantes = append(detalleComprobantes, go_basic_orm.Model{ //total
		Name:        "total",
		Description: "total",
		Required:    true,
		Type:        "float64",
		Float:       go_basic_orm.Floats{},
	})
	detalleComprobantes = append(detalleComprobantes, go_basic_orm.Model{ //years
		Name:        "years",
		Description: "years",
		Required:    true,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 3000,
		},
	})
	detalleComprobantes = append(detalleComprobantes, go_basic_orm.Model{ //months
		Name:        "months",
		Description: "months",
		Required:    true,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 12,
		},
	})
	return detalleComprobantes, tableName
}
