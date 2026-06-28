package avatar

import (
	"fmt"
	"image"
	"image/png"
	"math/rand"
	"os"
	"time"

	"github.com/ethercod3/goAvatar/colors"
	"github.com/ethercod3/goAvatar/images"
	"github.com/ethercod3/goAvatar/patterns"
)

// Options controls avatar generation.
type Options struct {
	Dimensions int
	FileSizePx int
	Seed       int64
}

// Validate checks whether options can produce a usable avatar image.
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

// New validates options and returns a generated avatar image.
func New(options Options) (*image.RGBA, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}
	return Generate(options), nil
}

// Generate returns a generated avatar image.
//
// Use New when options come from user input and need validation.
func Generate(options Options) *image.RGBA {
	rng := newRand(options.Seed)
	return generateWithRand(options, rng)
}

func generateWithRand(options Options, rng *rand.Rand) *image.RGBA {
	scheme := colors.GenerateColorSchemeWithRand(rng)
	pattern := patterns.GeneratePatternWithRand(rng, options.Dimensions)
	return images.Draw(pattern, scheme, options.FileSizePx, options.Dimensions)
}

// GenerateWithSeed returns a reproducible avatar image for seed.
func GenerateWithSeed(options Options, seed int64) *image.RGBA {
	rng := rand.New(rand.NewSource(seed))
	return generateWithRand(options, rng)
}

// SavePNG writes img as a PNG file.
func SavePNG(path string, img image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

// Save validates options, generates an avatar, and writes it as PNG.
func Save(path string, options Options) error {
	img, err := New(options)
	if err != nil {
		return err
	}
	return SavePNG(path, img)
}

func newRand(seed int64) *rand.Rand {
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	return rand.New(rand.NewSource(seed))
}
