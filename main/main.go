package main

import (
	"fmt"
	"randomAvatarApi/patterns"
)

func main() {
	left := patterns.GenerateLeftSideIndexes()
	fmt.Println(left)
	fmt.Println(patterns.GenerateMiddleIndexes())
	fmt.Println(patterns.GenerateRightSideIndexes(left))
}
