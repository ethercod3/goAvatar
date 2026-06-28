package avatar

import (
	"image"
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
