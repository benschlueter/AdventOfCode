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

func movByOne(newPosH, posT *[2]int, positions map[string]struct{}, tail bool) {
	x, y := newPosH[0]-posT[0], newPosH[1]-posT[1]
	if (abs(posT[0]-newPosH[0]) == 2) && (abs(posT[1]-newPosH[1]) == 1) ||
		(abs(posT[0]-newPosH[0]) == 1) && (abs(posT[1]-newPosH[1]) == 2) {
		posT[0] += (x / abs(x))
		posT[1] += (y / abs(y))
	}
	if (abs(posT[0]-newPosH[0]) == 2) || (abs(posT[1]-newPosH[1]) == 2) {
		posT[0] += (x / 2)
		posT[1] += (y / 2)
	}
	if tail {
		// fmt.Println("update T", newPosH, posT)
		positions[fmt.Sprintf("%d:%d", posT[0], posT[1])] = struct{}{}
	}
}

func move(nodes *[lengthTailPlusOne][2]int, x, y int, positions map[string]struct{}) {
	nodes[0][0] += x
	nodes[0][1] += y
	for i := 1; i < lengthTailPlusOne; i++ {
		movByOne(&nodes[i-1], &nodes[i], positions, i == (lengthTailPlusOne-1))
	}
}

const lengthTailPlusOne = 10

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	positions := make(map[string]struct{})
	nodes := [lengthTailPlusOne][2]int{}
	for _, v := range nodes {
		v[0] = 0
		v[1] = 0
	}

	var x, y int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splitline := strings.Split(line, " ")
		value, _ := strconv.Atoi(splitline[1])
		switch splitline[0] {
		case "R":
			x, y = 1, 0
		case "L":
			x, y = -1, 0
		case "U":
			x, y = 0, 1
		case "D":
			x, y = 0, -1
		}

		// fmt.Println(nodes)

		fmt.Println(x, y, value)
		for i := 0; i < value; i++ {
			// x, y := posH[0]-old[0], posH[1]-old[1]
			move(&nodes, x, y, positions)
		}

	}

	fmt.Println(1 == 1)
	fmt.Println(len(positions))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
