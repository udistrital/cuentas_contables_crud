package models

// TipoCuentaCollection ...
var TipoCuentaCollection = "tipo_cuenta"

// TipoCuenta es la estructura de un tipo de cuenta
type TipoCuenta struct {
	ID    string `json:"Id" bson:"_id,omitempty"`
	Label string `json:"Label" bson:"label"`
}
