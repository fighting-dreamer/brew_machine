package main

import "testing"

func TestFail(t *testing.T) {
	got := -1
	expected := 1
	if got != expected {
		t.Error("expected not same as got")
	}
}
