/// echo prints its command-line arguments
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

//Ex 1.1: print the name of the command
func exercise1_1() {
	fmt.Println(os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}

//Ex 1.2: Print index and value of each argument
func exercise1_2() {
	// format string docs here: https://gobyexample.com/string-formatting
	for i, v := range os.Args[1:] {
		fmt.Printf("%d\t%s", i, v)
	}
}

//Ex 1.3: time string concat solution and vs the join solution
func timeit(cmd func()) {
	start := time.Now()
	cmd()
	log.Printf("Finished, took %s", time.Since(start))

}

func echoSlow() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echoFast() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func exercise1_3() {
	fmt.Println("Testing with string concat")
	timeit(echoSlow)
	fmt.Println("Testing with string join")
	timeit(echoFast)

}
func main() {
	exercise1_1()
	exercise1_2()
	exercise1_3()
}
