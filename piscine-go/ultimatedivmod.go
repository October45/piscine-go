package piscine

func UltiimateDivMod(a *int, b *int) {
	div := *a / *b
	mod := *a % *b
	*a = div
	*b = mod
}
