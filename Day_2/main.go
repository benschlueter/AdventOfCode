package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		switch line[2] {
		case 'X': // Rock
			accumulate += 1
		case 'Y': // Paper
			accumulate += 2
		case 'Z': // Scissor
			accumulate += 3
		}
		// draw round
		if line[0] == line[2]-23 {
			accumulate += 3
			continue
		}
		if (line[0] == 'A' && line[2] == 'Y') || (line[0] == 'B' && line[2] == 'Z') || (line[0] == 'C' && line[2] == 'X') {
			accumulate += 6
			continue
		}
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
		switch line[2] {
		case 'X': // Lose
			accumulate += 0
			switch line[0] {
			case 'A': // Scissor loses against Rock
				accumulate += 3
			case 'B': // Rock loses against Paper
				accumulate += 1
			case 'C': // Paper loses against Scissor
				accumulate += 2
			}
		case 'Y': // Tie
			accumulate += 3
			switch line[0] {
			case 'A': // Rock tie against Rock
				accumulate += 1
			case 'B': // Paper tie against Paper
				accumulate += 2
			case 'C': // Scissor tie against Scissor
				accumulate += 3
			}
		case 'Z': // Win
			accumulate += 6
			switch line[0] {
			case 'A': // Paper wins against Rock
				accumulate += 2
			case 'B': // Scissor wins against Paper
				accumulate += 3
			case 'C': // Rock wins against Scissor
				accumulate += 1
			}
		}
	}
	fmt.Printf("Total Score %d", accumulate)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
