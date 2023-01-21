package main

import (
	"testing"
)

func TestbinarySearsh(t *testing.T) {
	a := []int{1, 3, 5, 7, 9}
	got := binarySearch(a, 5)
	want := 2
	if got != want {
		t.Errorf("binarySearch == %d, want %d", got, want)
	}
}
