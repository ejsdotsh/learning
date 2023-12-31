package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/ejsdotsh/learning-generative-art/sketch"
)

var (
	sourceImageName = "source.jpg"
	outputImageName = "out.png"
	totalCycleCount = 5000
)

// func cpuProf(fn func()) {
// 	f, err := os.Create("cpuprof.out")
// 	if err != nil {
// 		fmt.Println("Error in creating file for writing cpu profile: ", err)
// 		return
// 	}
// 	defer f.Close()

// 	if err := pprof.StartCPUProfile(f); err != nil {
// 		fmt.Println("Error starting CPU profile: ", err)
// 		return
// 	}

// 	fn()
// 	defer pprof.StopCPUProfile()
// }

// uncomment for a random image
// func loadRandomUnsplashImage(width, height int) (image.Image, error) {
// 	url := fmt.Sprintf("https://source.unsplash.com/random/%dx%d", width, height)
// 	res, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	img, _, err := image.Decode(res.Body)
// 	return img, err
// }

func loadImage(src string) (image.Image, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, fmt.Errorf("source image could not be loaded: %w", err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("source image format could not be decoded: %w", err)
	}

	return img, nil
}

func saveOutput(img image.Image, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	// img, err := loadRandomUnsplashImage(2000, 2000)
	img, err := loadImage(sourceImageName)
	if err != nil {
		log.Panicln(err)
	}

	destWidth := 2000
	s := sketch.NewSketch(img, sketch.UserParams{
		DestWidth:                destWidth,
		DestHeight:               2000,
		StrokeRatio:              0.75,
		StrokeReduction:          0.002,
		StrokeInversionThreshold: 0.5,
		StrokeJitter:             int(0.1 * float64(destWidth)),
		InitialAlpha:             0.1,
		AlphaIncrease:            0.06,
		MinEdgeCount:             5,
		MaxEdgeCount:             16,
	})

	rand.Seed(time.Now().Unix())

	for i := 0; i < totalCycleCount; i++ {
		s.Update()
	}

	saveOutput(s.Output(), outputImageName)
}
