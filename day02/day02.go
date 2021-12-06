package day02

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	depth              int
	horizontalDistance int
	multiplier         int
}

func calculatePosition(lines []string, useAim bool) Position {
	var depth = 0
	var horizontalDistance = 0
	var aim = 0
	for _, line := range lines {
		if len(line) < 2 {
			continue
		}
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
		depth:              depth,
		horizontalDistance: horizontalDistance,
		multiplier:         depth * horizontalDistance,
	}
	return finalPosition
}

func Day02() (int, int) {
	pwd, _ := os.Getwd()
	data, _ := ioutil.ReadFile(pwd + "/day02/input")
	lines := strings.Split(string(data), "\n")
	answer1 := calculatePosition(lines, false)
	answer2 := calculatePosition(lines, true)
	return answer1.multiplier, answer2.multiplier

}
