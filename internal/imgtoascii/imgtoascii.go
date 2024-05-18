package imgtoascii

import (
	"image"
	"image/color"
	"math"
	"strings"

	"github.com/disintegration/imaging"
)

type Options struct {
	Columns int
}

func Convert(img image.Image, options ...Options) (string, error) {
	option := mergeOptions(options...)

	cols := img.Bounds().Dx()
	if option.Columns != 0 {
		cols = option.Columns
	}
	rows := cols / 2

	var data strings.Builder
	data.Grow(cols * rows)

	resizedImg := imaging.Resize(img, cols, rows, imaging.Lanczos)

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			clr := resizedImg.At(x, y)
			gray := colorToGrayScale(clr)
			char := grayToChar(gray)
			data.WriteByte(char)
		}
		data.WriteByte('\n')
	}
	return data.String(), nil
}

func mergeOptions(options ...Options) *Options {
	return &options[0]
}

func colorToGrayScale(clr color.Color) uint8 {
	r, g, b, _ := clr.RGBA()
	return uint8((2126*r + 7152*g + 722*b) / 65536 / 256)
}

// chars from https://paulbourke.net/dataformats/asciiart/
var charlist = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "
var chars = []byte(charlist)

// var chars = [26]byte{'M', 'N', 'H', '#', 'Q', 'U', 'A', 'D', 'O', 'Y', '2', '6', '8', 'Z', '0', 'L', 'C', 'J', 'P', 'G', '9', 'S', 'I', ':', '!', ' '}
var charLength = uint8(math.Ceil(256 / float64(len(chars))))

func grayToChar(gray uint8) byte {
	return chars[gray/charLength]
}
