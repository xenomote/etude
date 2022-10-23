package num

type NumWidthMin struct {
	width int
	min   int
	prec  int
}

func (n NumWidthMin) Min() int {
	return n.min
}

func (n NumWidthMin) Max() int {
	// TODO factor in precision
	
	return n.min + 2 ^ n.width*8
}

func (n NumWidthMin) Width() int {
	return n.width
}

func (n NumWidthMin) Precision() int {
	return n.prec
}
