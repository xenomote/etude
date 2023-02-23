package num

type NumWidth struct {
	width int
	prec  int
}

func (n NumWidth) Min() int {
	// TODO factor in precision

	return -n.Max() + 1
}

func (n NumWidth) Max() int {
	// TODO factor in precision

	return 2 ^ (n.width*8 - 1)
}

func (n NumWidth) Width() int {
	return n.width
}

func (n NumWidth) Precision() int {
	return n.prec
}
