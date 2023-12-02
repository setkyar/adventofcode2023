package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fptr := flag.String("fpath", "input.txt", "file path to read data from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	total := 0
	for s.Scan() {
		line := s.Text()
		total += gameResult(line)
	}

	fmt.Println("Game Result ", total)
}

func gameResult(line string) int {
	game := line[:strings.IndexByte(line, ':')]
	currentGame := strings.Split(game, "Game ")
	currentGameInt, _ := strconv.Atoi(currentGame[1])

	results := strings.Replace(line, game+":", "", 1)

	// get sets
	sets := strings.Split(results, ";")

	return checkForMinSetOfCubes(results)

	// This is for part I
	for _, set := range sets {
		currentSet := strings.Split(set, ",")
		if !checkForGameRule(currentSet) {
			return 0
		}
	}

	return currentGameInt
}

func checkForMinSetOfCubes(s string) int {
	result := strings.ReplaceAll(s, ";", ",")
	sli := strings.Split(result, ",")

	cubes := make(map[string]int)

	for i := 0; i < len(sli); i++ {
		keyPair := strings.Split(strings.TrimSpace(sli[i]), " ")
		val, _ := strconv.Atoi(keyPair[0])
		if cubes[keyPair[1]] < val {
			cubes[keyPair[1]] = val
		}
	}

	total := 0
	for _, v := range cubes {
		if total == 0 {
			total = v
		} else {
			total *= v
		}
	}

	return total
}

func checkForGameRule(s []string) bool {
	gameRule := map[string]int{"red": 12, "green": 13, "blue": 14}

	for _, val := range s {
		result := strings.Split(val, " ")
		i, _ := strconv.Atoi(result[1])

		if gameRule[result[2]] < i {
			return false
		}
	}

	return true
}
