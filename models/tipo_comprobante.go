package models

var TipoComprobanteCollection = "tipo_comprobante"

// TipoComprobante es la estructura de un tipo de comprobantel
type TipoComprobante struct {
	*General
	ID            string `json:"Codigo" bson:"_id,omitempty"`
	TipoDocumento string `json:"TipoDocumento" bson:"tipo_documento"`
	Descripcion   string `json:"Descripcion" bson:"descripcion"`
}
