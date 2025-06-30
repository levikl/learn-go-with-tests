package main

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
	result := initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}

func Find[A any](items []A, predicate func(A) bool) (item A, found bool) {
	for _, x := range items {
		if predicate(x) {
			return x, true
		}
	}
	return
}
