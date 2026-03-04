package effingo

type Array[T any] struct {
	data []T
}

func From[T any](arr []T) Array[T] {
	return Array[T]{
		data: arr,
	}
}

func (arr Array[T]) Filter(f func(T) bool) Array[T] {
	return From(Filter(arr.data, f))
}

func ReduceTo[T any, Acc any](arr Array[T], acc Acc, f func(T, Acc) Acc) Acc {
	return Reduce(arr.data, acc, f)
}

func MapTo[T any, U any](arr Array[T], f func(element T) U) Array[U] {
	return From(Map(arr.data, f))
}

func (arr Array[T]) ForEach(f func(el T)) Array[T] {
	for _, v := range arr.data {
		f(v)
	}

	return arr
}

func (arr Array[T]) Find(f func(T) bool) (int, *T) {
	index, value := Find(arr.data, f)

	return index, value
}
