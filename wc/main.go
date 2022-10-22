package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type scanByteCounter struct {
	bytesRead int
}

func (s *scanByteCounter) wrap(split bufio.SplitFunc) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (int, []byte, error) {
		adv, tok, err := split(data, atEOF)
		s.bytesRead += adv
		return adv, tok, err
	}
}

func count(r io.Reader, countType string) int {
	sc := bufio.NewScanner(r)

	if countType == "bytes" {
		bc := scanByteCounter{}
		bc.bytesRead = 0
		splitFunc := bc.wrap(bufio.ScanWords)
		sc.Split(splitFunc)
		for sc.Scan() {
			// fmt.Printf("split text: %s|\n", sc.Text())
			// fmt.Printf("bytes read: %d\n\n", bc.bytesRead)
		}
		return bc.bytesRead
	}

	if countType == "word" {
		sc.Split(bufio.ScanWords)
	}

	wc := 0

	for sc.Scan() {
		wc++
	}

	return wc
}

func main() {
	var countType string

	flag.StringVar(&countType, "t", "word", "count type")
	flag.Parse()

	if countType == "line" {
		fmt.Println(count(os.Stdin, "line"))
	} else if countType == "bytes" {
		// must use `echo -n ...`, or the counter reads the newline as a character
		fmt.Println(count(os.Stdin, "bytes"))
	} else {
		fmt.Println(count(os.Stdin, "word"))
	}
}
