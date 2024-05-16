package toplvl

func Power(x, n uint) uint {
	var r uint = 1
	for i := 0; i < int(n); i++ {
		r = r * x
	}
	return r
}
