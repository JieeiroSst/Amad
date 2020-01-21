package images

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

func loadImages(fileName string) image.Image {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

// func getPixel(img image.Image) []dto.Pixel {
// 	bounds := img.Bounds()
// 	fmt.Println(bounds.Dx(), "x", bounds.Dy())
// 	pixels := make([]dto.Pixel, bounds.Dx()*bounds.Dy())

// 	for i := 0; i < bounds.Dx()*bounds.Dy(); i++ {
// 		x := i % bounds.Dx()
// 		y := i / bounds.Dx()
// 		r, g, b, a := img.At(x, y).RGBA()
// 		pixels[i].r = r
// 		pixels[i].g = g
// 		pixels[i].b = b
// 		pixels[i].a = a
// 	}
// 	return pixels
// }

// func getimages(dir string) [][]dto.Pixel {
// 	var images [][]dto.Pixel

// 	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
// 		if info.IsDir() {
// 			return nil
// 		}
// 		img := loadImages(path)
// 		pixels := getPixel(img)
// 		images := append(images, pixels)
// 		return nil
// 	})
// 	return images
// }

// func Load() {
// 	images := getimages("../../uploads/")
// 	for i,img=images {
// 		for j,pixel=range img {
// 			fmt.Println("images",i,"\t pixel",j,"\t r g b a",pixel)
// 			if j==10{
// 				break
// 			}
// 		}
// 	}
// }
