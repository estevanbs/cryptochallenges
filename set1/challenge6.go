package main

import (
	"fmt"
	"math"
)

func hammingDistance(bytes1, bytes2 []byte) int {
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

func estimateKeySize(input []byte) float32 {
	results := make([][]float32, 5)
	for i := 0; i < 5; i++ {
		results[i] = make([]float32, 30)
		for keySize := 2; keySize <= 30; keySize++ {
			firstBytes := input[i : keySize+i]
			secondBytes := input[keySize+i : (keySize+i)*2]

			hd := hammingDistance(firstBytes, secondBytes)
			nhd := float32(hd) / float32(keySize)

			results[i][keySize-2] = nhd
		}
	}
	averages := make([]float32, 30)
	for i := 0; i < 30; i++ {
		for j := 0; j < 5; j++ {
			averages[i] += results[j][i]
		}
		averages[i] = averages[i] / 5
	}
	fmt.Println(averages)
	return 0
}

func breakRepeatingKeyXor(input []byte) string {
	return ""
}
