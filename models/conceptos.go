package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ArbolPlanMaestroCuentasContCollection ...
var ArbolConceptosCollection = "conceptos"

// ArbolCuentasContParametersCollection ...
var ArbolConceptosParametersCollection = "parametros_conceptos"

type Conceptos struct {
	*General			`bson:"inline"`
	ID					primitive.ObjectID	 `bson:"_id" json:"id"`
	Padre				*string   `json:"Padre" bson:"concepto_padre_id"`
	Nombre				string   `json:"Nombre" bson:"nombre"`
	AreaFuncional		int      `json:"AreaFuncional" bson:"area_funcional"`
	EntidadPresupuestal	int      `json:"EntidadPresupuestal" bson:"entidad_presupuestal"`
	ClaseTransaccionId  int   `json:"ClaseTransaccionId" bson:"clase_transaccion_id"`
	TipoTransaccionId	int      `json:"TipoTransaccionId" bson:"tipo_transaccion_id"`
	Codigo				string   `json:"Codigo" bson:"codigo"`
	TipoComprobanteId   string   `json:"TipoComprobanteId" bson:"tipo_comprobante_id"`
	RubroPresupuestalId string   `json:"RubroPresupuestalId" bson:"rubro_presupuestal_id"`
	CodigoBogdata		string   `json:"CodigoBogdata" bson:"codigo_bogdata"`
	CuentasCredito		[]string `json:"CuentasCredito" bson:"cuentas_credito"`
	CuentasDebito		[]string `json:"CuentasDebito"  bson:"cuentas_debito"`
	Nivel				int      `json:"Nivel" bson:"nivel"`
	Hijos				[]string `json:"Hijos" bson:"hijos"`
	Aplicacion			string	 `json:"Aplicacion" bson:"aplicacion"`
	Metadatos			string   `json:"Metadatos" bson:"metadatos"`
}

// NodoArbolConceptos This struct is iseful for reduce band with usage in services that build a tree.
type NodoArbolConceptos struct {
	*General `bson:"inline"`
	ID		primitive.ObjectID	 		`bson:"_id" json:"id"`
	Hijos    []string                   `json:"Hijos" bson:"hijos,omitempty"`
	HijosRef []*NodoArbolConceptos 		`json:"children" bson:"-"`
	Padre    *string                    `json:"Padre" bson:"concepto_padre_id,omitempty"` // if the field is optional we put it as pointer.
	Nombre   string                     `json:"Nombre" bson:"nombre"`
	Nivel    int                        `json:"Nivel" bson:"nivel"`
	Codigo	 string   					`json:"Codigo" bson:"codigo"`
}

// ArbolConceptosParameters represents the paremeters for some Arbol Conceptos process.
type ArbolConceptosParameters struct {
	ID         string `json:"ID" bson:"_id,omitempty"`
	Nivel      *int   `json:"Nivel" bson:"nivel"`
	CodeLenght *int   `json:"CodeLenght" bson:"longitud_codigo"`
}

// ArbolConceptosFormatNode ...
type ArbolConceptosFormatNode struct {
	Data     *NodoArbolConceptos `json:"data" bson:"-"`
	Children []*ArbolConceptosFormatNode     `json:"children" bson:"-"`
}
