package main

import (
	"randomAvatarApi/colors"
	"randomAvatarApi/images"
	"randomAvatarApi/patterns"
)

func main() {
	dimensions := 5
	pattern := patterns.GeneratePattern(dimensions)
	images.Draw(pattern, colors.GenerateColorScheme(), 1000, dimensions)
}
