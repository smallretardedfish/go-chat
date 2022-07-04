package slice

func Map[T, K any](s []T, f func(item T) K) []K {
	res := make([]K, 0)
	for _, item := range s {
		res = append(res, f(item))
	}
	return res
}
