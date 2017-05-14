// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.Black, color.RGBA{0x4A, 0xF6, 0x26, 0xff}, color.RGBA{0xF6, 0x4A, 0x26, 0xff}}

const (
	whiteIndex  = 0 // first color in palette
	secondIndex = 1 // next color in palette
	thirdIndex  = 2 // third
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*math.Pi; t += res {
			x, y := math.Sin(t), math.Sin(t*freq+phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), secondIndex)
		}
		for t := cycles * math.Pi; t < cycles*2*math.Pi; t += res {
			x, y := math.Sin(t), math.Sin(t*freq+phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), thirdIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOT: ignoring encoding errors
}
