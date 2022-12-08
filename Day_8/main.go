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
	currentHeight := finalArray[y][x]

	tmpScore := 0
	for j := x + 1; j < len(finalArray); j++ {
		if finalArray[y][j] < currentHeight {
			tmpScore += 1
		} else {
			tmpScore += 1
			break
		}
	}

	return tmpScore
}

func calcScoreLeft(x, y int, finalArray [][]int) int {
	currentHeight := finalArray[y][x]

	tmpScore := 0
	for j := x - 1; j >= 0; j-- {
		if finalArray[y][j] < currentHeight {
			tmpScore += 1
		} else {
			tmpScore += 1
			break
		}
	}

	return tmpScore
}

func calcScoreUp(x, y int, finalArray [][]int) int {
	currentHeight := finalArray[y][x]

	tmpScore := 0
	for j := y - 1; j >= 0; j-- {
		if finalArray[j][x] < currentHeight {
			tmpScore += 1
		} else {
			tmpScore += 1
			break
		}
	}

	return tmpScore
}

func calcScoreDown(x, y int, finalArray [][]int) int {
	currentHeight := finalArray[y][x]

	tmpScore := 0
	for j := y + 1; j < len(finalArray); j++ {
		if finalArray[j][x] < currentHeight {
			tmpScore += 1
		} else {
			tmpScore += 1
			break
		}
	}

	return tmpScore
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
	cnt := 0
	for j := 0; j < len(finalArray); j++ {
		for i := 0; i < len(finalArray); i++ {
			if x := calcScore(i, j, finalArray); x > cnt {
				cnt = x
			}
		}
	}

	fmt.Printf("Final count %d\n", cnt)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
