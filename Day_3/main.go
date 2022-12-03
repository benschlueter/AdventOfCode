package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var valueLookup map[byte]int

func fillTable() {
	valueLookup = make(map[byte]int)
	x := byte('a')
	for i := 1; i <= 26; i++ {
		valueLookup[x] = i
		x += 1
	}
	x = byte('A')
	for i := 27; i <= 52; i++ {
		valueLookup[x] = i
		x += 1
	}
}

func task1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	fillTable()

	accumulate := 0
	var x map[byte]struct{}

	for scanner.Scan() {
		line := scanner.Text()
		length := len(line) >> 1
		x = make(map[byte]struct{})
		for i := 0; i < length; i++ {
			x[line[i]] = struct{}{}
		}
		for i := length; i < 2*length; i++ {
			if _, ok := x[line[i]]; ok {
				fmt.Printf("Found match %c | %d\n", line[i], i)
				accumulate += valueLookup[line[i]]
				break
			}
		}
		fmt.Println(line)
	}
	fmt.Println(accumulate)

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
	fillTable()

	accumulate := 0
	var x map[byte]struct{}
	var y map[byte]struct{}

	for scanner.Scan() {
		x = make(map[byte]struct{})
		y = make(map[byte]struct{})
		line1 := scanner.Text()
		if !scanner.Scan() {
			log.Fatal("no lines left")
		}
		line2 := scanner.Text()
		if !scanner.Scan() {
			log.Fatal("no lines left")
		}
		line3 := scanner.Text()
		for i := 0; i < len(line1); i++ {
			x[line1[i]] = struct{}{}
		}
		for i := 0; i < len(line2); i++ {
			y[line2[i]] = struct{}{}
		}
		for i := 0; i < len(line3); i++ {
			if _, ok := x[line3[i]]; ok {
				if _, ok := y[line3[i]]; ok {
					accumulate += valueLookup[line3[i]]
					break
				}
			}
		}

	}
	fmt.Println(accumulate)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
