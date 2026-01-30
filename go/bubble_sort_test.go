package main

import "testing"

func TestEmpty(t *testing.T) {
  i := [0]int{}
  d := [0]int{}
  e := [0]int{}
  BubbleSort(&d)
  if d != e {
    t.Errorf("BubbleSort(%v) = %v; want %v", i, d, e);
  }
}

func TestOne(t *testing.T) {
  i := [1]int{1}
  d := [1]int{1}
  e := [1]int{1}
  BubbleSort(&d)
  if d != e {
    t.Errorf("BubbleSort(%v) = %v; want %v", i, d, e);
  }
}

func TestSwap(t *testing.T) {
  i := [2]int{2, 1}
  d := [2]int{2, 1}
  e := [2]int{1, 2}
  BubbleSort(&d)
  if d != e {
    t.Errorf("BubbleSort(%v) = %v; want %v", i, d, e);
  }
}