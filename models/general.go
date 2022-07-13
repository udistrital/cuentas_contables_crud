package models

// General estructura general de una entidad en plan de cuentas
type General struct {
	FechaCreacion     string `json:"FechaCreacion" bson:"fecha_creacion"`
	FechaModificacion string `json:"FechaModificacion" bson:"fecha_modificacion"`
	Activo            bool   `json:"Activo" bson:"activo"`
}

type RespuestaApi struct {
	Data interface{}
}
