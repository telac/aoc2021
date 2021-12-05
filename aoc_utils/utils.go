package utils

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func ReadLines() []int {
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

func ReadStrLines() []string {
    var lines []string
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) == 1 {
            if line[0] == '\n' {
                break
            }
        }
        nums = append(lines, line)
    }
    return lines
}

type Solution struct {
    part1 int
    part2 int
}