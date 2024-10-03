package main

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	file_content, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(file_content), "\n")

	var sum int64

	for i := 0; i < len(lines)-1; i++ {
		seq := lines[i]
		var single_digits []string

		for j := 0; j < len(seq); j++ {
			char := seq[j]
			if unicode.IsNumber(rune(char)) {
				single_digits = append(single_digits, string(seq[j]))
			}
		}
		digit_sum, _ := strconv.ParseInt(single_digits[0]+single_digits[len(single_digits)-1], 10, 8)
		sum += digit_sum
	}
	println("Result: ", sum)
}
