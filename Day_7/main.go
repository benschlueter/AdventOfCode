package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type dir struct {
	files  []*dir
	size   int
	parent *dir
	isDir  bool
	name   string
}

var (
	totalSum                     = 0
	maxDiskSpace                 = 70000000
	neededDiskSpace              = 30000000
	spaceWeNeedToFree            = 0
	smallestDirSizeToAchieveGoal = 70000000
)

func getSizeOfDir(baseFile *dir) int {
	x := smallestDirSizeToAchieveGoal
	if baseFile.size < smallestDirSizeToAchieveGoal && baseFile.size >= spaceWeNeedToFree {
		x = baseFile.size
	}
	for _, v := range baseFile.files {
		if v.isDir {
			if y := getSizeOfDir(v); y < x {
				x = y
				fmt.Printf("new smallest dir is %s %d\n", v.name, v.size)
			}
		}
	}
	return x
}

func initializeSize(baseFile *dir) int {
	for _, v := range baseFile.files {
		if !v.isDir {
			baseFile.size += v.size
		} else {
			baseFile.size += initializeSize(v)
		}
	}
	if baseFile.size <= 100000 {
		totalSum += baseFile.size
	}
	fmt.Printf("dir: %s size: %d\n", baseFile.name, baseFile.size)
	return baseFile.size
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	base := dir{
		name:  "/",
		isDir: true,
	}
	base.parent = &base

	currentDir := &base
	for scanner.Scan() {
		line := scanner.Text()
		splitline := strings.Split(line, " ")
		if line == "$ cd /" {
			currentDir = &base
			continue
		}
		// fmt.Println(currentDir.name)

		if splitline[0] == "dir" {
			currentDir.files = append(currentDir.files,
				&dir{
					files:  nil,
					size:   0,
					parent: currentDir,
					isDir:  true,
					name:   splitline[1],
				},
			)
			// fmt.Println("added dir", splitline[1], "to", currentDir.name)
			continue
		}

		if match, _ := regexp.MatchString("^([0-9])", line); match {
			size, _ := strconv.Atoi(splitline[0])
			currentDir.files = append(currentDir.files,
				&dir{
					files: nil,
					size:  size,
					isDir: false,
					name:  splitline[1],
				},
			)
			// fmt.Println("added file", splitline[1], "to", currentDir.name)
			continue
		}
		if len(splitline) == 3 && splitline[0] == "$" && splitline[1] == "cd" {
			if splitline[2] == ".." {
				currentDir = currentDir.parent
				continue
			}
			for i := 0; i < len(currentDir.files); i++ {
				if currentDir.files[i].isDir && currentDir.files[i].name == splitline[2] {
					currentDir = currentDir.files[i]
					break
				}
			}
			continue
		}
	}

	totalSum = 0
	rootDirSize := initializeSize(&base)
	fmt.Printf("Root dir size %d\n", rootDirSize)
	fmt.Printf("Answer to question 1 %d\n", totalSum)
	// fmt.Printf("sum of directories under %d threshhold: %d\n", 100000, totalSum)
	spaceWeNeedToFree = neededDiskSpace - (maxDiskSpace - rootDirSize)
	fmt.Printf("We need to free %d space\n", spaceWeNeedToFree)
	totalSum = 0
	guess := getSizeOfDir(&base)
	if guess < spaceWeNeedToFree {
		fmt.Printf("guess %d is smaller then required space %d", guess, spaceWeNeedToFree)
	}
	fmt.Println(guess)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
