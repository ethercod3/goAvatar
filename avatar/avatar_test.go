package avatar

import (
	"bytes"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

func TestOptionsValidate(t *testing.T) {
	tests := []struct {
		name    string
		options Options
		wantErr bool
	}{
		{
			name:    "valid options",
			options: Options{Dimensions: 5, FileSizePx: 500},
		},
		{
			name:    "dimensions too small",
			options: Options{Dimensions: 1, FileSizePx: 500},
			wantErr: true,
		},
		{
			name:    "file size too small",
			options: Options{Dimensions: 5, FileSizePx: 0},
			wantErr: true,
		},
		{
			name:    "file size smaller than dimensions",
			options: Options{Dimensions: 10, FileSizePx: 5},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.options.Validate()
			if tt.wantErr && err == nil {
				t.Fatal("expected error")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestGenerateReturnsExpectedBounds(t *testing.T) {
	options := Options{
		Dimensions: 5,
		FileSizePx: 500,
	}

	img := Generate(options)
	want := image.Rect(0, 0, options.FileSizePx, options.FileSizePx)
	if img.Bounds() != want {
		t.Fatalf("bounds = %v, want %v", img.Bounds(), want)
	}
}

func TestNewValidatesOptions(t *testing.T) {
	_, err := New(Options{Dimensions: 1, FileSizePx: 500})
	if err == nil {
		t.Fatal("expected validation error")
	}
}

func TestGenerateWithSeedIsDeterministic(t *testing.T) {
	options := Options{
		Dimensions: 5,
		FileSizePx: 500,
	}

	first := GenerateWithSeed(options, 42)
	second := GenerateWithSeed(options, 42)

	if !bytes.Equal(first.Pix, second.Pix) {
		t.Fatal("expected same seed to generate same image")
	}
}

func TestGenerateWithSeedZeroIsDeterministic(t *testing.T) {
	options := Options{
		Dimensions: 5,
		FileSizePx: 500,
	}

	first := GenerateWithSeed(options, 0)
	second := GenerateWithSeed(options, 0)

	if !bytes.Equal(first.Pix, second.Pix) {
		t.Fatal("expected zero seed to be deterministic when passed explicitly")
	}
}

func TestSaveWritesPNG(t *testing.T) {
	path := filepath.Join(t.TempDir(), "avatar.png")
	err := Save(path, Options{
		Dimensions: 5,
		FileSizePx: 64,
		Seed:       42,
	})
	if err != nil {
		t.Fatalf("save failed: %v", err)
	}

	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("open saved file: %v", err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		t.Fatalf("decode saved png: %v", err)
	}

	want := image.Rect(0, 0, 64, 64)
	if img.Bounds() != want {
		t.Fatalf("bounds = %v, want %v", img.Bounds(), want)
	}
}
