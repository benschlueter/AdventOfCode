package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func movByOne(newPosH, posT *[2]int, positions map[string]struct{}) {
	if (abs(posT[0]-newPosH[0]) == 2) && (abs(posT[1]-newPosH[1]) == 1) ||
		(abs(posT[0]-newPosH[0]) == 1) && (abs(posT[1]-newPosH[1]) == 2) {
		x, y := newPosH[0]-posT[0], newPosH[1]-posT[1]
		if abs(x) == 2 {
			posT[0] += (x / 2)
			posT[1] += y
		}
		if abs(y) == 2 {
			posT[0] += (x)
			posT[1] += (y / 2)
		}
		fmt.Println("update T", posT)
	}
	if (abs(posT[0]-newPosH[0]) == 2) || (abs(posT[1]-newPosH[1]) == 2) {
		x, y := newPosH[0]-posT[0], newPosH[1]-posT[1]
		posT[0] += (x / 2)
		posT[1] += (y / 2)
		fmt.Println("update T", posT)
	}
	positions[fmt.Sprintf("%d:%d", posT[0], posT[1])] = struct{}{}
}

func move(posH, newPosH, posT *[2]int, positions map[string]struct{}) {
	x, y := newPosH[0]-posH[0], newPosH[1]-posH[1]

	for i := 1; i <= x; i++ {
		movByOne(&[2]int{posH[0] + i, posH[1]}, posT, positions)
	}
	for i := -1; i >= x; i-- {
		movByOne(&[2]int{posH[0] + i, posH[1]}, posT, positions)
	}
	for i := 1; i <= y; i++ {
		movByOne(&[2]int{posH[0], posH[1] + i}, posT, positions)
	}
	for i := -1; i >= y; i-- {
		movByOne(&[2]int{posH[0], posH[1] + i}, posT, positions)
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	positions := make(map[string]struct{})
	posH := [2]int{0, 0}
	posT := [2]int{0, 0}
	newPosH := [2]int{0, 0}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splitline := strings.Split(line, " ")
		value, _ := strconv.Atoi(splitline[1])
		switch splitline[0] {
		case "R":
			newPosH[0] += value
		case "L":
			newPosH[0] -= value
		case "U":
			newPosH[1] += value
		case "D":
			newPosH[1] -= value
		}
		move(&posH, &newPosH, &posT, positions)
		fmt.Println(posT, newPosH)
		posH[0] = newPosH[0]
		posH[1] = newPosH[1]
	}

	fmt.Println(len(positions))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
