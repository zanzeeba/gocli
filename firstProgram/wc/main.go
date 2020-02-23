package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// add a line and a few words
	lines := flag.Bool("l",false, "Count Lines")
	bytes := flag.Bool("b",false, "Count Bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int  {
	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	} else if !countBytes {
		scanner.Split(bufio.ScanBytes)
	} 
	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}