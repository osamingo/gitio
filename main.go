package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/osamingo/gitio/shortener"
)

var code = ""

func init() {
	flag.StringVar(&code, "code", "", "if you will be use any code, set code flag")
}

func main() {

	flag.Parse()

	as := flag.Args()
	if len(as) < 1 {
		PrintUsage()
		return
	}

	u, err := url.ParseRequestURI(as[0])
	if err != nil {
		ErrExit(err)
	}

	r, err := shortener.GenerateShortURL(u, code)
	if err != nil {
		ErrExit(err)
	}

	fmt.Fprintln(os.Stdout, r)
}

// PrintUsage shows usage sentence.
func PrintUsage() {
	fmt.Fprintln(os.Stdout, "Usage: gitio [-code=] url\nIf you will be use any code, set code flag")
}

// ErrExit exits with error.
func ErrExit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
