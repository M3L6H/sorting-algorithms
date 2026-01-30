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

func TestAllSame(t *testing.T) {
  i := []int{8, 8, 8, 8, 8}
  d := []int{8, 8, 8, 8, 8}
  e := []int{8, 8, 8, 8, 8}
  BubbleSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("BubbleSort(%v) = %v; want %v", i, d, e);
  }
}

func TestMany(t *testing.T) {
  i := []int{7, 8, 6, 3, 9, 1, 6, 5, 4, 8, 3}
  d := []int{7, 8, 6, 3, 9, 1, 6, 5, 4, 8, 3}
  e := []int{1, 3, 3, 4, 5, 6, 6, 7, 8, 8, 9}
  BubbleSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("BubbleSort(%v) = %v; want %v", i, d, e);
  }
}
