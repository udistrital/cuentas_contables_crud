package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"],
        beego.ControllerComments{
            Method: "AddNode",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"],
        beego.ControllerComments{
            Method: "GetTree",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cuentas_contables_crud/controllers:NodoCuentaContableController"],
        beego.ControllerComments{
            Method: "GetByUUID",
            Router: `/:UUID`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
