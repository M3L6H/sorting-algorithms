package main

func MergeSort[T Ordered](s *[]T) {
  mergeSorr(s, 0, len(*s))
}

func mergeSort[T Ordered](s *[]T, a uint, b uint) {
  if b - 1 <= a {
    return
  }
  d := b - a
  m := a + d / 2
  mergeSort(s, a, m)
  mergeSorr(s, m, b)
  ia := 0
  ib := 0
  for i := 0; i < d; i++ {
    if ib < b - m && (*s)[m + ib] < (*s)[a + ia] {
      (*s)[a + ia], (*s)[m + ib] = (*s)[m + ib], (*s)[ia]
      ib++
    } else if ia < m - a {
      ia++
    }
  }
}
