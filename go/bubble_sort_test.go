package main

import "testing"

func TestEmpty(t *testing.T) {
  i := []int{}
  d := make([]int, len(i))
  copy(d, i)
  e := []int{}
  BubbleSort(&d)
  if d != []int{} {
    t.Errorf("BubbleSort(%v) = %v; want %v", i, d, e);
  }
}

func TestOne(t *testing.T) {
  i := []int{1}
  d := make([]int, len (i))
  copy(d, i)
  e := []int{1}
  BubbleSort(&d)
  if d != []int{} {
    t.Errorf("BubbleSort(%v) = %v; want %v", i, d, e);
  }
}

func TestSwap(t *testing.T) {
  i := []int{2, 1}
  d := make([]int, len (i))
  copy(d, i)
  e := []int{1, 2}
  BubbleSort(&d)
  if d != []int{} {
    t.Errorf("BubbleSort(%v) = %v; want %v", i, d, e);
  }
}