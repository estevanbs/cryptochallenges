package main

import "testing"

func TestFixStringXordedChar(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	test := "Cooking MC's like a pound of bacon"

	result := fixStringXordedChar(input)

	if string(result.value) != test {
		t.Fatalf("Fixing from string %s failed. expected=%s, got=%s", input, test, result.value)
	}
}
