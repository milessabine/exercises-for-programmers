package main

import "testing"

func TestCheckPassword(t *testing.T) {
	if checkPassword("12345") {
		t.Error("12345 is not a valid password.")
	}
	if !checkPassword("abc$123") {
		t.Error("abc$123 should be a valid password.")
	}
}
