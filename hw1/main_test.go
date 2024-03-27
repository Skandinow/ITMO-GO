package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	a, b := 2, 2
	want := 4

	if want != sum(a, b) {
		t.Errorf("want %d + %d = %d, but got %d", a, b, want, sum(a, b))
	}

}
