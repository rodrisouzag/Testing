package calculadora

// Se importa el package testing
import "testing"

func TestRestar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 5
	num2 := 3
	resultadoEsperado := 0
	// Se ejecuta el test
	resultado := Restar(num1, num2)
	// Se validan los resultados
	if resultado != resultadoEsperado {
		t.Errorf("Funcion resta() arrojo el resultado = %v, pero el esperado es %v", resultado, resultadoEsperado)
	}
}
