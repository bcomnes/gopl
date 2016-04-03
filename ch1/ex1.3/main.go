package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// whatevs

	start := time.Now()
	echo1()
	ns := time.Since(start).Nanoseconds()
	fmt.Printf("echo1: %dns\n", ns)

	start = time.Now()
	echo2()
	ns = time.Since(start).Nanoseconds()
	fmt.Printf("echo2: %dns\n", ns)

	start = time.Now()
	echo3()
	ns = time.Since(start).Nanoseconds()
	fmt.Printf("echo3: %dns\n", ns)

}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
