package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

const Size = 512

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: spinning <output.png>")
		os.Exit(1)
	}

	image := image.NewRGBA64(image.Rect(0, 0, Size, Size))

	for x := 0; x < Size; x++ {
		for y := 0; y < Size; y++ {
			distanceVecX := float64(x - Size/2)
			distanceVecY := float64(y - Size/2)
			mag := math.Sqrt(math.Pow(distanceVecX, 2) + math.Pow(distanceVecY, 2))
			angle := mag / (Size / 3)
			newX := math.Cos(angle)*distanceVecX - math.Sin(angle)*distanceVecY
			newY := math.Sin(angle)*distanceVecX + math.Cos(angle)*distanceVecY
			diffX := newX - distanceVecX
			diffY := newY - distanceVecY
			diffScaler := float64(0x8000) / Size
			image.Set(x, y, color.RGBA64{
				R: uint16(diffScaler*diffX + 0x8000),
				G: uint16(diffScaler*diffY + 0x8000),
				B: 0x8000,
				A: 0xffff,
			})
		}
	}

	out, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer out.Close()
	png.Encode(out, image)
}
