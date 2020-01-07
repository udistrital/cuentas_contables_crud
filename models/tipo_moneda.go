package models

var TipoMonedaCollection = "tipo_moneda"

type TipoMoneda struct {
	ID    string `json:"ID" bson:"_id,omitempty"`
	Label string `json:"Label" bson:"label"`
}
