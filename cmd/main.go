package main

import (
	"bufio"
	"fmt"
	"os"

	"../parser"
)

func main() {
	fn := os.Args[1]
	res := parseFile(fn)

	c := make(chan string)

	for _, x := range res {
		go parser.ParseAddress(x, c)
	}

	for i := 1; i < len(res); i++ {
		_ = []string{<-c}
	}

}

func parseFile(fn string) []string {
	fh, err := os.Open(fn)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(fh)

	content := make([]string, 0)

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	return content
}
