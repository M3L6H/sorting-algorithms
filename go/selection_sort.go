package main

func SelectionSort[T Ordered](s *[]T) {
	for i, _ := range *s {
	    mi := i
		for j := i + 1; j < len(*s); j++ {
			if (*s)[j] < (*s)[mi] {
				mi = j
			}
		}
		(*s)[i], (*s)[mi] = (*s)[mi], (*s)[i]
	}
}
