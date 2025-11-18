package main

import (
	"encoding/hex"
	"testing"
)

func TestEstimateKeySize(t *testing.T) {
	text := "hey, this is a test. It should pass. My name isn't Jorginho!!"
	key := "secret"

	encryptedHex := encryptRepeatingKeyXor(text, key)

	encrypted, e := hex.DecodeString(encryptedHex)

	if e != nil {
		panic(e)
	}

	test := len(key)

	result := estimateKeySize(encrypted)

	if result != test {
		t.Fatalf(
			"Estimating hamming distance for string \"%s\" failed. expected=%d, got=%d",
			text,
			test,
			result)
	}
}

func TestHammingDistance(t *testing.T) {
	string1 := []byte("this is a test")
	string2 := []byte("wokka wokka!!!")

	result := hammingDistance(string1, string2)

	test := 37

	if result != test {
		t.Fatalf(
			"Calculating hamming distance of strings \"%s\" and \"%s\" failed. expected=%d, got=%d",
			string1,
			string2,
			test,
			result)
	}
}
