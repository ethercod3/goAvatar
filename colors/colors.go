package colors

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"
)

const (
	DARKEN = iota
	LIGHTEN
)

// ColorChannel is one RGB channel.
type ColorChannel uint8

// Color is an RGB color.
type Color struct {
	R, G, B ColorChannel
}

// ColorScheme contains the background and foreground colors.
type ColorScheme struct {
	First, Second Color
}

func generateRandomColorChannel(rng *rand.Rand) ColorChannel {
	return ColorChannel(rng.Intn(256))
}

func generateRandomColor(rng *rand.Rand) Color {
	r := generateRandomColorChannel(rng)
	g := generateRandomColorChannel(rng)
	b := generateRandomColorChannel(rng)
	return Color{r, g, b}
}

func getLuminance(c Color) float32 {
	return (0.2126)*float32(c.R) + (0.7152)*float32(c.G) + (0.0722)*float32(c.B)
}

func darken(channel ColorChannel) ColorChannel {
	return ColorChannel(float32(channel) * 0.5)
}

func lighten(channel ColorChannel) ColorChannel {
	c := ColorChannel(float32(channel) * 1.5)
	if c < channel {
		return ColorChannel(255)
	}
	return c
}

func filter(c Color, filter int) Color {
	switch filter {
	case DARKEN:
		return Color{R: darken(c.R), G: darken(c.G), B: darken(c.B)}
	case LIGHTEN:
		return Color{R: lighten(c.R), G: lighten(c.G), B: lighten(c.B)}
	default:
		panic(fmt.Sprintf("Unknown filter: %v", filter))
	}
}

func generateSecondColor(c Color) Color {
	luminance := getLuminance(c)
	if luminance <= 128 {
		return filter(c, LIGHTEN)
	} else {
		return filter(c, DARKEN)
	}
}

// GenerateColorScheme returns a random two-color scheme.
func GenerateColorScheme() ColorScheme {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return GenerateColorSchemeWithRand(rng)
}

// GenerateColorSchemeWithRand returns a two-color scheme using rng.
func GenerateColorSchemeWithRand(rng *rand.Rand) ColorScheme {
	firstColor := generateRandomColor(rng)
	secondColor := generateSecondColor(firstColor)
	return ColorScheme{firstColor, secondColor}
}

// ColorToRGBA converts Color to an opaque color.RGBA.
func ColorToRGBA(c Color) color.RGBA {
	return color.RGBA{uint8(c.R), uint8(c.G), uint8(c.B), 255}
}
