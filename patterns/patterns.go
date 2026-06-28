package patterns

import (
	"math/rand"
	"time"
)

// Pattern stores the generated cell indexes for each avatar side.
type Pattern struct {
	Left, Middle, Right []int
}

func generateLeftSideIndexes(rng *rand.Rand, dimensions int) []int {
	var indexes []int
	cropped := dimensions / 2
	for range dimensions {
		side := rng.Intn(cropped)
		indexes = append(indexes, side)
	}
	return indexes
}

func generateMiddleIndexes(rng *rand.Rand, dimensions int) []int {
	var indexes []int
	for range dimensions {
		if include := rng.Intn(2) == 1; include {
			indexes = append(indexes, 1)
		} else {
			indexes = append(indexes, -1)
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

// GeneratePattern returns a symmetric avatar pattern.
func GeneratePattern(dimensions int) Pattern {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return GeneratePatternWithRand(rng, dimensions)
}

// GeneratePatternWithRand returns a symmetric avatar pattern using rng.
func GeneratePatternWithRand(rng *rand.Rand, dimensions int) Pattern {
	left := generateLeftSideIndexes(rng, dimensions)
	middle := generateMiddleIndexes(rng, dimensions)
	right := generateRightSideIndexes(left, dimensions)
	return Pattern{left, middle, right}
}
