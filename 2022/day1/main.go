package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func run() error {
	f, err := os.Open("/Users/noahstride/aoc-day1-input.txt")
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println("Failed to close file:", err)
		}
	}()

	maxCalCountSeen := 0
	currentCalCount := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// handle case where end of elf
		if line == "" {
			if currentCalCount > maxCalCountSeen {
				maxCalCountSeen = currentCalCount
			}
			currentCalCount = 0
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		currentCalCount += num
	}
	if currentCalCount > maxCalCountSeen {
		maxCalCountSeen = currentCalCount
	}

	log.Println("Highest caloried elf has", maxCalCountSeen, "calories")

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
