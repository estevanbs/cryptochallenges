package main

import (
	"encoding/hex"
	"testing"
)

func TestEstimateHammingDistance(t *testing.T) {
	text, e := hex.DecodeString(`0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`)

	if e != nil {
		panic(e)
	}

	var test float32 = 10

	result := estimateKeySize(text)

	if result != test {
		t.Fatalf(
			"Estimating hamming distance for string \"%s\" failed. expected=%f, got=%f",
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
