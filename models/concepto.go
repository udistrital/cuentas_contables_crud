package models

var ConceptoCollection = "concepto"

// Concepto ...
type Concepto struct {
	*General         `bson:"inline"`
	ID               string   `json:"ID" bson:"_id,omitempty"`
	Nombre           string   `json:"Nombre" bson:"nombre"`
	CuentasContables []string `json:"CuentasContables" bson:"cuentas_contables,omitempty"`
	SistemaID        string   `json:"SistemaID" bson:"sistema_id"`
}
