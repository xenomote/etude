package num

import "math"

type NumMinMax struct {
	min  int
	max  int
	prec int
}

func (n NumMinMax) Min() int {
	return n.min
}

func (n NumMinMax) Max() int {
	return n.max
}

func (n NumMinMax) Width() int {
	diff := float64(n.max - n.min)
	bits := math.Ceil(math.Log2(diff))
	bytes := int(math.Ceil(bits / 8))

	// TODO factor in precision

	return bytes
}

func (n NumMinMax) Precision() int {
	return n.prec
}
