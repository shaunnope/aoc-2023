package main

import (
	"bufio"
	"math"
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

func parse(lines []string) ([]int, []int) {
	digits := regexp.MustCompile(`((\d+)|;)`)
	timesStr := digits.FindAllString(lines[0], -1)
	times := make([]int, len(timesStr))
	for i, t := range timesStr {
		times[i], _ = strconv.Atoi(t)
	}
	distanceStr := digits.FindAllString(lines[1], -1)
	distances := make([]int, len(distanceStr))
	for i, d := range distanceStr {
		distances[i], _ = strconv.Atoi(d)
	}
	return times, distances
}

func part1(times []int, distances []int) {
	res := 1
	for i, t := range times {
		d := distances[i]

		D := (math.Sqrt(float64(t*t - 4*d))) / 2
		low := math.Ceil(float64(t)/2 - D)
		high := math.Floor(float64(t)/2 + D)
		if low < 0 {
			low = 0
		}
		if high > float64(t) {
			high = float64(t)
		}
		res *= int(high - low + 1)

	}

	println(res)
}

func part2(lines []string) {
	res := 0
	digits := regexp.MustCompile(`((\d+)|;)`)
	timesStr := digits.FindAllString(lines[0], -1)
	totalTime := ""
	for _, t := range timesStr {
		totalTime += t
	}
	time, _ := strconv.Atoi(totalTime)
	distanceStr := digits.FindAllString(lines[1], -1)
	totalDistance := ""
	for _, d := range distanceStr {
		totalDistance += d
	}
	distance, _ := strconv.Atoi(totalDistance)
	D := (math.Sqrt(float64(time*time - 4*distance))) / 2
	low := math.Ceil(float64(time)/2 - D)
	high := math.Floor(float64(time)/2 + D)
	if low < 0 {
		low = 0
	}
	if high > float64(time) {
		high = float64(time)
	}
	res = int(high - low + 1)

	println(res)
}

func main() {
	lines := load("input.txt")
	times, distances := parse(lines)
	part1(times, distances)
	part2(lines)

}
