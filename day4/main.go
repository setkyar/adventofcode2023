package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const DefaultFilePath = "input.txt"

type Card struct {
	givenNumbers   []string
	winningNumbers []string
}

func (c *Card) checkPrizeForPartOne() int {
	winningNumberMap := make(map[string]struct{})
	for _, num := range c.winningNumbers {
		winningNumberMap[num] = struct{}{}
	}

	winner := 0
	for _, num := range c.givenNumbers {
		if _, exists := winningNumberMap[num]; exists {
			if winner == 0 {
				winner = 1
			} else {
				winner *= 2
			}
		}
	}
	return winner
}

func (c *Card) countMatches() int {
	matches := 0
	winningNumberMap := make(map[string]struct{})
	for _, num := range c.winningNumbers {
		winningNumberMap[num] = struct{}{}
	}

	for _, givenNum := range c.givenNumbers {
		if _, exists := winningNumberMap[givenNum]; exists {
			matches++
		}
	}
	return matches
}

func main() {
	fptr := flag.String("fpath", DefaultFilePath, "file path to read data from")
	flag.Parse()

	file, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	var cards []Card
	for s.Scan() {
		line := s.Text()
		parts := strings.Split(line, " | ")
		if len(parts) != 2 {
			log.Fatal("Invalid line format:", line)
		}
		cards = append(cards, Card{
			givenNumbers:   strings.Fields(parts[1]),
			winningNumbers: strings.Fields(parts[0]),
		})
	}

	totalPartOne := 0
	for _, card := range cards {
		totalPartOne += card.checkPrizeForPartOne()
	}

	fmt.Println("Part one result:", totalPartOne)
	fmt.Println("Part two result:", checkPrizeForPartTwo(cards))
}

func checkPrizeForPartTwo(cards []Card) int {
	cardCounts := make(map[int]int)
	for i := range cards {
		cardCounts[i] = 1 // Card index starts from 0
	}

	for i, card := range cards {
		matches := card.countMatches()
		for j := 1; j <= matches; j++ {
			if i+j < len(cards) {
				cardCounts[i+j] += cardCounts[i]
			}
		}
	}

	totalCards := 0
	for _, count := range cardCounts {
		totalCards += count
	}

	return totalCards
}
