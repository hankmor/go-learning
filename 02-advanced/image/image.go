package main

import (
	"github.com/disintegration/imaging"
)

const (
	// image = "/Users/sam/Desktop/test/1.jpeg"
	image = "/Users/sam/Desktop/test/2.png"
)

func main() {
	// imaging 库：https://pkg.go.dev/github.com/disintegration/imaging
	println("打开image...")
	srcImage, err := imaging.Open(image)
	if err != nil {
		panic(err)
	}
	println("图片等比例缩放...")
	// Resize the image to width = 200px preserving the aspect ratio.
	srcImage = imaging.Resize(srcImage, 200, 0, imaging.Lanczos)

	println("保存image...")
	err = imaging.Save(srcImage, "/Users/sam/Desktop/test/2-1.jpg")
	if err != nil {
		panic(err)
	}
}
