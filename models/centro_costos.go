package models

var CentroCostosCollection = "centro_gestor"

type CentroCostos struct {
	ID    string `json:"Id" bson:"_id,omitempty"`
	Label string `json:"Label" bson:"label"`
}
