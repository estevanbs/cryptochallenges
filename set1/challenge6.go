package main

import (
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

func estimateKeySize(input []byte) int {
	const MIN_KEY_SIZE = 2
	const BLOCKS_TO_CHECK = 4
	// use this MAX_KEY_SIZE to calculate hamming distance of all blocks
	MAX_KEY_SIZE := len(input)/2 - BLOCKS_TO_CHECK - 1

	keySizeChecks := MAX_KEY_SIZE - MIN_KEY_SIZE

	results := make([][]float64, BLOCKS_TO_CHECK)
	for i := 0; i < BLOCKS_TO_CHECK; i++ {
		results[i] = make([]float64, keySizeChecks)
		for keySize := MIN_KEY_SIZE; keySize < MAX_KEY_SIZE; keySize++ {
			firstBytes := input[i : keySize+i]
			secondBytes := input[keySize+i : (keySize+i)*2]

			hd := hammingDistance(firstBytes, secondBytes)
			nhd := float64(hd) / float64(keySize)

			results[i][keySize-MIN_KEY_SIZE] = nhd
		}
	}
	smallestAverage := math.MaxFloat64
	probableKeySize := 0
	for i := 0; i < keySizeChecks; i++ {
		var resultsSum float64 = 0
		for j := 0; j < BLOCKS_TO_CHECK; j++ {
			resultsSum += results[j][i]
		}
		average := resultsSum / BLOCKS_TO_CHECK
		if average < smallestAverage {
			smallestAverage = average
			probableKeySize = i + MIN_KEY_SIZE
		}
	}
	return probableKeySize
}

func breakRepeatingKeyXor(input []byte) string {
	return ""
}
