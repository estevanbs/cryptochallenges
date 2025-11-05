package main

import "testing"

func TestHammingDistance(t *testing.T) {
	string1 := "this is a test"
	string2 := "wokka wokka!!!"

	result := hammingDistance(string1, string2)

	test := 37

	if result != test {
		t.Fatalf(
			"Calculating hamming distance from string \"%s\" and \"%s\" failed. expected=%d, got=%d",
			string1,
			string2,
			test,
			result)
	}
}
