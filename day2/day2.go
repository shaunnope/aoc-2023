package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
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

type Color string

const (
	RED   Color = "red"
	BLUE  Color = "blue"
	GREEN Color = "green"
)

func NewColor(s string) Color {
	switch s {
	case "r":
		return RED
	case "g":
		return GREEN
	case "b":
		return BLUE
	default:
		panic("invalid color")
	}
}

type Draw struct {
	Num   int
	Color Color
}

func parseDraws(line []string) [][]Draw {
	res := make([][]Draw, len(line))
	res[0] = make([]Draw, 0)
	i := 0
	for _, l := range line {
		if l == ";" {
			i++
			res[i] = make([]Draw, 0)
			continue
		}
		num, err := strconv.Atoi(l[:len(l)-2])
		if err != nil {
			panic(err)
		}
		draw := Draw{Num: num, Color: NewColor(l[len(l)-1:])}
		res[i] = append(res[i], draw)
	}
	return res
}

func parseLine(line string) [][]Draw {
	digits := regexp.MustCompile(`((\d+ (r|g|b))|;)`)
	return parseDraws(digits.FindAllString(line, -1))
}

func part1(lines []string) {
	limits := map[Color]int{
		RED:   12,
		GREEN: 13,
		BLUE:  14,
	}
	res := 0
Loop:
	for i, line := range lines {
		draws := parseLine(line)

		for _, draw := range draws {
			for _, d := range draw {
				if d.Num > limits[d.Color] {
					// fmt.Printf("line %d: %v\n", i, draw)
					continue Loop
				}
			}
		}
		res += i + 1

	}
	println(res)
}

func part2(lines []string) {
	res := 0
	for _, line := range lines {
		mins := map[Color]int{
			RED:   0,
			GREEN: 0,
			BLUE:  0,
		}
		draws := parseLine(line)

		for _, draw := range draws {
			for _, d := range draw {
				mins[d.Color] = max(mins[d.Color], d.Num)
			}
		}
		res += mins[RED] * mins[GREEN] * mins[BLUE]

	}
	println(res)
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

func main() {
	lines := load("input.txt")
	part1(lines)
	part2(lines)

}
