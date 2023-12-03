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

type EntType string

const (
	NUM EntType = "num"
	SYM EntType = "sym"
)

type TableRow map[int]Ent
type Table []TableRow

type Index struct {
	Row int
	Col int
}

type Ent struct {
	Type   EntType
	Val    string
	Offset int
}

func parseLine(i int, line string) TableRow {
	row := make(TableRow)
	digits := regexp.MustCompile(`\d+`)
	symbols := regexp.MustCompile(`[^\d|.]`)

	dig_indices := digits.FindAllStringIndex(line, -1)
	sym_indices := symbols.FindAllStringIndex(line, -1)

	for _, j := range dig_indices {
		for k := j[0]; k < j[1]; k++ {
			row[k] = Ent{Type: NUM, Val: line[j[0]:j[1]], Offset: j[0]}
		}
	}

	for _, j := range sym_indices {
		// if j[0] != j[1]-1 {
		// 	panic(fmt.Sprintf("invalid symbol: %v %v", j, line[j[0]:j[1]]))
		// }
		row[j[0]] = Ent{Type: SYM, Val: line[j[0]:j[1]], Offset: j[0]}
	}
	return row
}

func parse(lines []string) Table {
	res := make([]TableRow, len(lines)+2)
	for i, l := range lines {
		res[i+1] = parseLine(i, l)
	}
	return res
}

func part1(table Table) {
	res := 0
	set := make(map[Index]Ent)
	for i, row := range table {
		for _, ent := range row {
			if ent.Type == SYM {
				// check current row
				left, ok := row[ent.Offset-1]
				if ok && left.Type == NUM {
					set[Index{i, left.Offset}] = left
				}
				right, ok := row[ent.Offset+1]
				if ok && right.Type == NUM {
					set[Index{i, right.Offset}] = right
				}
				// check previous row
				prev_row := table[i-1]
				for k := ent.Offset - 1; k <= ent.Offset+1; k++ {
					prev, ok := prev_row[k]
					if ok && prev.Type == NUM {
						set[Index{i - 1, prev.Offset}] = prev
					}
				}
				// check next row
				next_row := table[i+1]
				for k := ent.Offset - 1; k <= ent.Offset+1; k++ {
					next, ok := next_row[k]
					if ok && next.Type == NUM {
						set[Index{i + 1, next.Offset}] = next
					}
				}
			}
		}
	}

	for _, ent := range set {
		num, err := strconv.Atoi(ent.Val)
		if err != nil {
			panic(err)
		}
		res += num
	}
	println(res)
}

func part2(table Table) {
	res := 0
	for i, row := range table {
		for _, ent := range row {
			if ent.Type == SYM && ent.Val == "*" {
				set := make(map[Index]Ent)
				// check current row
				left, ok := row[ent.Offset-1]
				if ok && left.Type == NUM {
					set[Index{i, left.Offset}] = left
				}
				right, ok := row[ent.Offset+1]
				if ok && right.Type == NUM {
					set[Index{i, right.Offset}] = right
				}
				// check previous row
				prev_row := table[i-1]
				for k := ent.Offset - 1; k <= ent.Offset+1; k++ {
					prev, ok := prev_row[k]
					if ok && prev.Type == NUM {
						set[Index{i - 1, prev.Offset}] = prev
					}
				}
				// check next row
				next_row := table[i+1]
				for k := ent.Offset - 1; k <= ent.Offset+1; k++ {
					next, ok := next_row[k]
					if ok && next.Type == NUM {
						set[Index{i + 1, next.Offset}] = next
					}
				}
				if len(set) != 2 {
					continue
				}
				prod := 1
				for _, ent := range set {
					num, err := strconv.Atoi(ent.Val)
					if err != nil {
						panic(err)
					}
					prod *= num
				}
				res += prod
			}
		}
	}

	println(res)
}

func main() {
	lines := load("input.txt")
	table := parse(lines)
	part1(table)
	part2(table)

}
