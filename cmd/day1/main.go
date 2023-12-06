package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := os.Open("./input.txt")
	if err != nil {
		panic(fmt.Sprintf("couldn't open file: %s", err))
	}
	defer input.Close()

	fileContent := bufio.NewScanner(input)
	fileContent.Split(bufio.ScanLines)

	fileLines := []string{}

	for fileContent.Scan() {
		fileLines = append(fileLines, fileContent.Text())
	}

	part1total, err := Part1(fileLines)
	if err != nil {
		panic(fmt.Sprintf("part 1 failed: %s", err))
	}

	fmt.Printf("Part1:\nThe total sum of all the calibration values is: %d\n", part1total)

	part2total, err := Part2(fileLines)
	if err != nil {
		panic(fmt.Sprintf("part 2 failed: %s", err))
	}

	fmt.Printf("Part2:\nThe total sum of all the calibration values is: %d\n", part2total)

}

func Part2(fileLines []string) (int, error) {
	var err error
	total := 0

	for _, line := range fileLines {
		result := FindNumbersInText(line)

		lastSplit := strings.Split(result, "")

		firstAndLast := lastSplit[0] + lastSplit[len(lastSplit)-1]

		number, err := strconv.Atoi(firstAndLast)
		if err != nil {
			return 0, err
		}

		total = total + number
	}

	return total, err
}

func FindNumbersInText(text string) string {

	//chopped up number names to make it possible to parse out numbers that share a letter...
	//ugly solve, but works
	writtenNumbers := []string{"one", "tw", "thre", "four", "fiv", "six", "seven", "igh", "nin"}

	for i, n := range writtenNumbers {

		splitString := strings.SplitAfter(text, n)

		for _, sp := range splitString {
			if strings.Contains(text, sp) {
				text = strings.ReplaceAll(text, n, fmt.Sprint(i+1))
			}
		}

	}

	numbers := ""
	splitAgain := strings.Split(text, "")
	for _, v := range splitAgain {
		_, err := strconv.Atoi(v)
		if err != nil {
			//not an integer, check next
			continue
		} else {
			numbers = numbers + v
		}
	}

	return numbers

}

func Part1(fileLines []string) (int, error) {
	var err error

	numbersToAdd := []int{}

	for _, n := range fileLines {

		characters := strings.Split(n, "")

		var allNumbers string

		for _, c := range characters {
			_, err := strconv.Atoi(c)
			if err != nil {
				//not an integer, check next
				continue
			} else {
				allNumbers = allNumbers + c
			}
		}

		firstAndLast := strings.Split(allNumbers, "")

		n = firstAndLast[0] + firstAndLast[len(firstAndLast)-1]

		number, err := strconv.Atoi(n)
		if err != nil {
			return 0, fmt.Errorf("couldn't parse integer from string: %s", n)
		}

		numbersToAdd = append(numbersToAdd, number)

	}

	total := 0
	for _, n := range numbersToAdd {
		total += n
	}

	return total, err
}
