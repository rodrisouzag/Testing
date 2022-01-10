package ordenamiento

import "testing"

func TestOrdenar(t *testing.T) {
	unsorted := []int{-5, 23, 5, 19, 8, 1, 0, 567}
	resultadoEsperado := []int{-5, 0, 1, 5, 8, 19, 23, 567}

	resultado := Ordenar(unsorted)

	equals := true
	for i, n := range resultadoEsperado {
		if n != resultado[i] {
			equals = false
			break
		}
	}
	if !equals {
		t.Errorf("Funcion Orenar() arrojo el resultado = %v, pero el esperado es %v", resultado, resultadoEsperado)
	}
}
