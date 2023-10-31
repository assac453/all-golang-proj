package abs

func Abs[V ~int | ~float64](x V) V {
	if x > 0 {
		return x
	} else {
		return -x
	}
}
