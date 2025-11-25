package main

import (
	"math"
)

func hammingDistance(block1, block2 []byte) int {
	distance := 0
	for i := 0; i < len(block1); i++ {
		xor := block1[i] ^ block2[i]
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

// maxKeySize deve ter no máximo metade de len(input)
func estimateKeySize(input []byte, minKeySize, maxKeySize int) int {
	keySizeChecks := maxKeySize - minKeySize

	results := make([][]float64, keySizeChecks)

	for keySize := minKeySize; keySize < maxKeySize; keySize++ {
		currentKeySizeIndex := keySize - minKeySize

		// use this blocksToCheck to calculate hamming distance of all blocks
		blocksToCheck := (len(input) / keySize) - 1

		results[currentKeySizeIndex] = make([]float64, blocksToCheck)
		for i := 0; i < blocksToCheck; i++ {
			firstBlock := input[i : keySize+i]
			secondBlock := input[keySize+i : keySize*2+i]

			hd := hammingDistance(firstBlock, secondBlock)
			nhd := float64(hd) / float64(keySize)

			results[currentKeySizeIndex][i] = nhd
		}
	}
	// smallestAverage := math.MaxFloat64
	// probableKeySize := 0
	// for i := 0; i < keySizeChecks; i++ {
	// 	var resultsSum float64 = 0
	// 	for j := 0; j < blocksToCheck; j++ {
	// 		resultsSum += results[j][i]
	// 	}
	// 	average := resultsSum / float64(blocksToCheck)
	// 	if average < smallestAverage {
	// 		smallestAverage = average
	// 		probableKeySize = i + minKeySize
	// 	}
	// }
	return 0
}

func breakRepeatingKeyXor(input []byte) string {
	return ""
}
