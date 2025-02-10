package patterns

import (
	"math/rand"
)

type Pattern struct {
	Left, Middle, Right []int
}

func generateLeftSideIndexes(dimensions int) []int {
	var indexes []int
	cropped := dimensions / 2
	for range dimensions {
		side := rand.Intn(cropped)
		indexes = append(indexes, side)
	}
	return indexes
}

func generateMiddleIndexes(dimensions int) []int {
	var indexes []int
	for row := range dimensions {
		if include := rand.Intn(2) == 1; include {
			indexes = append(indexes, row)
		}
	}
	return indexes
}

func generateRightSideIndexes(leftSide []int, dimensions int) []int {
	var indexes []int
	cropped := dimensions / 2
	for _, cellIndex := range leftSide {
		indexes = append(indexes, cropped-cellIndex-1)
	}
	return indexes
}

func GeneratePattern(dimensions int) Pattern {
	left := generateLeftSideIndexes(dimensions)
	middle := generateMiddleIndexes(dimensions)
	right := generateRightSideIndexes(left, dimensions)
	return Pattern{left, middle, right}
}
