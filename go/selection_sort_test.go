package main

import (
  "slices"
  "testing"
) 

func TestSelectionSortEmpty(t *testing.T) {
  i := []int{}
  d := []int{}
  e := []int{}
  SelectionSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("SelectionSort(%v) = %v; want %v", i, d, e);
  }
}

func TestSelectionSortOne(t *testing.T) {
  i := []int{1}
  d := []int{1}
  e := []int{1}
  SelectionSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("SelectionSort(%v) = %v; want %v", i, d, e);
  }
}

func TestSelectionSortSwap(t *testing.T) {
  i := []int{2, 1}
  d := []int{2, 1}
  e := []int{1, 2}
  SelectionSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("SelectionSort(%v) = %v; want %v", i, d, e);
  }
}

func TestSelectionSortAllSame(t *testing.T) {
  i := []int{8, 8, 8, 8, 8}
  d := []int{8, 8, 8, 8, 8}
  e := []int{8, 8, 8, 8, 8}
  SelectionSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("SelectionSort(%v) = %v; want %v", i, d, e);
  }
}

func TestSelectionSortMany(t *testing.T) {
  i := []int{7, 8, 6, 3, 9, 1, 6, 5, 4, 8, 3}
  d := []int{7, 8, 6, 3, 9, 1, 6, 5, 4, 8, 3}
  e := []int{1, 3, 3, 4, 5, 6, 6, 7, 8, 8, 9}
  SelectionSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("SelectionSort(%v) = %v; want %v", i, d, e);
  }
}

func TestSelectionSortNeg(t *testing.T) {
  i := []int{-5, -8, 3, 1, -5, 10, 7, 3, -5}
  d := []int{-5, -8, 3, 1, -5, 10, 7, 3, -5}
  e := []int{-8, -5, -5, -5, 1, 3, 3, 7, 10}
  SelectionSorr(&d)
  if !slices.Equal(d, e) {
    t.Errorf("SelectionSort(%v) = %v; want %v", i, d, e);
  }
}
