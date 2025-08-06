package main

import (
	"os"
	"strings"
)

func fixFileMultipleStringXordedChar(filePath string) string {
	content, e := os.ReadFile(filePath)
	if e != nil {
		panic(e)
	}

	lines := strings.Split(strings.TrimSuffix(string(content), "\n"), "\n")
	maxScore := ValueScore{value: []byte{0}, score: 0}

	for _, line := range lines {
		valueScore := fixStringXordedChar(line)
		if valueScore.score > maxScore.score {
			maxScore = valueScore
		}
	}

	return strings.TrimSuffix(string(maxScore.value), "\n")
}
