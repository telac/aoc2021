package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "time"
)

func readStrLines() []string {
    var lines []string
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) == 1 {
            if line[0] == '\n' {
                break
            }
        }
        lines = append(lines, line)
    }
    return lines
}

type Position struct {
    depth int
    horizontalDistance int
    multiplier int
}

func calculatePosition(lines []string, useAim bool) Position {
    var depth = 0
    var horizontalDistance = 0
    var aim = 0
    for _, line := range lines {
        items := strings.Fields(line)
        var direction = items[0]
        magnitude, err := strconv.Atoi(items[1])
        if err != nil {
           fmt.Println(err)
        }
        if useAim {
            switch direction {
            case "up":
                aim -= magnitude
            case "down":
                aim += magnitude
            case "forward":
                horizontalDistance += magnitude
                depth += magnitude * aim
            }
        } else {
            switch direction {
            case "up":
                depth -= magnitude
            case "down":
                depth += magnitude
            case "forward":
                horizontalDistance += magnitude
            }
        }
    }
    var finalPosition = Position{
            depth: depth,
            horizontalDistance: horizontalDistance,
            multiplier: depth*horizontalDistance,
        }
    return finalPosition
}

func main() {
    start := time.Now()
    var lines []string = readStrLines()
    answer1 := calculatePosition(lines, false)
    fmt.Println("answer 1: ", answer1)
    answer2 := calculatePosition(lines, true)
    fmt.Println("answer 2: ", answer2)
    duration := time.Since(start)
    fmt.Println("time elapsed: ", duration)
}
