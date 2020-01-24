package models

// CentroCostosCollection ...
var CentroCostosCollection = "centro_gestor"

// CentroCostos ...
type CentroCostos struct {
	ID    string `json:"Id" bson:"_id,omitempty"`
	Label string `json:"Label" bson:"label"`
}
