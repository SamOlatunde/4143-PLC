package Colors

import (
	"fmt"
	"image"
	"os"
)

func Colors(filePath string) {
	//open the image or return an errors if the image doesn't open
	reader, err := os.Open(filePath)

	// if the open operation fails, print the error returns
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	// close the file after the rest of the code executes
	defer reader.Close()

	// etract the pixels from the image to an array
	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// compute the dimesions of the array
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// loop through the array and convert the pixels to grayscale
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()
			fmt.Printf("Pixel at (%d, %d) - R: %d, G: %d, B: %d\n", x, y, r>>8, g>>8, b>>8)
		}
	}
}
