package main

func BubbleSort[T Ordered](s *[]T) {
  sorted := false
  for !sorted {
    sorted = true
    for i, v := range *s {
      if i == len(*s) - 1 {
        break
      }

      if v > (*s)[i + 1] {
        (*s)[i], (*s)[i + 1] = (*s)[i + 1], (*s)[i]
        sorted = false
      }
    }
  }
}
