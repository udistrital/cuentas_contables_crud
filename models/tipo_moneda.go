package models

// TipoMonedaCollection ...
var TipoMonedaCollection = "tipo_moneda"

// TipoMoneda ...
type TipoMoneda struct {
	ID    string `json:"Id" bson:"_id,omitempty"`
	Label string `json:"Label" bson:"label"`
}
