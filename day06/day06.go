package day06

import (
    "io/ioutil"
    "os"
    "strings"
    "strconv"
	"sync"
	//"fmt"
)

func getNumbers(line string) []int {
	numbers := strings.Split(line, ",")
	in := make([]int, len(numbers))
	for i, s := range numbers {
		val, _ := strconv.Atoi(strings.TrimSpace(s))
		in[i] = val
		
	}
	return in
}

func simulateFish(startNum int, numDays int, index int, result []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var fishSlice []int
	fishSlice = append(fishSlice, startNum)
	for i := 0; i < numDays; i++ {
		newFish := 0
		for i, f := range fishSlice {
			if f == 0 {
				newFish++
				fishSlice[i] = 6
			} else {
				fishSlice[i]--
			}
		}
		for i := 0; i < newFish; i++ {
			fishSlice = append(fishSlice, 8)
		}
	}
	result[index] = len(fishSlice)
}


func simulateFishMap(startNum int, numDays int, index int, result []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var fishMap = make(map[int]int, 9)
	for i := 0; i < 9; i++ {
		fishMap[i] = 0
	}

	fishMap[startNum] += 1
	for i := 0; i < numDays; i++ {
		var tempMap = make(map[int]int, 9)
		for i := 0; i < 9; i++ {
			tempMap[i] = 0
		}
		for key, _ := range fishMap {
			if key == 0  {
				tempMap[6] += fishMap[0]
				tempMap[8] += fishMap[0]
			} else {
				tempMap[key - 1] += fishMap[key]
			}

		}
		/*
		if index == 0 {
			fmt.Println("before", fishMap)
		}
		*/
		for key, _ := range tempMap {
			fishMap[key] = tempMap[key]
		}
				/*
		if index == 0 {
			fmt.Println("after", fishMap)
		}
		*/

	}


	var sum int
	for _, val := range fishMap {
		sum += val
	}
	result[index] = sum
}

func task1(numbers []int) int{
	results := make([]int, len(numbers))
	var wg sync.WaitGroup
	for i, val := range numbers {
		wg.Add(1)
		go simulateFishMap(val, 80, i, results, &wg)
	}
	wg.Wait()
	var sum int
	for _, val := range results {
		sum += val
	}
	return sum
}

func task2(numbers []int) int {
	results := make([]int, len(numbers))
	var wg sync.WaitGroup
	for i, val := range numbers {
		wg.Add(1)
		go simulateFishMap(val, 256, i, results, &wg)
	}
	wg.Wait()
	var sum int
	for _, val := range results {
		sum += val
	}
	return sum
}


func Day06() (int, int) {
	pwd, _ := os.Getwd()
    data, _ := ioutil.ReadFile(pwd + "/day06/input")
	lines := strings.Split(string(data), "\n")
	numbers := getNumbers(lines[0])
	return task1(numbers), task2(numbers)
		
}
