package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func task1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var finalArray [][]int

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "")
		tmpArr := []int{}
		for i := 0; i < len(splitLine); i++ {
			val, _ := strconv.Atoi(splitLine[i])
			tmpArr = append(tmpArr, val)
		}
		finalArray = append(finalArray, tmpArr)
	}
	fmt.Println(finalArray)

	var b [][]bool
	for i := 0; i < len(finalArray); i++ {
		a := make([]bool, len(finalArray))
		if i == 0 || i == (len(finalArray)-1) {
			for j := 0; j < len(finalArray); j++ {
				a[j] = true
			}
		} else {
			a[0] = true
			a[len(finalArray)-1] = true
		}
		b = append(b, a)
	}

	for i := 0; i < len(finalArray); i++ {
		currentMaxH := 0
		currentMaxV := 0
		for j := 0; j < len(finalArray); j++ {
			if finalArray[i][j] > currentMaxH {
				currentMaxH = finalArray[i][j]
				b[i][j] = true
			}
			if finalArray[j][i] > currentMaxV {
				currentMaxV = finalArray[j][i]
				b[j][i] = true
			}
		}
		currentMaxV = 0
		currentMaxH = 0
		for j := len(finalArray) - 1; j >= 0; j-- {
			if finalArray[i][j] > currentMaxH {
				currentMaxH = finalArray[i][j]
				b[i][j] = true
			}
			if finalArray[j][i] > currentMaxV {
				currentMaxV = finalArray[j][i]
				b[j][i] = true
			}
			// fmt.Println(j, i, currentMaxV, finalArray[j][i])
		}
	}
	cnt := 0
	for j := 0; j < len(finalArray); j++ {
		for i := 0; i < len(finalArray); i++ {
			if b[j][i] {
				cnt += 1
			}
		}
	}

	fmt.Println(b)
	fmt.Printf("Final count %d\n", cnt)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func calcScoreRight(x, y int, finalArray [][]int) int {
	currentMax := -1
	finalScore := 1

	tmpScore := 0
	for j := x + 1; j < len(finalArray); j++ {
		if finalArray[y][j] >= currentMax {
			currentMax = finalArray[y][j]
			tmpScore += 1
		}
		/* 		if currentMax > finalArray[y][x] {
			break
		} */
	}
	if tmpScore != 0 {
		finalScore *= tmpScore
	}
	fmt.Println("right score", finalScore)
	return finalScore
}

func calcScoreLeft(x, y int, finalArray [][]int) int {
	currentMax := -1
	finalScore := 1

	tmpScore := 0
	for j := x - 1; j >= 0; j-- {
		if finalArray[y][j] >= currentMax {
			currentMax = finalArray[y][j]
			tmpScore += 1
		}
		/* 		if currentMax > finalArray[y][x] {
			break
		} */
	}
	if tmpScore != 0 {
		finalScore *= tmpScore
	}
	fmt.Println("left score", finalScore)
	return finalScore
}

func calcScoreUp(x, y int, finalArray [][]int) int {
	currentMax := -1
	finalScore := 1

	tmpScore := 0
	for j := y - 1; j >= 0; j-- {
		if finalArray[j][x] >= currentMax {
			currentMax = finalArray[j][x]
			tmpScore += 1
		}
		/* 		if currentMax > finalArray[y][x] {
			break
		} */
	}
	if tmpScore != 0 {
		finalScore *= tmpScore
	}

	fmt.Println("up score", finalScore)
	return finalScore
}

func calcScoreDown(x, y int, finalArray [][]int) int {
	currentMax := -1
	finalScore := 1

	tmpScore := 0
	for j := y + 1; j < len(finalArray); j++ {
		if finalArray[j][x] >= currentMax {
			currentMax = finalArray[j][x]
			tmpScore += 1
		}
		/* 		if currentMax > finalArray[y][x] {
			break
		} */
	}
	if tmpScore != 0 {
		finalScore *= tmpScore
	}

	fmt.Println("down score", finalScore)
	return finalScore
}

func calcScore(x, y int, finalArray [][]int) int {
	result := 1
	result *= calcScoreRight(x, y, finalArray)
	result *= calcScoreLeft(x, y, finalArray)
	result *= calcScoreDown(x, y, finalArray)
	result *= calcScoreUp(x, y, finalArray)
	return result
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var finalArray [][]int

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "")
		tmpArr := []int{}
		for i := 0; i < len(splitLine); i++ {
			val, _ := strconv.Atoi(splitLine[i])
			tmpArr = append(tmpArr, val)
		}
		finalArray = append(finalArray, tmpArr)
	}
	fmt.Println(finalArray)
	cnt := 0
	for j := 0; j < len(finalArray); j++ {
		for i := 0; i < len(finalArray); i++ {
			if x := calcScore(i, j, finalArray); x > cnt {
				cnt = x
			}
		}
	}

	fmt.Printf("Final count %d\n", cnt)
	fmt.Println("score,", calcScore(2, 3, finalArray))
	fmt.Println("score,", calcScore(2, 1, finalArray))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
