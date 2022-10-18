package utils

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
)

type GBImage struct {
	Rows    int
	Columns int
	Source  []byte
	Palette color.Palette
}

// PNG 형식의 이미지를 디코딩하여 로딩
func ReadGBImage(pngFilename string) (*GBImage, bool) {
	file, err := os.Open(pngFilename)
	if err != nil {
		return nil, false
	}

	defer file.Close()

	file.Seek(0, 0)
	palette, ok := decodePNGPalette(file, 4)
	if !ok {
		return nil, false
	}

	file.Seek(0, 0)
	src, r, c, ok := decodePNGSource(file, palette)
	if !ok {
		return nil, false
	}

	return &GBImage{
		Source:  src,
		Palette: palette,
		Rows:    r,
		Columns: c,
	}, true
}

func decodePNGSource(reader io.Reader, palette color.Palette) ([]byte, int, int, bool) {
	var bytes []byte

	src, err := png.Decode(reader)
	if err != nil {
		return nil, 0, 0, false
	}

	size := src.Bounds().Size()
	rows := size.Y / 8
	columns := size.X / 8

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			for y := 0; y < 8; y++ {
				first := byte(0)
				second := byte(0)

				for x := 0; x < 8; x++ {
					indexX := (r * 8) + x
					indexY := (c * 8) + y
					color := src.At(indexX, indexY)
					colorIndex := palette.Index(color)
					hi := (colorIndex & 0b00000010) >> 1
					lo := (colorIndex & 0b00000001)
					first |= byte(lo << (7 - x))
					second |= byte(hi << (7 - x))
				}

				bytes = append(bytes, first, second)
			}
		}
	}

	return bytes, rows, columns, true
}

func decodePNGPalette(reader io.Reader, paletteCount int) (color.Palette, bool) {
	const (
		pngHeader    = "\x89PNG\r\n\x1a\n"
		paletteLabel = "PLTE"
	)
	var result color.Palette

	bytes, _ := io.ReadAll(reader)

	// PNG 헤더 체크
	if string(bytes[:len(pngHeader)]) != pngHeader {
		return nil, false
	}

	// PLTE 영역 파싱
	currentIndex := 0
	for {
		if currentIndex+len(paletteLabel) >= len(bytes) {
			return nil, false
		}

		if string(bytes[currentIndex:currentIndex+len(paletteLabel)]) == paletteLabel {
			currentIndex += 4
			break
		}
		currentIndex++
	}

	for i := 0; i < paletteCount; i++ {
		color := color.RGBA{
			bytes[currentIndex+0],
			bytes[currentIndex+1],
			bytes[currentIndex+2],
			0xff,
		}
		result = append(result, color)
		currentIndex += 3
	}

	return result, true
}

// 이미지를 PNG 형식으로 인코딩하여 저장
func WriteGBImage(filename string, gbImage GBImage) bool {
	paletted := image.NewPaletted(image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: gbImage.Columns * 8,
			Y: gbImage.Rows * 8,
		},
	}, gbImage.Palette)

	for i := 0; i < len(gbImage.Source); i += 2 {
		first := gbImage.Source[i+0]
		second := gbImage.Source[i+1]

		x := (i / (16 * gbImage.Columns)) * 8
		y := (i % (16 * gbImage.Columns)) / 2

		for pixelX := 0; pixelX < 8; pixelX++ {
			hi := byte((second >> (7 - pixelX)) & 1)
			lo := byte((first >> (7 - pixelX)) & 1)
			colorIndex := (hi << 1) | lo
			paletted.SetColorIndex(x+pixelX, y, colorIndex)
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return false
	}

	defer file.Close()
	return png.Encode(file, paletted) == nil
}
