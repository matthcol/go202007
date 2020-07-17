package calcul

func fibo(rank uint) uint {
	var a uint = 0
	var b uint = 1
	switch rank {
	case 0:
		return a
	case 1:
		return b
	default:
		for i := uint(2); i < rank+1; i++ {
			c := a + b
			a, b = b, c
		}
		return b
	}
}
