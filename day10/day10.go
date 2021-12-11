package day10

import (
	"fmt"
	"io/ioutil"
	"os"
	//"strconv"
	"strings"
	"sort"
)

func completeLine(line string, expected []string) int {
	brackets := map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
		"<": ">",
	}
	points := 0
	for i := len(expected) - 1; i >= 0; i-- {
		points *= 5
		switch brackets[expected[i]] {
		case ")":
			points += 1
		case "]":
			points += 2
		case "}":
			points += 3
		case ">":
			points += 4
		}
	}
	return points
}

func part1(lines []string) (int, int) {
	brackets := map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
		"<": ">",
	}
	points := 0
	var points2 []int
	var incompleteLines []string
	for _, line := range lines {
		if len(line) < 2 {
			continue
		}
		var expected []string
		var corrupted = false
		var corruptedChar string
		for _, character := range line {
			sChar := string(character)
			var current string
			switch sChar {
			case "{", "(", "<", "[":
				expected = append(expected, sChar)
			case "}", ")", "]", ">":
				current, expected = expected[len(expected)-1], expected[:len(expected)-1]
				if sChar != brackets[current] {
					corrupted = true
					corruptedChar = sChar
					break
				}
			default:
				fmt.Println("not supported character: ", sChar)
			}			
		}
		if corrupted {
			switch corruptedChar {
			case ")":
				points += 3
			case "]":
				points += 57
			case "}":
				points += 1197
			case ">":
				points += 25137
			}
		} else {
			incompleteLines = append(incompleteLines, line)
			points2 = append(points2, completeLine(line,expected))
		}

	}
	sort.Ints(points2)
	midPoint := points2[len(points2) / 2]
	return points, midPoint
}


func Day10() (int, int) {
	pwd, _ := os.Getwd()
	data, _ := ioutil.ReadFile(pwd + "/day10/input")
	lines := strings.Split(string(data), "\n")
	points, incompleteLines := part1(lines)
	return points, incompleteLines
}
