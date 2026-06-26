package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	entries, err := file.Readdir(0)
	if err != nil {
		return err
	}
	var filtered []os.FileInfo
	for _, entry := range entries {
		if entry.Name() == ".DS_Store" {
			continue
		}
		if !printFiles && !entry.IsDir() {
			continue
		}
		filtered = append(filtered, entry)
	}
	sort.Slice(filtered, func(i, j int) bool { return filtered[i].Name() < filtered[j].Name() })
	for i, entry := range filtered {
		prefix := "├───"
		if i == len(filtered)-1 {
			prefix = "└───"
		}
		fmt.Fprintln(out, prefix+entry.Name())
	}
	return nil
}
