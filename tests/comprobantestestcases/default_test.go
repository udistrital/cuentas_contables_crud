package comprobantestestcases

import "testing"

func TestPresupuestalAssignationPipeline(t *testing.T) {
	// <setup code>
	t.Run("Registro del comprobante exitoso, success process test", TestComprobanteSuccess)
	t.Run("Registro del comprobante fallido, fail process test", TestComprobanteFail)
	t.Run("Registro del TipoComprobante exitoso, success process test", TestTipoComprobanteSuccess)
	t.Run("Update de Parametros del Comprobante exitoso, success process test", TestComprobanteParamSuccess)
	t.Run("Update de Parametros del Comprobante fallido, fail process test", TestComprobanteParamFail)
	// <tear-down code>
}
