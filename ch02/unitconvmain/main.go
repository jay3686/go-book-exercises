// unitconv converts its numeric argument to various units.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jay3686/go-book-exercises/ch02/tempconv"
	"github.com/jay3686/go-book-exercises/ch02/unitconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))

		p := unitconv.Pound(t)
		o := unitconv.Ounce(t)
		k := unitconv.Kilogram(t)

		fmt.Printf("%s = %s, %s = %s\n", p, unitconv.PToO(p), o, unitconv.OToP(o))
		fmt.Printf("%s = %s, %s = %s\n", p, unitconv.PToK(p), k, unitconv.KToP(k))
	}
}
