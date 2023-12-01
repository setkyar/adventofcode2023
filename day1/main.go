package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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

	total := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		t := convertStringToNumber(line)
		re := regexp.MustCompile("[0-9]+")
		stringNums := re.FindAllString(t, -1)
		total += sum(stringNums)
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

func convertStringToNumber(s string) string {
	n := map[string]string{
		"one":   "o1ne",
		"two":   "t2wo",
		"three": "t3hree",
		"four":  "f4our",
		"five":  "f5ive",
		"six":   "s6x",
		"seven": "s7even",
		"eight": "e8ight",
		"nine":  "n9ine",
	}

	for i, v := range n {
		s = strings.ReplaceAll(s, i, v)
	}

	return s
}

func sum(slice []string) int {
	str := ""
	for _, v := range slice {
		str += v
	}
	first := string([]rune(str)[0])
	last := string([]rune(str)[len(str)-1])

	num := first + last
	i, _ := strconv.Atoi(num)
	return i
}
