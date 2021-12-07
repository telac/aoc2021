package day07

import (
	"io/ioutil"
	"os"
	"strconv"
	"math"
	"strings"
	//"fmt"
)

func getNumbers(line string) ([]int, map[int]int) {
	numbers := strings.Split(line, ",")
	mapNumbers := map[int]int{}
	in := make([]int, len(numbers))
	for i, s := range numbers {
		val, _ := strconv.Atoi(strings.TrimSpace(s))
		in[i] = val
		if _, exist := mapNumbers[i]; !exist {
            mapNumbers[i] = 0
        } else {
            mapNumbers[i]++
        }
	}
	return in, mapNumbers
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func findClosest(positions []int, uniquePositions map[int]int) (int, int) {
	crabEngineeredDistance := map[int]int{}
	for _, posValue := range positions {
		for key, _ := range uniquePositions {
			distance := abs(posValue - key)
			uniquePositions[key] += distance
			crabEngineeredDistance[key] += distance * (distance  + 1) / 2
			//fmt.Println(posValue, "->" , "key", key, "distance", distance, uniquePositions[key])
		}
	}
	minSeen := math.MaxInt64
	minCrabSeen := math.MaxInt64
	for k, value := range uniquePositions {
		//fmt.Println(k, value)
		if value < minSeen {
			minSeen = value
		}
		if crabEngineeredDistance[k] < minCrabSeen {
			minCrabSeen = crabEngineeredDistance[k]
		}
	}
	return minSeen, minCrabSeen
}


func Day07() (int, int) {
	pwd, _ := os.Getwd()
	data, _ := ioutil.ReadFile(pwd + "/day07/input")
	lines := strings.Split(string(data), "\n")
	numbers, m := getNumbers(lines[0])
	return findClosest(numbers, m)

}
