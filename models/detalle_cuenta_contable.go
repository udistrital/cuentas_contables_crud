package models

var DetalleCuentaContableCollection = "detalle_cuenta_contable"

type DetalleCuentaContable struct {
	ID    string `json:"ID" bson:"_id,omitempty"`
	Label string `json:"Label" bson:"label"`
}
