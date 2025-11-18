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

func estimateKeySize(input []byte, minKeySize, maxKeySize int) int {
	// use this blocksToCheck to calculate hamming distance of all blocks
	blocksToCheck := len(input)/2 - maxKeySize - 1

	keySizeChecks := maxKeySize - minKeySize

	results := make([][]float64, blocksToCheck)
	for i := 0; i < blocksToCheck; i++ {
		results[i] = make([]float64, keySizeChecks)
		for keySize := minKeySize; keySize < maxKeySize; keySize++ {
			firstBytes := input[i : keySize+i]
			secondBytes := input[keySize+i : (keySize+i)*2]

			hd := hammingDistance(firstBytes, secondBytes)
			nhd := float64(hd) / float64(keySize)

			results[i][keySize-minKeySize] = nhd
		}
	}
	smallestAverage := math.MaxFloat64
	probableKeySize := 0
	for i := 0; i < keySizeChecks; i++ {
		var resultsSum float64 = 0
		for j := 0; j < blocksToCheck; j++ {
			resultsSum += results[j][i]
		}
		average := resultsSum / float64(blocksToCheck)
		if average < smallestAverage {
			smallestAverage = average
			probableKeySize = i + minKeySize
		}
	}
	return probableKeySize
}

func breakRepeatingKeyXor(input []byte) string {
	return ""
}
