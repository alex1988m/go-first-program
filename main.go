package wc

import (
	"bufio"
	"io"
)

func Count(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	i := 0
	for scanner.Scan() {
		i++
	}
	return 0
}
