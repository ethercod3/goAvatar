package argparser

import (
	"flag"
)

func Parse() (int, int, string) {
	var dimensions int
	var size int
	var output string

	flag.IntVar(&dimensions, "d", 0, "Dimensions of the image")
	flag.IntVar(&dimensions, "dimensions", 0, "Dimensions of the image")

	flag.IntVar(&size, "s", 0, "Size in pixels")
	flag.IntVar(&size, "size", 0, "Size in pixels")

	flag.StringVar(&output, "o", "avatar.png", "output file")
	flag.StringVar(&output, "output", "avatar.png", "output file")

	flag.Parse()

	return dimensions, size, output
}
