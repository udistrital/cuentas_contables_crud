package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:CentroCostos"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:CentroCostos"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ComprobanteController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ComprobanteController"],
		beego.ControllerComments{
			Method:           "AddComprobante",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ComprobanteController"],
		beego.ControllerComments{
			Method:           "GetByUUID",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ComprobanteController"],
		beego.ControllerComments{
			Method:           "UpdateComprobante",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ComprobanteController"],
		beego.ControllerComments{
			Method:           "DeleteComprobante",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ConceptoController"],
		beego.ControllerComments{
			Method:           "Add",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ConceptoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ConceptoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:UUID`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ConceptoController"],
		beego.ControllerComments{
			Method:           "UpdateNode",
			Router:           `/:UUID`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:ConceptoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:UUID`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:DetalleCuentaContable"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:DetalleCuentaContable"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NaturalezaCuentaContable"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NaturalezaCuentaContable"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"],
		beego.ControllerComments{
			Method:           "AddNode",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"],
		beego.ControllerComments{
			Method:           "GetTree",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"],
		beego.ControllerComments{
			Method:           "GetByUUID",
			Router:           `/:UUID`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"],
		beego.ControllerComments{
			Method:           "UpdateNode",
			Router:           `/:UUID`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"],
		beego.ControllerComments{
			Method:           "ChangeNodeState",
			Router:           `/change_node_state/:UUID`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/cuentas`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"],
		beego.ControllerComments{
			Method:           "GetByNaturalezaCuentaContable",
			Router:           `/cuentas/:NaturalezaCuentaContable`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"],
		beego.ControllerComments{
			Method:           "GetCuentasUsablesByNaturaleza",
			Router:           `/getCuentas/:NaturalezaCuentaContable`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"],
		beego.ControllerComments{
			Method:           "GetByNaturalezaArka",
			Router:           `/getNodosCuentasArka/:NaturalezaCuentaContable`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:TipoComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:TipoComprobanteController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:TipoComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:TipoComprobanteController"],
		beego.ControllerComments{
			Method:           "AddTipoComprobante",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:TipoComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:TipoComprobanteController"],
		beego.ControllerComments{
			Method:           "GetByUUID",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:TipoComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:TipoComprobanteController"],
		beego.ControllerComments{
			Method:           "UpdateTipoComprobante",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:TipoComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:TipoComprobanteController"],
		beego.ControllerComments{
			Method:           "DeleteTipoComprobante",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:TipoMoneda"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:TipoMoneda"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:TipoCuentaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:TipoCuentaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
