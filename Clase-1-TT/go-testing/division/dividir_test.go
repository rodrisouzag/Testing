package division

import "testing"

func TestDividir(t *testing.T) {

	num := 20
	den := 0
	resultadoEsperado := 4

	resultado, err := Dividir(num, den)
	if err != nil {
		if err != ErrDen0 {
			t.Error(err)
		}
		resultadoEsperado = 0
	}
	if resultado != resultadoEsperado {
		t.Errorf("Funcion dividir() arrojo el resultado = %v, pero el esperado es %v", resultado, resultadoEsperado)
	}

}
