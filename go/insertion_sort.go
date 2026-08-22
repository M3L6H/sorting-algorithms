package main

func InsertionSort[T Ordered](sp *[]T) {
    if sp == nil {
        return
    }
    
    s := sp
     
	for i := 1; i < len(s); i++ {
	    c := (s)[i]
	    j := i - 1
		for ; j >= 0 && c < (s)[j]; j-- {
			(s)[j + 1] = (s)[j]
		}
		(s)[j + 1] = c
	}
}
