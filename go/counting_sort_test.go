package main

import (
  "slices"
  "testing"
) 

func TestCountingSortEmpty(t *testing.T) {
  i := []int{}
  d := []int{}
  e := []int{}
  CountingSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("CountingSort(%v) = %v; want %v", i, d, e);
  }
}

func TestCountingSortOne(t *testing.T) {
  i := []int{1}
  d := []int{1}
  e := []int{1}
  CountingSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("CountingSort(%v) = %v; want %v", i, d, e);
  }
}

func TestCountingSortSwap(t *testing.T) {
  i := []int{2, 1}
  d := []int{2, 1}
  e := []int{1, 2}
  CountingSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("CountingSort(%v) = %v; want %v", i, d, e);
  }
}

func TestCountingSortOdd(t *testing.T) {
  i := []int{5, 1, 4}
  d := []int{5, 1, 4}
  e := []int{1, 4, 5}
  CountingSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("CountingSort(%v) = %v; want %v", i, d, e);
  }
}

func TestCountingSortAllSame(t *testing.T) {
  i := []int{8, 8, 8, 8, 8}
  d := []int{8, 8, 8, 8, 8}
  e := []int{8, 8, 8, 8, 8}
  CountingSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("CountingSort(%v) = %v; want %v", i, d, e);
  }
}

func TestCountingSortMany(t *testing.T) {
  i := []int{7, 8, 6, 3, 9, 1, 6, 5, 4, 8, 3}
  d := []int{7, 8, 6, 3, 9, 1, 6, 5, 4, 8, 3}
  e := []int{1, 3, 3, 4, 5, 6, 6, 7, 8, 8, 9}
  CountingSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("CountingSort(%v) = %v; want %v", i, d, e);
  }
}

func TestCountingSortGaps(t *testing.T) {
  i := []int{10, 50, 50, 10, 50, 20, 10, 40}
  d := []int{10, 50, 50, 10, 50, 20, 10, 40}
  e := []int{10, 10, 10, 20, 40, 50, 50, 50}
  CountingSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("CountingSort(%v) = %v; want %v", i, d, e);
  }
}

func TestCountingSortNeg(t *testing.T) {
  i := []int{-5, -8, 3, 1, -5, 10, 7, 3, -5}
  d := []int{-5, -8, 3, 1, -5, 10, 7, 3, -5}
  e := []int{-8, -5, -5, -5, 1, 3, 3, 7, 10}
  CountingSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("CountingSort(%v) = %v; want %v", i, d, e);
  }
}
