package main

import (
  "slices"
  "testing"
) 

func TestMergeSortEmpty(t *testing.T) {
  i := []int{}
  d := []int{}
  e := []int{}
  MergeSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("MergeSort(%v) = %v; want %v", i, d, e);
  }
}

func TestMergeSortOne(t *testing.T) {
  i := []int{1}
  d := []int{1}
  e := []int{1}
  MergeSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("MergeSort(%v) = %v; want %v", i, d, e);
  }
}

func TestMergeSortSwap(t *testing.T) {
  i := []int{2, 1}
  d := []int{2, 1}
  e := []int{1, 2}
  MergeSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("MergeSort(%v) = %v; want %v", i, d, e);
  }
}

func TestMergeSortOdd(t *testing.T) {
  i := []int{5, 1, 4}
  d := []int{5, 1, 4}
  e := []int{1, 4, 5}
  MergeSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("MergeSort(%v) = %v; want %v", i, d, e);
  }
}

func TestMergeSortAllSame(t *testing.T) {
  i := []int{8, 8, 8, 8, 8}
  d := []int{8, 8, 8, 8, 8}
  e := []int{8, 8, 8, 8, 8}
  MergeSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("MergeSort(%v) = %v; want %v", i, d, e);
  }
}

func TestMergeSortMany(t *testing.T) {
  i := []int{7, 8, 6, 3, 9, 1, 6, 5, 4, 8, 3}
  d := []int{7, 8, 6, 3, 9, 1, 6, 5, 4, 8, 3}
  e := []int{1, 3, 3, 4, 5, 6, 6, 7, 8, 8, 9}
  MergeSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("MergeSort(%v) = %v; want %v", i, d, e);
  }
}
