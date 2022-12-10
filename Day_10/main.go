package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func update(cycle, regX, strength *int) {
	if ((*cycle + 20) % 40) == 0 {
		*strength += *cycle * *regX
		fmt.Println(*cycle, *regX)
	}
}

func update2(cycle, regX int) {
	cycleMod := cycle % 40
	if ((cycleMod)%40) == 0 && cycle != 0 {
		fmt.Printf("\n")
	}
	if (cycleMod == regX) || (cycleMod == (regX - 1)) || (cycleMod == (regX + 1)) {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	cycle := 0
	strength := 0
	regX := 1
	for scanner.Scan() {
		line := scanner.Text()
		splitline := strings.Split(line, " ")
		if splitline[0] == "noop" {
			update(&cycle, &regX, &strength)
			cycle += 1
		} else {
			update(&cycle, &regX, &strength)
			cycle += 1
			update(&cycle, &regX, &strength)
			cycle += 1
			val, _ := strconv.Atoi(splitline[1])
			regX += val
		}

	}
	fmt.Println(strength)

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
	cycle := 0
	regX := 1
	for scanner.Scan() {
		line := scanner.Text()
		splitline := strings.Split(line, " ")
		if splitline[0] == "noop" {
			update2(cycle, regX)
			cycle += 1
		} else {
			update2(cycle, regX)
			val, _ := strconv.Atoi(splitline[1])
			cycle += 1
			update2(cycle, regX)
			cycle += 1
			regX += val
		}
	}
	update2(cycle, regX)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
