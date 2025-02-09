package patterns

import (
	"math/rand"
)

func GenerateLeftSideIndexes() []int {

	var indexes []int

	for row := range 5 {
		row += 1
		side := rand.Intn(2)
		index := (row * 2) - (1 + side)
		if index < 0 {
			index = 0
		}
		indexes = append(indexes, index)
	}

	return indexes

}

func GenerateMiddleIndexes() []int {

	var indexes []int

	for row := range 5 {
		row += 1
		side := rand.Intn(2)
		index := row - (1 + side)
		if index < 0 {
			index = 0
		}
		if side > 0 {
			indexes = append(indexes, index)
		}
	}

	return indexes

}

func GenerateRightSideIndexes(leftSide []int) []int {
	var indexes []int
	for _, cellIndex := range leftSide {
		if cellIndex%2 == 0 {
			indexes = append(indexes, cellIndex+1)
		} else {
			indexes = append(indexes, cellIndex-1)
		}
	}
	return indexes
}
