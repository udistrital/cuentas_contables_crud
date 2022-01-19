package models

// ArbolPlanMaestroCuentasContCollection ...
var ArbolPlanMaestroCuentasContCollection = "plan_mestro_cuentas_contables"

// ArbolCuentasContParametersCollection ...
var ArbolCuentasContParametersCollection = "parametros_plan_cuentas_contables"

// NodoCuentaContable This struct represents a tree's node of "plan cuentas contable" bussines model.
type NodoCuentaContable struct {
	*General            `bson:"inline"`
	ID                  string   `json:"Id" bson:"_id,omitempty"`
	Codigo              string   `json:"Codigo" bson:"codigo,omitempty"`
	Hijos               []string `json:"Hijos" bson:"hijos,omitempty"`
	Padre               *string  `json:"Padre" bson:"padre,omitempty"` // if the field is optional we put it as pointer.
	Nombre              string   `json:"Nombre" bson:"nombre"`
	Nivel               int      `json:"Nivel" bson:"nivel"`
	DetalleCuentaID     string   `json:"DetalleCuentaID" bson:"detalle_cuenta_id"`
	NaturalezaCuentaID  string   `json:"NaturalezaCuentaID" bson:"naturaleza_id"`
	CodigoCuentaAlterna string   `json:"CodigoCuentaAlterna" bson:"codigo_cuenta_alterna"`
	Ajustable           bool     `json:"Ajustable" bson:"ajustable"`
	MonedaID            string   `json:"MonedaID" bson:"moneda_id"`
	RequiereTercero     bool     `json:"RequiereTercero" bson:"requiere_tercero"`
	CentroDecostosID    string   `json:"CentroDecostosID" bson:"centro_costos_id"`
	Nmnc                bool     `json:"Nmnc" bson:"nmnc"`
}

// NodoArbolCuentaContable This struct is iseful for reduce band with usage in services that build a tree.
type NodoArbolCuentaContable struct {
	*General `bson:"inline"`
	ID       string                     `json:"Id" bson:"_id,omitempty"`
	Codigo   string                     `json:"Codigo" bson:"codigo,omitempty"`
	Hijos    []string                   `json:"Hijos" bson:"hijos,omitempty"`
	HijosRef []*NodoArbolCuentaContable `json:"children" bson:"-"`
	Padre    *string                    `json:"Padre" bson:"padre,omitempty"` // if the field is optional we put it as pointer.
	Nombre   string                     `json:"Nombre" bson:"nombre"`
	Nivel    int                        `json:"Nivel" bson:"nivel"`
}

// ArbolCuentaContableParameters represents the paremeters for some Arbol Cuentas contables process.
type ArbolCuentaContableParameters struct {
	ID         string `json:"ID" bson:"_id,omitempty"`
	Nivel      *int   `json:"Nivel" bson:"nivel"`
	CodeLenght *int   `json:"CodeLenght" bson:"longitud_codigo"`
}

// ArbolNbFormatNode ...
type ArbolNbFormatNode struct {
	Data     *NodoArbolCuentaContable `json:"data" bson:"-"`
	Children []*ArbolNbFormatNode     `json:"children" bson:"-"`
}

// ArkaCuentasContables...
type ArkaCuentasContables struct {
	ID              string `json:"Id" bson:"_id,omitempty"`
	Codigo          string `json:"Codigo" bson:"codigo,omitempty"`
	Descripcion     string `json:"DetalleCuentaID" bson:"detalle_cuenta_id"`
	Naturaleza      string `json:"Naturaleza" bson:"naturaleza_id"`
	Nombre          string `json:"Nombre" bson:"nombre"`
	RequiereTercero bool   `json:"RequiereTercero" bson:"requiere_tercero"`
}
