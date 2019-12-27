package models

// TipoComprobante es la estructura de un tipo de comprobantel
type TipoComprobante struct {
	ID            string `json:"Codigo" bson:"_id,omitempty"`
	TipoDocumento string `json:"TipoDocumento" bson:"tipodocumento"`
	Descripcion   string `json:"Descripcion" bson:"descripcion"`
}
