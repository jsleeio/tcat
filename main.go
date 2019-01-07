package main

import (
	"bufio"
	"compress/bzip2"
	"compress/gzip"
  "flag"
  "fmt"
	"io"
  "log"
	"os"
	"path"
  "reflect"
)

// TransparentExpandingReader creates a Reader that transparently decompresses based
// on filename. Supports gzip and bzip2; falls back to assuming uncompressed
func TransparentExpandingReader(key string, source io.ReadCloser) (io.Reader, error) {
	ext := path.Ext(key)
	var reader io.Reader
	var err error
	switch {
	case ext == ".gz":
		reader, err = gzip.NewReader(source)
		if err != nil {
			return nil, err
		}
	case ext == ".bz2":
		reader = bzip2.NewReader(source)
	default:
		reader = bufio.NewReader(source)
	}
	return reader, nil
}

func main() {
  summary := flag.Bool("summary", false, "summarize file types and line counts only")
  flag.Parse()
  for _,arg := range flag.Args() {
    f,err := os.Open(arg)
    defer f.Close()
    if err != nil {
      log.Printf("error opening %s: %v", arg, err)
      continue
    }
    reader,err := TransparentExpandingReader(arg, f)
    if err != nil {
      log.Printf("error reading %s: %v", arg, err)
      continue
    }
    scanner := bufio.NewScanner(reader)
    n := 0
    for scanner.Scan() {
      n++
      if ! *summary {
        fmt.Println(scanner.Text())
      }
    }
    if *summary {
      fmt.Printf("%20s %10d %s\n", reflect.TypeOf(reader), n, arg)
    }
  }
}
