package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
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

type Index struct {
	From string
	To   string
}

type MarkerType int64

const (
	START MarkerType = iota
	END
)

type Marker struct {
	From int64
	To   int64
	Type MarkerType
	Next string
}

type Data struct {
	Seeds []int64
	Maps  map[Index][]Marker
}

type Draw struct {
	Winning map[int64]bool
	Nums    map[int64]bool
}

func parse(lines []string) (Data, []Index) {
	// get seeds
	digits := regexp.MustCompile(`\d+`)
	seeds_raw := digits.FindAllString(lines[0], -1)
	seeds := make([]int64, len(seeds_raw))
	for i, s := range seeds_raw {
		n, _ := strconv.Atoi(s)
		seeds[i] = int64(n)
	}

	mapper := regexp.MustCompile(`([^\s]+)-to-([^\s]+)`)
	idx := Index{}
	idxs := make([]Index, 0)
	maps := make(map[Index][]Marker)
	maps[idx] = make([]Marker, 0)
	for _, l := range lines[2:] {
		matches := mapper.FindStringSubmatch(l)
		if len(matches) != 0 {
			mapping := maps[idx]
			sort.Slice(mapping, func(i, j int) bool {
				return mapping[i].From < mapping[j].From
			})
			maps[idx] = mapping
			idx = Index{From: matches[1], To: matches[2]}
			idxs = append(idxs, idx)
			maps[idx] = make([]Marker, 0)
			continue
		}
		digits := regexp.MustCompile(`\d+`)
		nums := digits.FindAllString(l, -1)
		if len(nums) != 3 {
			continue
		}
		target, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		source, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		d, err := strconv.Atoi(nums[2])
		if err != nil {
			panic(err)
		}

		maps[idx] = append(maps[idx], []Marker{
			{From: int64(source), To: int64(target), Type: START},
			{From: int64(source + d - 1), To: int64(target + d - 1), Type: END},
		}...)
	}

	res := Data{Seeds: seeds, Maps: maps}
	return res, idxs
}

func part1(data Data, idxs []Index) {
	res := int64(math.MaxInt64)
	for _, d := range data.Seeds {
		val := d
		parts := make([]int64, 0)
		for _, ix := range idxs {
			parts = append(parts, val)
			markers := data.Maps[ix]
			mix := sort.Search(len(markers), func(i int) bool {
				return markers[i].From >= val
			})
			if mix == len(markers) {
				// fmt.Print("o")
				continue
			} else if markers[mix].From == val {
				// fmt.Print("s")
				val = markers[mix].To
				continue
			} else if markers[mix].Type == END {
				// fmt.Print("e")
				val = markers[mix].To - (markers[mix].From - val)
				continue
			}
			// fmt.Print("m")
		}
		// fmt.Println()
		parts = append(parts, val)
		if val < res {
			res = val
		}
		fmt.Printf("%v\n", parts)
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
	// lines := load("example.txt")
	data, idxs := parse(lines)
	for _, marker := range data.Maps[idxs[1]] {
		fmt.Printf("%v\n", marker)
	}
	// part1(data, idxs)
	// part2(draws)

}
