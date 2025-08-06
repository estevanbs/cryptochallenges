package main

import "testing"

func TestFixFileMultipleStringXordedChar(t *testing.T) {
	file := "challenge4_data.txt"

	test := "Now that the party is jumping"

	result := fixFileMultipleStringXordedChar(file)

	if result != test {
		t.Fatalf("Fixing from file %s failed. expected=%s, got=%s", file, test, result)
	}
}
