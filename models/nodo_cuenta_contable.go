package models

var ArbolPlanMaestroCuentasContCollection = "plan_mestro_cuentas_contables"
var ArbolCuentasContCollection = "plan_cuentas_contables"

// NodoCuentaContable This struct represents a tree's node of "plan cuentas contable" bussines model.
type NodoCuentaContable struct {
	ID            string   `json:"Codigo" bson:"_id,omitempty"`
	Hijos         []string `json:"Hijos" bson:"hijos,omitempty"`
	Padre         *string  `json:"Padre" bson:"padre,omitempty"` // if the field is optional we put it as pointer.
	FechaRegistro string   `json:"FechaRegistro" bson:"fecha_registro"`
}
