package goAvatar

import (
	"image"

	"github.com/ethercod3/goAvatar/colors"
	"github.com/ethercod3/goAvatar/images"
	"github.com/ethercod3/goAvatar/patterns"
)

type AvatarOptions struct {
	Dimensions, FileSizePx int
}

func GenerateAvatar(options AvatarOptions) *image.RGBA {
	scheme := colors.GenerateColorScheme()
	pattern := patterns.GeneratePattern(options.Dimensions)
	img := images.Draw(pattern, scheme, options.FileSizePx, options.Dimensions)
	return img
}
