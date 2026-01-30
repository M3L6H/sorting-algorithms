package main

import (
  "slices"
  "testing"
) 

func TestEmpty(t *testing.T) {
  i := []int{}
  d := []int{}
  e := []int{}
  BubbleSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("BubbleSort(%v) = %v; want %v", i, d, e);
  }
}

func TestOne(t *testing.T) {
  i := []int{1}
  d := []int{1}
  e := []int{1}
  BubbleSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("BubbleSort(%v) = %v; want %v", i, d, e);
  }
}

func TestSwap(t *testing.T) {
  i := []int{2, 1}
  d := []int{2, 1}
  e := []int{1, 2}
  BubbleSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("BubbleSort(%v) = %v; want %v", i, d, e);
  }
}