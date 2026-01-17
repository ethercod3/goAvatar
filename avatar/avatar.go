package avatar

import (
	"image"

	"github.com/ethercod3/goAvatar/colors"
	"github.com/ethercod3/goAvatar/images"
	"github.com/ethercod3/goAvatar/patterns"
)

type Options struct {
	Dimensions int
	FileSizePx int
}

func Generate(options Options) *image.RGBA {
	scheme := colors.GenerateColorScheme()
	pattern := patterns.GeneratePattern(options.Dimensions)
	return images.Draw(pattern, scheme, options.FileSizePx, options.Dimensions)
}
