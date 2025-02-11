package goAvatar

import (
	"github.com/ethercod3/goAvatar/colors"
	"github.com/ethercod3/goAvatar/images"
	"github.com/ethercod3/goAvatar/patterns"
)

type AvatarOptions struct {
	Filepath               string
	Dimensions, FileSizePx int
}

func GenerateAvatar(options AvatarOptions) error {
	scheme := colors.GenerateColorScheme()
	pattern := patterns.GeneratePattern(options.Dimensions)
	err := images.Draw(pattern, scheme, options.FileSizePx, options.Dimensions, options.Filepath)
	return err
}
