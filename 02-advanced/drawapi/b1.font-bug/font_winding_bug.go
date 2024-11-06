package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
	"os"

	fixedtruetype "github.com/goki/freetype/truetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	ft      *truetype.Font
	fixedft *fixedtruetype.Font
)

func RenderBad() {
	s := "HelloWorld"
	w, h := 256, 128
	// draw a white image as background
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(img, img.Bounds(), image.NewUniform(color.White), image.Point{}, draw.Src)

	// draw text on to the background
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.Black),
		Face: truetype.NewFace(ft, &truetype.Options{Size: 40, DPI: 72, Hinting: font.HintingNone}),
	}
	// align center
	bd, _ := d.BoundString(s)
	sx := (w - bd.Max.X.Ceil()) / 2
	sy := (h-(bd.Max.Y.Ceil()-bd.Min.Y.Ceil()))/2 + (bd.Max.Y.Ceil() - bd.Min.Y.Ceil())
	d.Dot = fixed.Point26_6{X: fixed.I(sx), Y: fixed.I(sy)}
	d.DrawString(s)

	// create image file
	saveImg("bad.png", img)
}

func RenderWell() {
	s := "HelloWorld"
	w, h := 256, 128
	// draw a white image as background
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(img, img.Bounds(), image.NewUniform(color.White), image.Point{}, draw.Src)

	// draw text on to the background
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.Black),
		Face: fixedtruetype.NewFace(fixedft, &fixedtruetype.Options{Size: 40, DPI: 72, Hinting: font.HintingNone}),
	}
	// align center
	bd, _ := d.BoundString(s)
	sx := (w - bd.Max.X.Ceil()) / 2
	sy := (h-(bd.Max.Y.Ceil()-bd.Min.Y.Ceil()))/2 + (bd.Max.Y.Ceil() - bd.Min.Y.Ceil())
	d.Dot = fixed.Point26_6{X: fixed.I(sx), Y: fixed.I(sy)}
	d.DrawString(s)

	// save
	saveImg("good.png", img)
}

func saveImg(name string, img image.Image) {
	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		log.Fatalf("Failed to encode image: %v", err)
	}
}

func init() {
	f, err := os.Open("./Lexend-Bold.ttf")
	if err != nil {
		panic(fmt.Errorf("font not found"))
	}
	bs, _ := io.ReadAll(f)
	ft, err = truetype.Parse(bs)
	if err != nil {
		panic(fmt.Errorf("load font error: %v", err))
	}
	fixedft, err = fixedtruetype.Parse(bs)
	if err != nil {
		panic(fmt.Errorf("load font error: %v", err))
	}
}
