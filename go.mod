module aoc

go 1.17

replace (
	aoc/day01 => ./day01/day01
	aoc/day02 => ./day02/day02
	aoc/day03 => ./day03/day03
	aoc/day04 => ./day04/day04
	aoc/day05 => ./day05/day05
	aoc/day06 => ./day06/day06
	aoc/day07 => ./day07/day07
	aoc/day08 => ./day08/day08
	aoc/day09 => ./day09/day09
	aoc/day10 => ./day10/day10
	aoc/day11 => ./day11/day11
	aoc/day12 => ./day12/day12
	utils/utils => ./aoc_utils/utils
)

require github.com/yourbasic/graph v0.0.0-20210606180040-8ecfec1c2869 // indirect
