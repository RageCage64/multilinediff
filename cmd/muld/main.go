package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"

	"github.com/RageCage64/multilinediff"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if (len(os.Args) - 1) != 2 {
		return fmt.Errorf("requires exactly two arguments, got %d", len(os.Args))
	}

	leftPath := os.Args[1]
	rightPath := os.Args[2]

	leftContent, err := readContent(leftPath)
	if err != nil {
		return err
	}
	rightContent, err := readContent(rightPath)
	if err != nil {
		return err
	}

	lineSep := "\n"
	if runtime.GOOS == "windows" {
		lineSep = "\r\n"
	}

	fmt.Println(multilinediff.Diff(leftContent, rightContent, lineSep))

	return nil
}

func readContent(path string) (string, error) {
	if path == "-" {
		return readFromStdin()
	}
	content, err := os.ReadFile(path)
	return string(content), err
}

func readFromStdin() (string, error) {
	stdin := bufio.NewReader(os.Stdin)
	data := []byte{}
	for {
		b, err := stdin.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return "", err
			}
		}
		data = append(data, b)
	}
	return string(data), nil
}
