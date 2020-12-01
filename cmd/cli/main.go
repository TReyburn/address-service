package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/TReyburn/address-service/parser"
)

func main() {
	fn := os.Args[1]
	fh, err := os.Open(fn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res := parseFile(fh)

	c := make(chan parser.Address)

	for _, x := range res {
		go parser.ParseAddress(x, c)
	}

	for i := 1; i < len(res); i++ {
		fmt.Println(<-c)
	}

}

func parseFile(fh *os.File) []string {

	scanner := bufio.NewScanner(fh)

	content := make([]string, 0)

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	return content
}
