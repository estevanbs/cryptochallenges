package main

import (
	"encoding/hex"
	"math"
)

type ValueScore struct {
	value []byte
	score int
}

func byteScore(value byte) int {
	if value == ' ' {
		return 4
	}
	if 'a' <= value && value <= 'z' {
		return 2
	}
	if 'A' <= value && value <= 'Z' {
		return 1
	}
	return 0
}

func getScore(value []byte) int {
	bytesScore := 0

	for i := 0; i < len(value); i++ {
		bytesScore += byteScore(value[i])
	}

	return bytesScore
}

func fixStringXordedChar(value string) ValueScore {
	decoded, e := hex.DecodeString(value)
	if e != nil {
		panic(e)
	}

	maxScore := ValueScore{value: []byte{0}, score: 0}
	for i := byte(0); i < math.MaxUint8; i++ {
		tmp := make([]byte, len(decoded))

		for j := 0; j < len(decoded); j++ {
			tmp[j] = decoded[j] ^ i
		}

		score := getScore(tmp)
		if score > maxScore.score {
			maxScore = ValueScore{value: tmp, score: score}
		}
	}

	return maxScore
}
