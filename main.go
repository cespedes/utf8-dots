package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"

	"golang.org/x/image/font/basicfont"
)

func readImage(filename string) image.Image {
	var img image.Image
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	img, err = png.Decode(file)
	return img
}

func textToImage(text []string) image.Image {
	str := ""
	for i := range text {
		if i > 0 {
			str += " "
		}
		str += text[i]
	}

	face := basicfont.Face7x13
	img := image.NewAlpha(image.Rect(0, 0, len(str)*face.Advance, face.Height))
	for i, ru := range str {
		for _, ra := range face.Ranges {
			if ru >= ra.Low && ru <= ra.High {
				offset := ra.Offset + int(ru-ra.Low)
				for x := 0; x < face.Width; x++ {
					for y := 0; y < face.Height; y++ {
						img.Set(i*face.Advance+x, y, face.Mask.At(x, offset*face.Height+y))
					}
				}
			}
		}
	}
	return img
}

func imageToBraille(img image.Image) []string {
	braille := [][]rune{
		{0x2801, 0x2802, 0x2804, 0x2840},
		{0x2808, 0x2810, 0x2820, 0x2880},
	}
	result := make([]string, 0)
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y += 4 {
		var line string
		for x := bounds.Min.X; x < bounds.Max.X; x += 2 {
			var char rune
			for i := range braille {
				for j := range braille[i] {
					r, g, b, _ := img.At(x+i, y+j).RGBA()
					if (r > 64 || g > 64 || b > 64) != invert {
						char |= braille[i][j]
					}
				}
			}
			if char == 0 {
				char = ' '
			}
			line += string(char)
		}
		result = append(result, line)
	}
	return result
}

var invert bool

func init() {
	flag.BoolVar(&invert, "i", false, "invert colors")
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [-i] <image.png-or-text>\n", os.Args[0])
		os.Exit(1)
	}
	var img image.Image
	if len(flag.Args()) == 1 {
		img = readImage(flag.Arg(0))
	}
	if img == nil {
		img = textToImage(flag.Args())
	}
	if img == nil {
		// Shouldn't happen
		panic("No image?")
	}
	res := imageToBraille(img)
	for _, s := range res {
		fmt.Println(s)
	}
}
