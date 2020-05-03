package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var nLines = flag.Int("n", 10, "the number of lines to read")
	var fileName = flag.String("f", "", "filename to read")
	flag.Parse()

	file := getReader(fileName)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines); i++ {
		if i == *nLines {
			break
		}
		fmt.Println(lines[i])
	}
}

func getReader(fileName *string) io.Reader {
	r := os.Stdin
	var err error

	if *fileName != "" {
		r, err = os.Open(*fileName)
		if err != nil {
			panic(err)
		}
	}

	return r
}
