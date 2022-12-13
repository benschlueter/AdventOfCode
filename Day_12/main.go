package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type node struct {
	xPos            int
	yPos            int
	value           int
	visited         bool
	stepsToThisNode int
}

var (
	resolve = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"S": 1,
		"E": 26,
	}
	directions = [4][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
)

func traverse(arr [][]*node, current, end [2]int, step int) {
	xlength := len(arr[0])
	ylength := len(arr)
	currentNode := arr[current[1]][current[0]]
	if currentNode.stepsToThisNode > step {
		currentNode.stepsToThisNode = step
	} else {
		return
	}
	arr[current[1]][current[0]].visited = true

	for _, v := range directions {
		if current[0]+v[0] < xlength &&
			current[1]+v[1] < ylength &&
			current[0]+v[0] >= 0 &&
			current[1]+v[1] >= 0 &&
			!arr[current[1]+v[1]][current[0]+v[0]].visited &&
			arr[current[1]+v[1]][current[0]+v[0]].value-currentNode.value <= 1 {

			tmp := [2]int{current[0] + v[0], current[1] + v[1]}
			traverse(arr, tmp, end, step+1)
		}
	}
	arr[current[1]][current[0]].visited = false
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var array [][]*node
	var start, end [2]int
	i := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splitline := strings.Split(line, "")
		var tmp []*node
		for j, v := range splitline {
			if v == "S" {
				start[0] = j
				start[1] = i
			}
			if v == "E" {
				end[0] = j
				end[1] = i
			}
			tmp = append(tmp, &node{
				xPos:            j,
				yPos:            i,
				visited:         false,
				value:           resolve[v],
				stepsToThisNode: 100000000000,
			})
		}
		array = append(array, tmp)
		i += 1
	}
	traverse(array, start, end, 0)
	fmt.Println(array[end[1]][end[0]].stepsToThisNode)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
