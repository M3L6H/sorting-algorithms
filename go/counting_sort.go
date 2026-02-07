package main

func CountingSort(s *[]int) {
  if (len(*s) == 0) {
    return
  }
  mx := max((*s)[0], -(*s)[0])
  for _, elt := range *s {
    elt_abs := max(elt, -elt)
    if elt_abs > mx {
      mx = elt_abs
    }
  }
  mx *= 2
  c := make([]int, 0, mx + 1)
  for _, elt := range *s {
    if elt >= 0 {
      c[elt * 2]++
    } else {
      c[(elt * -2) - 1]++
    }
  }
  j := 0
  for i := 0; i < len(c); i++ {
    for k := 0; k < c[i]; k++ {
      (*s)[j] = i
      j++
    }
  }
}