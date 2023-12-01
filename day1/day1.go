package main

import (
	"bufio"
	"os"
	"regexp"
	"unicode/utf8"
)

func load(path string) []string {
	readFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}

func part1(lines []string) {
	res := 0
	for _, line := range lines {
		// line is a string of alphanumeric characters
		// use regexp to get the first and last digit in line
		// convert them to int
		// add them to res[idx]
		digits := regexp.MustCompile(`\d`)
		nums := digits.FindAllString(line, -1)
		res += 10*int(nums[0][0]-'0') + int(nums[len(nums)-1][0]-'0')
	}
	println(res)
}

func parseNum(num string) int {
	wordToDigit := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	val, ok := wordToDigit[num]
	if !ok {
		val = int(num[0] - '0')
	}
	return val
}

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func part2(lines []string) {
	res := 0
	for _, line := range lines {
		digits := regexp.MustCompile(`(\d|zero|one|two|three|four|five|six|seven|eight|nine)`)
		first := parseNum(digits.FindString(line))
		rev_digits := regexp.MustCompile(`(\d|orez|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)`)
		last := parseNum(Reverse(rev_digits.FindString(Reverse(line))))
		res += 10*first + last
	}
	println(res)
}

func main() {
	lines := load("input.txt")
	part1(lines)
	part2(lines)

}
