package nodoCuentaContableTestCases

import "testing"

func TestNodeAccountPipeline(t *testing.T) {
	// <setup code>
	t.Run("Registro del nodo fallido, success process test", TestNodoFail)
	// <tear-down code>
}
