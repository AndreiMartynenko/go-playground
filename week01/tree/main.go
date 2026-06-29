package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func main() {
}

// preparation for recursion
func walkDir(out io.Writer, path string, printFiles bool, prefix string) error {
	//open a folder
	file, err := os.Open(path)

	if err != nil {
		return err
	}
	defer file.Close()

	//read the content
	content, err := file.Readdir(0)
	if err != nil {
		return err
	}
	//don't show/include ".DS_Store"
	var entries []os.FileInfo
	for _, item := range content {
		if item.Name() == ".DS_Store" {
			continue
		}
		// don't show files only folders
		if !printFiles && !item.IsDir() {
			continue
		}
		entries = append(entries, item)
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})
	//let's add some graphic
	for i, entry := range entries {
		connector := "├───"
		if i == len(entries)-1 {
			connector = "└───"
		}

		name := entry.Name()
		if printFiles && !entry.IsDir() {
			if entry.Size() == 0 {
				name += " (empty)"
			} else {
				name += fmt.Sprintf(" (%db)", entry.Size())
			}
		}

		fmt.Fprintln(out, prefix+connector+name)
		//fmt.Fprintln(out, entry.Name())
		nextPath := filepath.Join(path, entry.Name())
		//fmt.Fprintln(out, nextPath)

		nextPrefix := prefix
		if i == len(entries)-1 {
			nextPrefix += "\t"
		} else {
			nextPrefix += "│\t"
		}

		if entry.IsDir() {
			err := walkDir(out, nextPath, printFiles, nextPrefix)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// wrapper
// It doesn't do anything itself. It simply runs the actual algorithm.
func dirTree(out io.Writer, path string, printFiles bool) error {
	return walkDir(out, path, printFiles, "")

}
