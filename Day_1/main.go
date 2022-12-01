package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	elfs := [3]int{0, 0, 0}
	sum := 0
	for scanner.Scan() {
		if x := scanner.Text(); x == "" {
			if sum > elfs[0] {
				elfs[2] = elfs[1]
				elfs[1] = elfs[0]
				elfs[0] = sum
			} else if sum > elfs[1] {
				elfs[2] = elfs[1]
				elfs[1] = sum
			} else if sum > elfs[2] {
				elfs[2] = sum
			}
			sum = 0
		} else {
			intVar, err := strconv.Atoi(x)
			if err != nil {
				log.Fatal(err)
			}
			sum += intVar
		}
	}
	fmt.Printf("elf sum %d", elfs[0]+elfs[1]+elfs[2])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
