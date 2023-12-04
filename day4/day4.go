package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
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

type Draw struct {
	Winning map[int]bool
	Nums    map[int]bool
}

func parseLine(i int, line string) Draw {
	res := Draw{Winning: make(map[int]bool), Nums: make(map[int]bool)}
	parts := regexp.MustCompile(`\d+|\|`)

	winDone := false
	for _, j := range parts.FindAllString(line, -1)[1:] {
		if !winDone {
			num, ok := strconv.Atoi(j)
			if ok != nil {
				winDone = true
				continue
			}
			res.Winning[num] = true
		} else {
			num, ok := strconv.Atoi(j)
			if ok != nil {
				panic("bad num")
			}
			res.Nums[num] = true
		}
	}

	return res
}

func parse(lines []string) []Draw {
	res := make([]Draw, len(lines))
	for i, l := range lines {
		res[i] = parseLine(i+1, l)
	}
	return res
}

func part1(draws []Draw) {
	res := 0
	for _, d := range draws {
		total := 0
		for k := range d.Winning {
			if _, ok := d.Nums[k]; ok {
				total++
			}
		}
		if total > 0 {
			res += 1 << (total - 1)
		}
	}

	println(res)
}

func part2(draws []Draw) {
	res := 0
	set := make(map[int]int)
	for i, d := range draws {
		set[i]++
		total := 0
		for k := range d.Winning {
			if _, ok := d.Nums[k]; ok {
				total++
			}
		}
		for j := i + 1; j <= i+total; j++ {
			set[j] += set[i]
		}

	}

	for _, v := range set {
		res += v
	}

	println(res)
}

func main() {
	lines := load("input.txt")
	draws := parse(lines)
	part1(draws)
	part2(draws)

}
