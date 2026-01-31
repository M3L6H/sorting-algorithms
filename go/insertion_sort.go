package main

func InsertionSort[T Ordered](s *[]T) {
	for i, _ := range *s {
		for j := 0; j < i; j++ {
			if (*s)[j] > (*s)[i] {
				(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
			}
		}
	}
}
