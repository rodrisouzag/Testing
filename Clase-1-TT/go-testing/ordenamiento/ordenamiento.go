package ordenamiento

import "sort"

func Ordenar(numeros []int) []int {
	s := numeros
	sort.Ints(sort.IntSlice(s))
	return s
}
