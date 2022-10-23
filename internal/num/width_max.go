package num

type NumWidthMax struct {
	width int
	max   int
	prec  int
}

func (n NumWidthMax) Min() int {
	// TODO factor in precision

	return n.max - 2 ^ n.width*8
}

func (n NumWidthMax) Max() int {
	return n.max
}

func (n NumWidthMax) Width() int {
	return n.width
}

func (n NumWidthMax) Precision() int {
	return n.prec
}
