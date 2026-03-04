package effingo

func Map[T any, U any](arr []T, f func(arrElement T) U) []U {
	result := make([]U, len(arr))

	for i, v := range arr {
		composeResult := f(v)

		result[i] = composeResult
	}

	return result
}

func Filter[T any, U any](arr []T, f func(arrElement T) bool) []T {
	result := []T{}

	for _, v := range arr {
		composeResult := f(v)

		if composeResult {
			result = append(result, v)
		}
	}
	return result
}

func ForEach[T any](arr []T, f func(T)) {
	for _, v := range arr {
		f(v)
	}
}

func Reduce[T any, Acc any](arr []T, acc Acc, f func(T, Acc) Acc) Acc {
	accumulator := acc

	ForEach(arr, func(elem T) { accumulator = f(elem, accumulator) })

	return accumulator
}

func Find[T any](arr []T, f func(T) bool) (index int, value *T) {
	for i, v := range arr {
		if f(v) {
			return i, &v
		}
	}

	return -1, nil
}
