package main

import "testing"

func TestIsAnagram(t *testing.T) {
	a := isAnagram("note", "tone")
	if a != true {
		t.Error("Expected note and tone to be anagrams.")
	}
	a = isAnagram("noted", "tone")
	if a != false {
		t.Error("Expected noted and tone not to be anagrams.")
	}
	/*	a = isAnagram("Arrigo Boito", "Tobia Gorrio")
		if a != true {
			t.Error("Expected Arrigo Boito and Tobia Gorrio to be anagrams.")
		}*/
}
