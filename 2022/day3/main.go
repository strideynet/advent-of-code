package main

import (
	"bufio"
	"errors"
	"log"
	"os"
)

func findShared(a []int, b []int) (int, error) {
	for _, item := range a {
		for _, secondItem := range b {
			if item == secondItem {
				return item, nil
			}
		}
	}

	return 0, errors.New("no match")
}

func run() error {
	f, err := os.Open("/Users/noahstride/aoc-day3-input.txt")
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println("Failed to close file:", err)
		}
	}()

	prioritySum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		backpack := scanner.Text()
		if backpack == "" {
			continue
		}
		// Break down input into items
		items := []int{}
		for _, item := range backpack {
			items = append(items, int(item))
		}
		// Split into halves
		firstHalf := items[:len(items)/2]
		secondHalf := items[len(items)/2:]
		// Find shared item
		sharedItem, err := findShared(firstHalf, secondHalf)
		if err != nil {
			return err
		}
		// Find value of shared item
		value := 0
		if sharedItem >= 96 && sharedItem <= 122 { // a-z
			value = (sharedItem - 96)
		} else if sharedItem >= 65 && sharedItem <= 90 { // A-Z
			value = (sharedItem - 64) + 26
		} else {
			return errors.New("unrecognized char")
		}
		prioritySum += value
		log.Println("Backpack", backpack, "Shared", string(sharedItem), "Shared value", value)
	}
	log.Println("Value:", prioritySum)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
