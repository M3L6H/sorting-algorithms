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
  c := make([]int, mx + 1)
  for _, elt := range *s {
    if elt >= 0 {
      c[elt * 2]++
    } else {
      c[(elt * -2) - 1]++
    }
  }
  j := 0
  for i := len(c) - (len(c) % 2); i > 0; i -= 2 {
    n := (i) / -2
    for k := 0; k < c[i - 1]; k++ {
      (*s)[j] = n
      j++
    }
  }
  for i := 0; i < len(c); i += 2 {
    n := i / 2
    for k := 0; k < c[i]; k++ {
      (*s)[j] = n
      j++
    }
  }
}