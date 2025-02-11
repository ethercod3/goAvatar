package randomavatarapi

import (
	"randomAvatarApi/colors"
	"randomAvatarApi/images"
	"randomAvatarApi/patterns"
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
