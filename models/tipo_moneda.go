package models

var TipoMonedaCollection = "tipo_moneda"

type TipoMoneda struct {
	ID    string `json:"Id" bson:"_id,omitempty"`
	Label string `json:"Label" bson:"label"`
}
