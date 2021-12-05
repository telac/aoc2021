package day01

import (
    "fmt"
    "os"
    "strconv"
    "io/ioutil"
	"strings"
)

func readLines(lines []string) []int {
    var nums []int
    for _, line := range lines {
        if len(line) < 2 {
			continue
        }
        iNum, err := strconv.Atoi(line)
        if err != nil {
           fmt.Println(err)
        }
		nums = append(nums, iNum)
    }
    return nums
}

func singleIncreasing(nums []int) int {
    var prev int
    var counter = 0
    for i, val := range nums {

        if i > 0  &&  val > prev {
            counter += 1
        }
        prev = val
    }
    return counter
}

func windowIncreasing(nums []int) int {
    var counter = 0
    for i, _ := range nums {
        if i > 2 {
            var prevSum = nums[i - 1] + nums[i - 2] + nums [i - 3]
            var currentSum = nums[i] + nums[i -1] + nums[i - 2]
            if currentSum > prevSum {
                counter += 1
            }
        }
    }
    return counter
}

func Day01() (int, int) {
	pwd, _ := os.Getwd()
	data, _ := ioutil.ReadFile(pwd + "/day01/input")
	lines := strings.Split(string(data), "\n")
    var nums []int = readLines(lines)
    var answer1 = singleIncreasing(nums)
    var answer2 = windowIncreasing(nums)
	return answer1, answer2

}
