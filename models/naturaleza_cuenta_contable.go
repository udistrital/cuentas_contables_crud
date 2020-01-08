package models

var NaturalezaCuentaContableCollection = "naturaleza_cuenta_contable"

// NaturalezaCuentaContable cuenta contable parameter.
type NaturalezaCuentaContable struct {
	ID    string `json:"ID" bson:"_id,omitempty"`
	Label string `json:"Label" bson:"label"`
}
