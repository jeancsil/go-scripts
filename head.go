package main

import(
  "fmt"
  "os"
  "bufio"
  "io"
  "flag"
)

func main() {
  var nLines = flag.Int("n", 10, "the number of lines to read")
  flag.Parse()

  fmt.Println("n:", *nLines)
  fmt.Println("tail:", flag.Args())
  fmt.Println("Length:", len(flag.Args()))
  fmt.Println("NArg:", flag.NArg())

  file := getReader()

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)
  var lines []string

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  for i := 0; i < *nLines; i++ {
    fmt.Println(lines[i])
  }
}

func getReader() io.Reader {
  var err error
  r := os.Stdin

  fileName := ""
  for i := 0; i < len(flag.Args()); i++ {
    if _, err := os.Stat(flag.Args()[i]); err == nil {
      fileName = flag.Args()[i]
      break
    }
  }

  if fileName != "" {
    r, err = os.Open(fileName)
    if err != nil {
      panic(err)
    }
  }

  return r
}
