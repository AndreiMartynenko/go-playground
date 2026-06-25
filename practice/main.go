package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// uniq reads sorted lines from input, removes duplicate neighbouring lines,
// and writes unique lines to output.
func uniq(input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)

	var previous string

	for scanner.Scan() {
		text := scanner.Text()

		if text == previous {
			continue
		}

		if text < previous {
			return fmt.Errorf("input is not sorted")
		}

		previous = text

		fmt.Fprintln(output, text)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := uniq(os.Stdin, os.Stdout); err != nil {
		panic(err)
	}
}
