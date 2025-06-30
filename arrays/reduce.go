package main

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
	result := initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}
