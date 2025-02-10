package patterns

import (
	"math/rand"
)

func GenerateLeftSideIndexes(dimensions int) []int {
	var indexes []int
	cropped := dimensions / 2
	for range dimensions {
		side := rand.Intn(cropped)
		indexes = append(indexes, side)
	}
	return indexes
}

func GenerateMiddleIndexes(dimensions int) []int {
	var indexes []int
	for row := range dimensions {
		if include := rand.Intn(2) == 1; include {
			indexes = append(indexes, row)
		}
	}
	return indexes
}

func GenerateRightSideIndexes(leftSide []int, dimensions int) []int {
	var indexes []int
	cropped := dimensions / 2
	for _, cellIndex := range leftSide {
		indexes = append(indexes, cropped-cellIndex-1)
	}
	return indexes
}
