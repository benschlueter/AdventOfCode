package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	accumulate := 0
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ",")

		result1 := strings.Split(splitLine[0], "-")
		result2 := strings.Split(splitLine[1], "-")

		a, _ := strconv.Atoi(result1[0])
		b, _ := strconv.Atoi(result1[1])
		c, _ := strconv.Atoi(result2[0])
		d, _ := strconv.Atoi(result2[1])

		if a <= c && b >= d {
			accumulate += 1
			fmt.Println(a, b, c, d)
			fmt.Println(splitLine, result1, result2, "match")
			continue
		}
		if a >= c && b <= d {
			accumulate += 1
			fmt.Println(a, c)
			fmt.Println(b, d)
			fmt.Println(splitLine, result1, result2, "match")
			continue
		}
		fmt.Println(splitLine, result1, result2)

	}
	fmt.Printf("Total Score %d", accumulate)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	accumulate := 0
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ",")

		result1 := strings.Split(splitLine[0], "-")
		result2 := strings.Split(splitLine[1], "-")

		a, _ := strconv.Atoi(result1[0])
		b, _ := strconv.Atoi(result1[1])
		c, _ := strconv.Atoi(result2[0])
		d, _ := strconv.Atoi(result2[1])

		//  a .... b
		//      c .....d
		if (b >= c && b <= d) || (d >= a && d <= b) {
			accumulate += 1
			fmt.Println(splitLine, result1, result2, "match")
			continue
		}
		// 			 a .... b
		//      c .....d
		fmt.Println(splitLine, result1, result2)

	}
	fmt.Printf("Total Score %d", accumulate)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
