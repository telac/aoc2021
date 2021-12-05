package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "time"
)

func readLines() []int {
    var nums []int
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) == 1 {
            if line[0] == '\n' {
                break
            }
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

func main() {
    start := time.Now()
    var nums []int = readLines()
    var answer1 = singleIncreasing(nums)
    fmt.Println("answer 1: ", answer1)
    var answer2 = windowIncreasing(nums)
    fmt.Println("answer 2: ", answer2)
    duration := time.Since(start)
    fmt.Println("time elapsed: ", duration)
}
