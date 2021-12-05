package main

import (
    "os"
    "aoc/day03"
    "aoc/day04"
    "aoc/day05"
    "time"
    "fmt"
)

func main() {
    tasks := map[string]func(){
        "3": day03.Day03,
        "4": day04.Day04,
        "5": day05.Day05,
    }
    
    if len(os.Args) > 1 {
        taskNum := os.Args[1]
        start := time.Now()
        tasks[taskNum]()
        duration := time.Since(start)
        fmt.Println("Time elapsed: ", duration)

    } else {
        totalTime := time.Now()
        for _, element := range tasks {
            element()
        }
        totalDuration := time.Since(totalTime)
        fmt.Println("Total time elapsed for all tasks: ", totalDuration)
    }
}

