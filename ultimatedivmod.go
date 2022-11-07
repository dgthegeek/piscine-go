package piscine

func UltimateDivMod(a, b *int) {
	temporallystorea := *a
	*a = *a / *b
	*b = temporallystorea % *b
}
