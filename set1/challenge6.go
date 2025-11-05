package main

import (
	"math"
)

func hammingDistance(string1, string2 string) int {
	bytes1 := []byte(string1)
	bytes2 := []byte(string2)

	distance := 0
	for i := 0; i < len(bytes1); i++ {
		xor := bytes1[i] ^ bytes2[i]
		for j := float64(0); j < 8; j++ {
			currentByte := byte(math.Pow(2, j))
			isDifferent := (xor & currentByte) > 0
			if isDifferent {
				distance++
			}
		}
	}

	return distance
}

func breakRepeatingKeyXor() string {

	return ""
}
