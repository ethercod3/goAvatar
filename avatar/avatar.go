package avatar

import (
	"fmt"
	"image"

	"github.com/ethercod3/goAvatar/colors"
	"github.com/ethercod3/goAvatar/images"
	"github.com/ethercod3/goAvatar/patterns"
)

type Options struct {
	Dimensions int
	FileSizePx int
}

func (o Options) Validate() error {
	if o.Dimensions < 2 {
		return fmt.Errorf("dimensions must be at least 2")
	}
	if o.FileSizePx < 1 {
		return fmt.Errorf("file size must be positive")
	}
	if o.FileSizePx < o.Dimensions {
		return fmt.Errorf("file size must be greater than or equal to dimensions")
	}
	return nil
}

func Generate(options Options) *image.RGBA {
	scheme := colors.GenerateColorScheme()
	pattern := patterns.GeneratePattern(options.Dimensions)
	return images.Draw(pattern, scheme, options.FileSizePx, options.Dimensions)
}
