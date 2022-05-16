package models

// TipoRetencionCollection ...
var TipoRetencionCollection = "tipo_retencion"

// TipoRetencion ...
type TipoRetencion struct {
	ID    string `json:"Id" bson:"_id,omitempty"`
	Label string `json:"Label" bson:"label"`
}
