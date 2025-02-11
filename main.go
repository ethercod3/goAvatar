package goAvatar

import (
	"github.com/ethercod3/goAvatar/colors"
	"github.com/ethercod3/goAvatar/images"
	"github.com/ethercod3/goAvatar/patterns"
)

type AvatarOptions struct {
	filepath               string
	dimensions, fileSizePx int
}

func GenerateAvatar(options AvatarOptions) error {
	scheme := colors.GenerateColorScheme()
	pattern := patterns.GeneratePattern(options.dimensions)
	err := images.Draw(pattern, scheme, options.fileSizePx, options.dimensions, options.filepath)
	return err
}
