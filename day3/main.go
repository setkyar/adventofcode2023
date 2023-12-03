package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Pair struct {
	First, Second int
}

func main() {
	fptr := flag.String("fpath", "input.txt", "file path to read data from")
	flag.Parse()

	file, err := os.Open(*fptr)

	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	n := len(grid)
	m := len(grid[0])

	total := 0
	table := make(map[Pair][]int)
	for i := 0; i < n; i++ {
		j := 0
		for j < m {
			if !unicode.IsDigit(rune(grid[i][j])) {
				j++
				continue
			}
			j2 := j + 1
			for j2 < m && unicode.IsDigit(rune(grid[i][j2])) {
				j2++
			}
			nbrs := getNeighbors(i, j, j2, n, m, grid)
			for nbr := range nbrs {
				nbri, nbrj := nbr.First, nbr.Second
				if isSymbol(rune(grid[nbri][nbrj])) {
					num, _ := strconv.Atoi(grid[i][j:j2])
					total += num
					// part 2
					if grid[nbri][nbrj] == '*' {
						table[Pair{nbri, nbrj}] = append(table[Pair{nbri, nbrj}], num)
					}
					break
				}
			}
			j = j2
		}
	}

	// part 1
	fmt.Println(total)

	// part 2
	gearRatioSum := 0
	for _, v := range table {
		if len(v) == 2 {
			gearRatioSum += v[0] * v[1]
		}
	}
	fmt.Println(gearRatioSum)
}

func getNeighbors(i, j1, j2, n, m int, grid []string) map[Pair]struct{} {
	nbrSet := make(map[Pair]struct{})
	for j := j1; j < j2; j++ {
		for di := -1; di <= 1; di++ {
			for dj := -1; dj <= 1; dj++ {
				if di == 0 && dj == 0 {
					continue
				}
				nbri, nbrj := i+di, j+dj
				if nbri >= 0 && nbri < n && nbrj >= 0 && nbrj < m {
					nbrSet[Pair{nbri, nbrj}] = struct{}{}
				}
			}
		}
	}
	return nbrSet
}

func isSymbol(char rune) bool {
	return !unicode.IsDigit(char) && char != '.'
}
