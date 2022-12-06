package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func count(str string, substr byte) (num int) {
	for i := 0; i < len(str); i++ {
		if str[i] == substr {
			num += 1
		}
	}
	return
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		i := 0
		for {
			// fmt.Println(line[i:i+4], i)
			if (count(line[i:i+14], line[i]) == 1) &&
				(count(line[i:i+14], line[i+1]) == 1) &&
				(count(line[i:i+14], line[i+2]) == 1) &&
				(count(line[i:i+14], line[i+3]) == 1) &&
				(count(line[i:i+14], line[i+4]) == 1) &&
				(count(line[i:i+14], line[i+5]) == 1) &&
				(count(line[i:i+14], line[i+6]) == 1) &&
				(count(line[i:i+14], line[i+7]) == 1) &&
				(count(line[i:i+14], line[i+8]) == 1) &&
				(count(line[i:i+14], line[i+9]) == 1) &&
				(count(line[i:i+14], line[i+10]) == 1) &&
				(count(line[i:i+14], line[i+11]) == 1) &&
				(count(line[i:i+14], line[i+12]) == 1) &&
				(count(line[i:i+14], line[i+13]) == 1) {
				fmt.Println(line[i:i+14], i+14)

				break
			}
			i += 1
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
