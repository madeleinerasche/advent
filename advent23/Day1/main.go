package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/**
--- Day 1: Trebuchet?! ---
Something is wrong with global snow production, and you've been selected to take a look. The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.

You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky") and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course, where do you think snow comes from") when you realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?
**/

func ReadFile(filepath string) ([]string, error) {
	content, err := os.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	return lines, nil
}

func calibrationPart1(lines []string) (int, error) {
	sum := 0
	var err error
	var start string
	var end string

	for i := 0; i < len(lines); i++ {
		// find end of string
		k := len(lines[i]) - 1

		// reset each time
		startFound := false
		endFound := false
		num := 0

		for j := 0; j < len(lines[i]); j++ {

			// looking for start
			if unicode.IsNumber(rune(lines[i][j])) && !startFound {
				start = string(lines[i][j])
				startFound = true
			}

			// looking for end
			if unicode.IsNumber(rune(lines[i][k])) && !endFound {
				end = string(lines[i][k])
				endFound = true
			}

			// once both found, add to sum and break from loop
			if startFound && endFound {
				fullNum := start + end
				num, err = strconv.Atoi(fullNum) // convert string to int
				if err == nil {
					sum += num
				}
				break
			}

			// decrement from the end location
			k--
		}
	}

	return sum, err
}

func main() {

	fmt.Println("Day 1: Trebuchet?!")

	lines, err := ReadFile("input.txt")
	if err != nil {
		fmt.Println("Failed to get file ", err)
	}

	num, err := calibrationPart1(lines)
	if err != nil {
		fmt.Println("Failed to get sum ", err)
	}

	fmt.Println(num)
}
