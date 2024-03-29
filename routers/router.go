// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/cuentas_contables_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/nodo_cuenta_contable",
			beego.NSInclude(
				&controllers.NodoCuentaContableController{},
			),
		),
		beego.NSNamespace("/comprobante",
			beego.NSInclude(
				&controllers.ComprobanteController{},
			),
		),
		beego.NSNamespace("/tipo_comprobante",
			beego.NSInclude(
				&controllers.TipoComprobanteController{},
			),
		),
		beego.NSNamespace("/naturaleza_cuenta_contable",
			beego.NSInclude(
				&controllers.NaturalezaCuentaContable{},
			),
		),
		beego.NSNamespace("/tipo_moneda",
			beego.NSInclude(
				&controllers.TipoMoneda{},
			),
		),
		beego.NSNamespace("/detalle_cuenta_contable",
			beego.NSInclude(
				&controllers.DetalleCuentaContable{},
			),
		),
		beego.NSNamespace("/centro_costos",
			beego.NSInclude(
				&controllers.CentroCostos{},
			),
		),
		beego.NSNamespace("/concepto",
			beego.NSInclude(
				&controllers.ConceptoController{},
			),
		),
		beego.NSNamespace("/tipo_cuenta",
			beego.NSInclude(
				&controllers.TipoCuentaController{},
			),
		),
		beego.NSNamespace("/conceptos",
			beego.NSInclude(
				&controllers.ConceptosController{},
			),
		),
		beego.NSNamespace("/tipo_retencion",
			beego.NSInclude(
				&controllers.TipoRetencionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
