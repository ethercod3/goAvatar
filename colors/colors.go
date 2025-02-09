package colors

import (
	"fmt"
	"math/rand"
)

type ColorChannel uint8

type Color struct {
	R, G, B ColorChannel
}

type ColorScheme struct {
	First, Second Color
}

func generateRandomColorChannel() ColorChannel {
	return ColorChannel(rand.Intn(256))
}

func generateRandomColor() Color {
	r := generateRandomColorChannel()
	g := generateRandomColorChannel()
	b := generateRandomColorChannel()
	return Color{r, g, b}
}

func getLuminance(c Color) float32 {
	return (0.2126)*float32(c.R) + (0.7152)*float32(c.G) + (0.0722)*float32(c.B)
}

func darken(channel ColorChannel) ColorChannel {
	return channel - ColorChannel(float32(channel)*0.5)
}

func lighten(channel ColorChannel) ColorChannel {
	return channel + ColorChannel(float32(channel)*0.5)
}

func filter(c Color, filter string) Color {
	switch filter {
	case "darken":
		return Color{R: darken(c.R), G: darken(c.G), B: darken(c.B)}
	case "lighten":
		return Color{R: lighten(c.R), G: lighten(c.G), B: lighten(c.B)}
	default:
		panic(fmt.Sprintf("Unknown filter: %v", filter))
	}
}

func generateSecondColor(c Color) Color {
	luminance := getLuminance(c)
	if luminance <= 128 {
		return filter(c, "lighten")
	} else {
		return filter(c, "darken")
	}
}

func GenerateColorScheme() ColorScheme {
	firstColor := generateRandomColor()
	secondColor := generateSecondColor(firstColor)
	return ColorScheme{firstColor, secondColor}
}
