package goAvatar_test

import (
	"image/png"
	"os"
	"testing"

	"github.com/ethercod3/goAvatar"
)

func TestMain(t *testing.T) {
	options := goAvatar.AvatarOptions{
		Dimensions: 5,
		FileSizePx: 500,
	}
	img := goAvatar.GenerateAvatar(options)
	file, err := os.Create("./avatar.png")
	if err != nil {
		t.Fatal(err)
	}
	png.Encode(file, img)
}
