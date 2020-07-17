package calcul

// Pgcd compute greater common diviser of 2 integers
func Pgcd(a, b uint) uint {
	for b != 0 {
		r := a % b
		a, b = b, r
	}
	return a
}
