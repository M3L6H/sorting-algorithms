package main

import (
  "slices"
  "testing"
) 

func TestBubbleSortEmpty(t *testing.T) {
  i := []int{}
  d := []int{}
  e := []int{}
  BubbleSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("BubbleSort(%v) = %v; want %v", i, d, e);
  }
}

func TestBubbleSortOne(t *testing.T) {
  i := []int{1}
  d := []int{1}
  e := []int{1}
  BubbleSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("BubbleSort(%v) = %v; want %v", i, d, e);
  }
}

func TestBubbleSortSwap(t *testing.T) {
  i := []int{2, 1}
  d := []int{2, 1}
  e := []int{1, 2}
  BubbleSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("BubbleSort(%v) = %v; want %v", i, d, e);
  }
}

func TestBubbleSortAllSame(t *testing.T) {
  i := []int{8, 8, 8, 8, 8}
  d := []int{8, 8, 8, 8, 8}
  e := []int{8, 8, 8, 8, 8}
  BubbleSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("BubbleSort(%v) = %v; want %v", i, d, e);
  }
}

func TestBubbleSortMany(t *testing.T) {
  i := []int{7, 8, 6, 3, 9, 1, 6, 5, 4, 8, 3}
  d := []int{7, 8, 6, 3, 9, 1, 6, 5, 4, 8, 3}
  e := []int{1, 3, 3, 4, 5, 6, 6, 7, 8, 8, 9}
  BubbleSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("BubbleSort(%v) = %v; want %v", i, d, e);
  }
}

func TestBubbleSortNeg(t *testing.T) {
  i := []int{-5, -8, 3, 1, -5, 10, 7, 3, -5}
  d := []int{-5, -8, 3, 1, -5, 10, 7, 3, -5}
  e := []int{-8, -5, -5, -5, 1, 3, 3, 7, 10}
  BubbleSort(&d)
  if !slices.Equal(d, e) {
    t.Errorf("BubbleSort(%v) = %v; want %v", i, d, e);
  }
}
