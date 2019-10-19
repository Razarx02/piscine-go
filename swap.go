package student

func Swap(a *int, b *int) {

	x := *b

	*a = *b

	*b = x

}
