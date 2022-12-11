package main

import "fmt"

type monkeyData struct {
	number      int
	items       chan int
	operation   func(int) int
	divisibleBy int
	tCase       int
	fCase       int
	activity    int
}

func generateMonkey(num, tCase, fCase, divisibleBy, operationConst int, operation int, items []int) *monkeyData {
	tmp := &monkeyData{
		number: num,
		items:  make(chan int, 100),
		operation: func(i int) int {
			if operation == 0 {
				return i + operationConst
			}
			if operation == 1 {
				return i * operationConst
			}
			return i * i
		},
		tCase:       tCase,
		fCase:       fCase,
		divisibleBy: divisibleBy,
		activity:    0,
	}
	for _, v := range items {
		tmp.items <- v
	}
	return tmp
}

//----------
/* Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
*/
func testCase() []*monkeyData {
	var monkeyList []*monkeyData
	m0 := generateMonkey(0, 2, 3, 23, 19, 1, []int{79, 98})
	monkeyList = append(monkeyList, m0)
	m1 := generateMonkey(1, 2, 0, 19, 6, 0, []int{54, 65, 75, 74})
	monkeyList = append(monkeyList, m1)
	m2 := generateMonkey(2, 1, 3, 13, 0, 2, []int{79, 60, 97})
	monkeyList = append(monkeyList, m2)
	m3 := generateMonkey(3, 0, 1, 17, 3, 0, []int{74})
	monkeyList = append(monkeyList, m3)
	return monkeyList
}

/*
Monkey 0:
  Starting items: 89, 73, 66, 57, 64, 80
  Operation: new = old * 3
  Test: divisible by 13
    If true: throw to monkey 6
    If false: throw to monkey 2

	m0 := generateMonkey(0, 6, 2, 13, 3, 1, []int{89, 73, 66, 57, 64, 80})

Monkey 1:
	Starting items: 83, 78, 81, 55, 81, 59, 69
	Operation: new = old + 1
	Test: divisible by 3
    If true: throw to monkey 7
    If false: throw to monkey 4

	m1 := generateMonkey(1, 7, 4, 3, 1, 0, []int{83, 78, 81, 55, 81, 59, 69})
Monkey 2:
  Starting items: 76, 91, 58, 85
  Operation: new = old * 13
  Test: divisible by 7
    If true: throw to monkey 1
    If false: throw to monkey 4

	m2 := generateMonkey(2, 1, 4, 7, 13, 1, []int{76, 91, 58, 85})

Monkey 3:
	Starting items: 71, 72, 74, 76, 68
	Operation: new = old * old
	Test: divisible by 2
    If true: throw to monkey 6
    If false: throw to monkey 0

	m3 := generateMonkey(3, 6, 0, 2, 0, 2, []int{71, 72, 74, 76, 68})

Monkey 4:
	Starting items: 98, 85, 84
	Operation: new = old + 7
	Test: divisible by 19
    If true: throw to monkey 5
    If false: throw to monkey 7

	m4 := generateMonkey(4, 5, 7, 19, 7, 0, []int{98, 85, 84})

	Monkey 5:
	Starting items: 78
	Operation: new = old + 8
	Test: divisible by 5
    If true: throw to monkey 3
    If false: throw to monkey 0

	m5 := generateMonkey(5, 3, 0, 5, 8, 0, []int{78})

	Monkey 6:
	Starting items: 86, 70, 60, 88, 88, 78, 74, 83
	Operation: new = old + 4
	Test: divisible by 11
    If true: throw to monkey 1
    If false: throw to monkey 2

	m6 := generateMonkey(6, 1, 2, 11, 4, 0, []int{86, 70, 60, 88, 88, 78, 74, 83})

Monkey 7:
  Starting items: 81, 58
  Operation: new = old + 5
  Test: divisible by 17
  	If true: throw to monkey 3
  	If false: throw to monkey 5

  m7 := generateMonkey(7, 3, 5, 17, 5, 0, []int{81,58})
*/

func main() {
	var monkeyList []*monkeyData
	m0 := generateMonkey(0, 6, 2, 13, 3, 1, []int{89, 73, 66, 57, 64, 80})
	monkeyList = append(monkeyList, m0)
	m1 := generateMonkey(1, 7, 4, 3, 1, 0, []int{83, 78, 81, 55, 81, 59, 69})
	monkeyList = append(monkeyList, m1)
	m2 := generateMonkey(2, 1, 4, 7, 13, 1, []int{76, 91, 58, 85})
	monkeyList = append(monkeyList, m2)
	m3 := generateMonkey(3, 6, 0, 2, 0, 2, []int{71, 72, 74, 76, 68})
	monkeyList = append(monkeyList, m3)
	m4 := generateMonkey(4, 5, 7, 19, 7, 0, []int{98, 85, 84})
	monkeyList = append(monkeyList, m4)
	m5 := generateMonkey(5, 3, 0, 5, 8, 0, []int{78})
	monkeyList = append(monkeyList, m5)
	m6 := generateMonkey(6, 1, 2, 11, 4, 0, []int{86, 70, 60, 88, 88, 78, 74, 83})
	monkeyList = append(monkeyList, m6)
	m7 := generateMonkey(7, 3, 5, 17, 5, 0, []int{81, 58})
	monkeyList = append(monkeyList, m7)

	monkeyList = testCase()

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeyList {
			for len(monkey.items) > 0 {
				item := <-monkey.items
				score := monkey.operation(item)
				// score /= 3
				if (score % monkey.divisibleBy) == 0 {
					monkeyList[monkey.tCase].items <- score
				} else {
					monkeyList[monkey.fCase].items <- score
				}
				monkey.activity += 1
			}
		}
	}
	var firstMax, secondMax int
	for i, monkey := range monkeyList {
		for len(monkey.items) > 0 {
			v := <-monkey.items
			fmt.Printf("%d,", v)
		}
		fmt.Printf("\n")
		fmt.Printf("Monkey: %d activity: %d\n", i, monkey.activity)
		if monkey.activity > firstMax {
			secondMax = firstMax
			firstMax = monkey.activity
			continue
		}
		if monkey.activity > secondMax {
			secondMax = monkey.activity
		}
	}
	fmt.Printf("Result %d", firstMax*secondMax)
}
