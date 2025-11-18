package main

import (
	"encoding/base64"
	"encoding/hex"
	"os"
	"testing"
)

func TestEstimateKeySize(t *testing.T) {
	text := "hey, this is a test. It should pass. My name isn't Jorginho!!"
	key := "secrete"

	encryptedHex := encryptRepeatingKeyXor(text, key)

	encrypted, e := hex.DecodeString(encryptedHex)

	if e != nil {
		panic(e)
	}

	test := len(key)

	result := estimateKeySize(encrypted, 2, 10)

	if result != test {
		t.Fatalf(
			"Estimating key size for string \"%s\" failed. expected=%d, got=%d",
			text,
			test,
			result)
	}
}


func TestEstimateKeySizeFile(t *testing.T) {
	content, e := os.ReadFile("challenge6_data.txt")
	if e != nil {
		panic(e)
	}

	encrypted, e := base64.StdEncoding.DecodeString(string(content))
	if e != nil {
		panic(e)
	}

	test := 29 // not sure

	result := estimateKeySize(encrypted, 2, 40)

	if result != test {
		t.Fatalf(
			"Estimating key size failed. expected=%d, got=%d",
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
