package models

// ComprobanteCollection ...
var ComprobanteCollection = "comprobante"

// Comprobante es la estructura de un comprobante con parametros opcionales
type Comprobante struct {
	ID              string           `json:"_id" bson:"_id,omitempty"`
	Codigo          int              `json:"Codigo" bson:"codigo"`
	Descripcion     string           `json:"Descripcion" bson:"descripcion"`
	Comprobante     string           `json:"Comprobante" bson:"comprobante"`
	Numero          int              `json:"Numero" bson:"numero"`
	TipoComprobante *TipoComprobante `json:"TipoComprobante" bson:"tipocomprobante"`
	*General
	*Parametros
}

// Parametros ...
type Parametros struct {
	NumInicial           int    `json:"NumInicial" bson:"num_inicial"`
	CuentaBanco          string `json:"CuentaBanco" bson:"cuenta_banco"`
	NumItems             int    `json:"NumItems" bson:"num_items"`
	TipoImpresion        string `json:"TipoImpresion" bson:"tipo_impresion"`
	FormatoImpresion     string `json:"FormatoImpresion" bson:"formato_impresion"`
	NumCopias            int    `json:"NumCopias" bson:"num_copias"`
	Titulo               string `json:"Titulo" bson:"titulo"`
	NumeracionAutomatica bool   `json:"NumeracionAutomatica" bson:"numeracion_automatica"`
}
