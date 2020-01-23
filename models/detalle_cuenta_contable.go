package models

// DetalleCuentaContableCollection ...
var DetalleCuentaContableCollection = "detalle_cuenta_contable"

// DetalleCuentaContable ...
type DetalleCuentaContable struct {
	ID    string `json:"Id" bson:"_id,omitempty"`
	Label string `json:"Label" bson:"label"`
}
