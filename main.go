package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	if clipboard.Unsupported {
		fmt.Println("clipboard unsupported")
		os.Exit(1)
	}

	flag.Usage = func() {
		fmt.Println("usage: cb [options...] <action>")
		fmt.Println("  -h, --help  show this help menu")
		fmt.Println("")
		fmt.Println("actions:")
		fmt.Println("  copy (default): copy, c, input, in, i")
		fmt.Println("  paste         : paste, p, output, out, o")
	}
	flag.Parse()
	args := flag.Args()

	if len(args) > 1 {
		fmt.Println("too many arguments")
		os.Exit(1)
	}

	if len(args) == 0 {
		copy()
		return
	}

	action := args[0]

	switch action {
	case "copy", "c", "input", "in", "i":
		copy()
		return
	case "paste", "p", "output", "out", "o":
		paste()
		return
	default:
		fmt.Printf("unknown action: %s\n", action)
		os.Exit(1)
	}
}

func copy() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := clipboard.WriteAll(string(b)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func paste() {
	s, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print(s)
}
