package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(Count(os.Stdin))
}

func Count(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	i := 0
	for scanner.Scan() {
		i++
	}
	return i
}
