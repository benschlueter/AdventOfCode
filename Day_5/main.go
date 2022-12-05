package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func part1() {
}

/*
	stacks

.		[F] [Q]         [Q]
[B]     [Q] [V] [D]     [S]
[S] [P] [T] [R] [M]     [D]
[J] [V] [W] [M] [F]     [J]     [J]
[Z] [G] [S] [W] [N] [D] [R]     [T]
[V] [M] [B] [G] [S] [C] [T] [V] [S]
[D] [S] [L] [J] [L] [G] [G] [F] [R]
[G] [Z] [C] [H] [C] [R] [H] [P] [D]
1   2   3   4   5   6   7   8   9
*/

func fillStacks(stackMap map[string]*stack.Stack) {
	stackMap["1"].Push("G")
	stackMap["1"].Push("D")
	stackMap["1"].Push("V")
	stackMap["1"].Push("Z")
	stackMap["1"].Push("J")
	stackMap["1"].Push("S")
	stackMap["1"].Push("B")

	stackMap["2"].Push("Z")
	stackMap["2"].Push("S")
	stackMap["2"].Push("M")
	stackMap["2"].Push("G")
	stackMap["2"].Push("V")
	stackMap["2"].Push("P")

	stackMap["3"].Push("C")
	stackMap["3"].Push("L")
	stackMap["3"].Push("B")
	stackMap["3"].Push("S")
	stackMap["3"].Push("W")
	stackMap["3"].Push("T")
	stackMap["3"].Push("Q")
	stackMap["3"].Push("F")

	stackMap["4"].Push("H")
	stackMap["4"].Push("J")
	stackMap["4"].Push("G")
	stackMap["4"].Push("W")
	stackMap["4"].Push("M")
	stackMap["4"].Push("R")
	stackMap["4"].Push("V")
	stackMap["4"].Push("Q")

	stackMap["5"].Push("C")
	stackMap["5"].Push("L")
	stackMap["5"].Push("S")
	stackMap["5"].Push("N")
	stackMap["5"].Push("F")
	stackMap["5"].Push("M")
	stackMap["5"].Push("D")

	stackMap["6"].Push("R")
	stackMap["6"].Push("G")
	stackMap["6"].Push("C")
	stackMap["6"].Push("D")

	stackMap["7"].Push("H")
	stackMap["7"].Push("G")
	stackMap["7"].Push("T")
	stackMap["7"].Push("R")
	stackMap["7"].Push("J")
	stackMap["7"].Push("D")
	stackMap["7"].Push("S")
	stackMap["7"].Push("Q")

	stackMap["8"].Push("P")
	stackMap["8"].Push("F")
	stackMap["8"].Push("V")

	stackMap["9"].Push("D")
	stackMap["9"].Push("R")
	stackMap["9"].Push("S")
	stackMap["9"].Push("T")
	stackMap["9"].Push("J")
}

func task1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	x := make(map[string]*stack.Stack)

	scanner := bufio.NewScanner(f)

	for i := 1; i < 10; i++ {
		x[fmt.Sprint(i)] = stack.New()
	}
	fillStacks(x)

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		fmt.Println(splitLine)
		numToTransfer, _ := strconv.Atoi(splitLine[1])
		for i := 0; i < numToTransfer; i++ {
			x[splitLine[5]].Push(x[splitLine[3]].Pop())
		}

	}
	for i := 1; i < 10; i++ {
		fmt.Printf("%s", x[fmt.Sprint(i)].Peek())
	}

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

	x := make(map[string]*stack.Stack)

	scanner := bufio.NewScanner(f)

	for i := 1; i < 10; i++ {
		x[fmt.Sprint(i)] = stack.New()
	}
	fillStacks(x)

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		fmt.Println(splitLine)
		numToTransfer, _ := strconv.Atoi(splitLine[1])
		var tmp []interface{}
		for i := 0; i < numToTransfer; i++ {
			tmp = append(tmp, x[splitLine[3]].Pop())
		}
		for i := len(tmp) - 1; i >= 0; i-- {
			x[splitLine[5]].Push(tmp[i])
		}

	}
	for i := 1; i < 10; i++ {
		fmt.Printf("%s", x[fmt.Sprint(i)].Peek())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
