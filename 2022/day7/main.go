package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int
}

type Directory struct {
	Name string
	// Subdirectories is a slice of all directories below this directory.
	Subdirectories []*Directory
	// Files is a slice of all files contained directly by this directory.
	Files []*File
	// Parent is the directory that contains this directory. For the root
	// directory this will be nil.
	Parent *Directory
}

// Size returns the total size of the directory, including the size of the
// subdirectories
func (d *Directory) Size() int {
	size := 0
	for _, file := range d.Files {
		size += file.Size
	}
	for _, directory := range d.Subdirectories {
		size += directory.Size()
	}
	return size
}

func (d *Directory) SumSub100() int {
	sum := 0
	if size := d.Size(); size <= 100000 {
		sum += size
	}
	for _, sub := range d.Subdirectories {
		sum += sub.SumSub100()
	}
	return sum
}

const (
	cdCommandPrefix = "$ cd "
	lsCommand       = "$ ls"
)

func buildTree(f io.Reader) (root *Directory, err error) {
	var currentDirectory *Directory
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, cdCommandPrefix) {
			dirName := strings.TrimPrefix(line, cdCommandPrefix)
			// Handle moving upwards
			if dirName == ".." {
				currentDirectory = currentDirectory.Parent
				continue
			}
			// Handle moving to a new subdirectory
			newDir := &Directory{
				Name:   dirName,
				Parent: currentDirectory,
			}
			if currentDirectory != nil {
				currentDirectory.Subdirectories = append(
					currentDirectory.Subdirectories, newDir,
				)
			}
			if root == nil {
				root = newDir
			}
			currentDirectory = newDir
		} else if strings.HasPrefix(line, lsCommand) {
			// We can actually ignore the ls, it's pretty meaningless.
			continue
		} else {
			split := strings.Split(line, " ")
			size := split[0]
			name := split[1]
			if size == "dir" {
				// Ignore directories
				continue
			}
			sizeInt, err := strconv.Atoi(size)
			if err != nil {
				return nil, err
			}
			currentDirectory.Files = append(currentDirectory.Files, &File{
				Name: name,
				Size: sizeInt,
			})
		}
	}

	return root, nil
}

func run() error {
	f, err := os.Open("/Users/noahstride/aoc-day7-input.txt")
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println("Failed to close file:", err)
		}
	}()

	// Break the input down into a tree.
	directory, err := buildTree(f)
	if err != nil {
		return err
	}

	// Traverse the tree looking for sub100K directories
	fmt.Println(directory.SumSub100())
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
