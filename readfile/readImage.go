/*
--------------------------------------------------------------

	student name: Souhail Daoudi
	student number: 300135458

--------------------------------------------------------------
*/
package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"math"
	"os"
	"sync"
)

// var imagePath string = "C:/Users/souhail/OneDrive/Bureau/queryImages/q00.jpg"
var coun = 0
var sum float64 = 0

type Histo struct {
	Name string
	H    []float64
}

func computeHistogram(imagePath string, depth int, wg *sync.WaitGroup) (Histo, error) {

	defer wg.Done()

	file, err := os.Open(imagePath)
	if err != nil {

		return Histo{"", nil}, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {

		return Histo{"", nil}, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	exponent := depth * 3
	capacityHisto := int(math.Pow(2, float64(exponent)))

	//define capacity of histograme
	h := Histo{imagePath, make([]float64, capacityHisto)}
	// avoid negatif shifting
	if depth > 8 {
		depth = 8
	}
	for y := 0; y < height && y < height; y++ {
		for x := 0; x < width && x < width; x++ {

			red, green, blue, _ := img.At(x, y).RGBA()
			red = red >> 8 >> (8 - depth)
			green = green >> 8 >> (8 - depth)
			blue = blue >> 8 >> (8 - depth)

			//print rgb values
			//fmt.Printf("Pixel at (%d, %d): R=%d, G=%d, B=%d\n", x, y, red, green, blue)

			//create histogram
			indexHistogramme := ((red << (2 * depth)) + (green << depth) + blue)
			h.H[indexHistogramme]++

		}
	}

	//normalize histogram
	for o := 0; o < capacityHisto; o++ {
		h.H[o] /= float64(width * height)

	}

	return h, nil
}

// compute data histograms , function computeHistograms call function compute histogram
func computeHistograms(imagePath []string, depth int, hChan chan<- Histo, wg *sync.WaitGroup) {
	defer wg.Done()

	// send each image path of data filenames to computeHistogram
	for i := range imagePath {
		wg.Add(1)
		go func(path string) {
			histo, err := computeHistogram(imagePath[i], depth, wg)
			if err != nil {
				fmt.Print("erreur")
			}
			hChan <- histo
		}(imagePath[i])
	}

}
