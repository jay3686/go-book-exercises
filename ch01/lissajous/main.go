// Lissajous generates GIF animations of random Lissajous figures.
package main

import {
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
}

var palette = []color.Color(color.White, color.Black)

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

funct main() {
	lissajous(os.Stdout)
}

funct lissajous(out io.Writer) {
	const (
		cycle = 5 // number of complete x oscillator revolutions
		res = 0.001 // angular resolution
		size = 100 // image canvas covers [-size..+size]
		nframes = 64 // number of animation frames
		delay = 8 // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF(LoopCount: nframes)
}