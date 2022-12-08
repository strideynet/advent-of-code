package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func rangeFromString(rng string) (start, end int, err error) {
	values := strings.Split(rng, "-")
	if len(values) != 2 {
		return 0, 0, errors.New("Invalid number of values in range")
	}
	start, err = strconv.Atoi(values[0])
	if err != nil {
		return 0, 0, err
	}
	end, err = strconv.Atoi(values[1])
	if err != nil {
		return 0, 0, err
	}

	return start, end, nil
}

func run() error {
	f, err := os.Open("/Users/noahstride/aoc-day4-input.txt")
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println("Failed to close file:", err)
		}
	}()

	redundantPairsCount := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		elfs := strings.Split(scanner.Text(), ",")
		if len(elfs) != 2 {
			return errors.New("Incorrect count of elves")
		}
		firstStart, firstEnd, err := rangeFromString(elfs[0])
		if err != nil {
			return err
		}
		secondStart, secondEnd, err := rangeFromString(elfs[1])
		if err != nil {
			return err
		}

		if (firstStart >= secondStart && firstEnd <= secondEnd) || (secondStart >= firstStart && secondEnd <= firstEnd) {
			log.Println(firstStart, "-", firstEnd, ",", secondStart, "-", secondEnd)
			redundantPairsCount++
		}

	}

	log.Println(redundantPairsCount)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
