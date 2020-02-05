package models

var ConceptoCollection = "concepto"

// Concepto ...
type Concepto struct {
	*General         `bson:"inline"`
	ID               string   `json:"ID" bson:"_id,omitempty"`
	Nombre           string   `json:"Nombre" bson:"nombre"`
	CuentaDebito     string   `json:"CuentaDebito" bson:"cuenta_debito,omitempty"`
	CuentaCredito    string   `json:"CuentaCredito" bson:"cuenta_credito,omitempty"`
	SistemaID        string   `json:"MovimientoID" bson:"movimiento_id"`
}
