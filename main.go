package main

import (
	"aoc/day01"
	"aoc/day02"
	"aoc/day03"
	"aoc/day04"
	"aoc/day05"
	"aoc/day06"
	"aoc/day07"
	"aoc/day08"
	"aoc/day09"
	"aoc/day10"
	"aoc/day11"
	"aoc/day12"
	"aoc/day13"
	"aoc/day14"
	"aoc/day15"
	"aoc/day16"
	"aoc/day17"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	tasks := map[int]func() (int, int){
		1:  day01.Day01,
		2:  day02.Day02,
		3:  day03.Day03,
		4:  day04.Day04,
		5:  day05.Day05,
		6:  day06.Day06,
		7:  day07.Day07,
		8:  day08.Day08,
		9:  day09.Day09,
		10: day10.Day10,
		11: day11.Day11,
		12: day12.Day12,
		13: day13.Day13,
		14: day14.Day14,
		15: day15.Day15,
		16: day16.Day16,
		17: day17.Day17,
	}

	if len(os.Args) > 1 {
		taskNum, _ := strconv.Atoi(os.Args[1])
		start := time.Now()
		answer1, answer2 := tasks[taskNum]()
		fmt.Println("Answer 1: ", answer1)
		fmt.Println("Answer 2: ", answer2)
		duration := time.Since(start)
		fmt.Println("Time elapsed: ", duration)

	} else {
		keys := make([]int, 0)
		for key, _ := range tasks {
			keys = append(keys, key)
		}
		totalTime := time.Now()
		sort.Ints(keys)
		for _, day := range keys {
			taskTimer := time.Now()
			tasks[day]()
			taskDuration := time.Since(taskTimer)
			fmt.Println("Day: ", day, " Duration: ", taskDuration)

		}
		totalDuration := time.Since(totalTime)
		fmt.Println("Total time elapsed for all tasks: ", totalDuration)
	}
}
