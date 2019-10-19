package student

func DivMod(a *int, b *int) {

	x := *a

	*a = *a / *b
	*b = x % *b

}
