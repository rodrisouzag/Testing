package division

import "errors"

var ErrDen0 = errors.New("no se puede dividir entre 0")

func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, ErrDen0
	}
	return num / den, nil
}
