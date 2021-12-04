package main

import (
    "os"
    day03 "aoc/day03"
    day04 "aoc/day04"
    "time"
    "fmt"
)

func main() {
    tasks := map[string]func(){
        "3": day03.Day03,
        "4": day04.Day04,
    }
    
    if len(os.Args) > 1 {
        task_num := os.Args[1]
        start := time.Now()
        tasks[task_num]()
        duration := time.Since(start)
        fmt.Println("Time elapsed: ", duration)

    } else {
        total_time := time.Now()
        for _, element := range tasks {
            element()
        }
        total_duration := time.Since(total_time)
        fmt.Println("Total time elapsed for all tasks: ", total_duration)
    }
}

