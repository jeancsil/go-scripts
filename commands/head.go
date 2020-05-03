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

	reader := getReader(fileName)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for i := 0; i < *nLines; i++ {
		if !scanner.Scan() {
			break
		}
		fmt.Println(scanner.Text())
	}
}

func getReader(fileName *string) io.Reader {
	if *fileName != "" {
		f, err := os.Open(*fileName)
		if err != nil {
			panic(err)
		}

		validateInput(f)

		return f
	}

	f := os.Stdin
	validateInput(f)

	return f
}

func validateInput(f *os.File) {
	fi, err := f.Stat()
	if err != nil {
		fmt.Println("Error stat()")
		os.Exit(1)
	}

	if fi.Size() <= 0 {
		fmt.Println("input has no length")
		os.Exit(1)
	}
}
