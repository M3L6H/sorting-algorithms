package main

func MergeSort[T Ordered](s *[]T) {
  mergeSort(s, 0, len(*s))
}

func mergeSort[T Ordered](s *[]T, a int, b int) {
  if b - 1 <= a {
    return
  }
  d := b - a
  m := a + d / 2
  mergeSort(s, a, m)
  mergeSort(s, m, b)
  ia := a
  ib := m
  sorted := make([]T, 0, b - a)
  
  for i := 0; i < d; i++ {
    if ia >= m {
      sorted = append(sorted, (*s)[ib])
      ib++
    } else if ib >= b {
      sorted = append(sorted, (*s)[ia])
      ia++
    } else if (*s)[ia] > (*s)[ib] {
      sorted = append(sorted, (*s)[ib])
      ib++
    } else {
      sorted = append(sorted, (*s)[ia])
      ia++
    }
  }

  for i := a; i < b; i++ {
    (*s)[i] = sorted[i - a]
  }
}
