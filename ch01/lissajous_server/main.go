// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.Black, color.RGBA{0x4A, 0xF6, 0x26, 0xff}, color.RGBA{0xF6, 0x4A, 0x26, 0xff}}

const (
	whiteIndex  = 0 // first color in palette
	secondIndex = 1 // next color in palette
	thirdIndex  = 2 // third
)

var (
	cycles  = 5.0   // number of complete x oscillator revolutions
	res     = 0.001 // angular resolution
	size    = 100.0 // image canvas covers [-size..+size]
	nframes = 64    // number of animation frames
	delay   = 8     // delay between frames in 10ms units
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		fmt.Println(params)
		if params["cycles"] != nil {
			//ignoring error for now
			cycles, _ = strconv.ParseFloat(params["cycles"][0], 64)
		}
		if params["res"] != nil {
			//ignoring error for now
			res, _ = strconv.ParseFloat(params["res"][0], 64)
		}
		if params["size"] != nil {
			//ignoring error for now
			size, _ = strconv.ParseFloat(params["size"][0], 64)
		}
		if params["nframes"] != nil {
			//ignoring error for now
			nframes, _ = strconv.Atoi(params["nframes"][0])
		}
		if params["delay"] != nil {
			//ignoring error for now
			delay, _ = strconv.Atoi(params["delay"][0])
		}
		lissajous(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*int(size)+1, 2*int(size)+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*math.Pi; t += res {
			x, y := math.Sin(t), math.Sin(t*freq+phase)
			img.SetColorIndex(int(size)+int(x*size+0.5), int(size)+int(y*size+0.5), secondIndex)
		}
		for t := cycles * math.Pi; t < cycles*2*math.Pi; t += res {
			x, y := math.Sin(t), math.Sin(t*freq+phase)
			img.SetColorIndex(int(size)+int(x*size+0.5), int(size)+int(y*size+0.5), thirdIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOT: ignoring encoding errors
}
