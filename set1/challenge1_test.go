package main

import "testing"

func TestHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	test := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	result := hexToBase64(input)

	if result != test {
		t.Fatalf("Conversion from hex %s to base64 failed. expected=%s, got=%s", input, test, result)
	}
}
