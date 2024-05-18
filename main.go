package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	isLine := flag.Bool("l", false, "count lines")
	flag.Parse()
	fmt.Println(Count(os.Stdin, *isLine))
}

func Count(r io.Reader, isLine bool) int {
	scanner := bufio.NewScanner(r)
	if !isLine {
		scanner.Split(bufio.ScanWords)
	}
	i := 0
	for scanner.Scan() {
		i++
	}
	return i
}
