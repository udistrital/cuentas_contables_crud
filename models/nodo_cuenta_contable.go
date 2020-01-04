package models

var ArbolPlanMaestroCuentasContCollection = "plan_mestro_cuentas_contables"
var ArbolCuentasContCollection = "plan_cuentas_contables"

// NodoCuentaContable This struct represents a tree's node of "plan cuentas contable" bussines model.
type NodoCuentaContable struct {
	*General
	ID                  string   `json:"Codigo" bson:"_id,omitempty"`
	Hijos               []string `json:"Hijos" bson:"hijos,omitempty"`
	Padre               *string  `json:"Padre" bson:"padre,omitempty"` // if the field is optional we put it as pointer.
	Nivel               int      `json:"Nivel" bson:"nivel"`
	DetalleCuenta       string   `json:"DetalleCuenta" bson:"detalle"`
	NaturalezaCuentaID  string   `json:"NaturalezaCuenta" bson:"naturaleza_id"`
	CodigoCuentaAlterna string   `json:"CodigoCuentaAlterna" bson:"codigo_cuenta_alterna"`
	Ajustable           bool     `json:"Ajustable" bson:"ajustable"`
	MonedaID            string   `json:"Moneda" bson:"moneda_id"`
	RequiereTercero     bool     `json:"RequiereTercero" bson:"requiere_tercero"`
	CentroDecostosID    string   `json:"CentroDecostosID" bson:"centro_costos_id"`
	Nmnc                bool     `json:"Nmnc" bson:"nmnc"`
}

// NodoArbolCuentaContable This struct is iseful for reduce band with usage in services that build a tree.
type NodoArbolCuentaContable struct {
	*General
	ID       string                     `json:"Codigo" bson:"_id,omitempty"`
	Hijos    []string                   `json:"Hijos" bson:"hijos,omitempty"`
	HijosRef []*NodoArbolCuentaContable `json:"children" bson:"-"`
	Padre    *string                    `json:"Padre" bson:"padre,omitempty"` // if the field is optional we put it as pointer.
}
