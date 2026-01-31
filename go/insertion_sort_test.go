package main

import (
  "slices"
  "testing"
) 

func TestInsertionSortEmpty(t *testing.T) {
  i := []int{}
  d := []int{}
  e := []int{}
  InsertionSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("InsertionSort(%v) = %v; want %v", i, d, e);
  }
}

func TestInsertionSortOne(t *testing.T) {
  i := []int{1}
  d := []int{1}
  e := []int{1}
  InsertionSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("InsertionSort(%v) = %v; want %v", i, d, e);
  }
}

func TestInsertionSortSwap(t *testing.T) {
  i := []int{2, 1}
  d := []int{2, 1}
  e := []int{1, 2}
  InsertionSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("InsertionSort(%v) = %v; want %v", i, d, e);
  }
}

func TestInsertionSortAllSame(t *testing.T) {
  i := []int{8, 8, 8, 8, 8}
  d := []int{8, 8, 8, 8, 8}
  e := []int{8, 8, 8, 8, 8}
  InsertionSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("InsertionSort(%v) = %v; want %v", i, d, e);
  }
}

func TestInsertionSortMany(t *testing.T) {
  i := []int{7, 8, 6, 3, 9, 1, 6, 5, 4, 8, 3}
  d := []int{7, 8, 6, 3, 9, 1, 6, 5, 4, 8, 3}
  e := []int{1, 3, 3, 4, 5, 6, 6, 7, 8, 8, 9}
  InsertionSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("InsertionSort(%v) = %v; want %v", i, d, e);
  }
}
