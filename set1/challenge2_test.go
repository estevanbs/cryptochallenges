package main

import "testing"

func TestFixedXor(t *testing.T) {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"

	test := "746865206b696420646f6e277420706c6179"

	result := fixedXor(input1, input2)

	if result != test {
		t.Fatalf("Fixed XOR failed. expected=%s, got=%s", test, result)
	}
}
