package day08

import (
	"io/ioutil"
	"os"
	"strings"
	"fmt"
	"strconv"
	"sort"
)

type DigitData struct {
	input []string
	output []string
}

func getOutput(lines []string) []DigitData {
	digits := []DigitData{}
	for _, line := range lines {
		splitLine := strings.Split(line, "|")
		if len(line) > 0 {
			_input := strings.Fields(splitLine[0])
			_output := strings.Fields(splitLine[1])
			digits = append(digits,DigitData{
				input: _input,
				output: _output,
			})
		}
	}
	return digits
}

func countDigits(digits []DigitData) map[int]int {
	nums := make(map[int]int, 9)
	for i := 0; i < 10; i++ {
		nums[i] = 0
	}
	for _, digit := range digits {
		for _, output := range digit.output {
			switch len(output) {
			case 2:
				nums[1] += 1
			case 3:
				nums[7] += 1
			case 4:
				nums[4] += 1
			case 7:
				nums[8] += 1
			}
		}
	}
	return nums
}

func appendUnique(s string, m map[string]int) map[string]int {
	for _, k := range s {
		kS := string(k)
		_, ok := m[kS]
		if ok {
			m[kS] += 1
		} else {
			m[kS] = 1
		}
	}
	return m
}

func getCounts(s []string) map[string]int {
	sMap := make(map[string]int)
	for _, v := range s{
		sMap = appendUnique(v, sMap)
	}
	return sMap
}

func findCount(m map[string]int, count int) []string {
	var s []string
	for k,v := range m {
		if v == count {
			s = append(s, k)
		}
	}
	return s
}

func findMissing(s1 string, s2 string) []string {
	sMap := getCounts([]string{s1, s2})
	var missingRunes []string
	for k, v := range sMap {
		if v == 1 {
			missingRunes = append(missingRunes, k)
		}
	}
	return missingRunes
}

func part2(digits []DigitData) int {
	sum := 0
	nums := make(map[int]string, 9)
	for _, digit := range digits {
		for _, input := range digit.input {
			switch len(input) {
			case 2:
				nums[1] = input
			case 3:
				nums[7] = input
			case 4:
				nums[4] = input
			case 7:
				nums[8] = input
			}
		}
		var midPart string
		var midTopLeft []string
		var topLeftPart string
		var topRight string
		var bottomRight string
		//var bottom string
		//var bottomLeft string
		knownLetters := make(map[string]int)
		indicesMapping := map[string]int {
			"012345" : 0,
			"12" : 1,
			"01346" : 2,
			"01236" : 3,
			"1256" : 4,
			"02356" : 5,
			"023456" : 6,
			"012" : 7,
			"0123456": 8,
			"012356" : 9,
		}
		
		/*
		map indices like this
			0 

		5		1

			6

		4		2

			3

		*/
		// find position of top part
		topPart := findMissing(nums[1], nums[7])[0]
		knownLetters[topPart] = 0
		// find position of mid and top left parts
		for k, v := range getCounts([]string{nums[4], nums[8]}) {
			// midpoint and sides have count of 2, but we know the sides already (from num[1])
			if v == 2 && k != string(nums[1][0]) && k != string(nums[1][1]) {
				midTopLeft = append(midTopLeft, k)
			}
		}
		// find position of mid part (from number 4 + midLeft)
		for _, v := range digit.input {
			if len(v) == 5 &&
				strings.Contains(v, string(nums[1][0])) &&
				strings.Contains(v, string(nums[1][1])) &&
				strings.Contains(v, topPart) &&
				( strings.Contains(v, midTopLeft[0]) || strings.Contains(v, midTopLeft[1]) ) {
					counts := getCounts([]string{strings.Join(midTopLeft[:], ""), v})
					for k, v := range counts {
						if v == 2 {
							midPart = k
							knownLetters[k] = 6
						}
					}
			}
		}
		// find position of top left part knowing mid part
		if midTopLeft[0] == midPart {
			topLeftPart = midTopLeft[1]
		} else {
			topLeftPart = midTopLeft[0]
		}
		knownLetters[topLeftPart] = 5
		// find bottom top and right from number 5
		for _, v := range digit.input {
			if len(v) == 5 &&
				strings.Contains(v, topPart) &&
				strings.Contains(v, midPart) &&
				strings.Contains(v, topLeftPart) &&
				(strings.Contains(v, string(nums[1][1])) || strings.Contains(v, string(nums[1][0]))) {
					if strings.Contains(v,  string(nums[1][0])) {
						topRight = string(nums[1][1])
						bottomRight = string(nums[1][0])
					} else {
						topRight = string(nums[1][0])
						bottomRight = string(nums[1][1])
					}
					knownLetters[topRight] = 1
					knownLetters[bottomRight] = 2
					// we know all other numbers except for one, so that will be the bottom part
					for _, c := range v {
						cs := string(c)
						_, ok := knownLetters[cs]
						if !ok {
							knownLetters[cs] = 3
							//bottom = cs
						} 
					}
				}
		}

		for _, v := range digit.input {
			if len(v) == 7 {
				for _, c := range v {
					cs := string(c)
					_, ok := knownLetters[cs]
					if !ok {
						knownLetters[cs] = 4
						//bottomLeft = cs
					} 
				}
			}
		}
		// sigh
		var rowValue string
		for _, v := range digit.output {
			var chars string
			for _, c := range v {
			chars += fmt.Sprint(knownLetters[string(c)])
			}
			number := indicesMapping[sortString(chars)]
			rowValue += fmt.Sprint(number)
		}
		rowIntVal, _ := strconv.Atoi(rowValue)
		//fmt.Println(knownLetters, digit.output, indicesMapping)
		//fmt.Println(rowValue)
		sum += rowIntVal
		/*
		fmt.Println("top part:", topPart)
		fmt.Println("mid part:", midPart)
		fmt.Println("right", nums[1])
		fmt.Println("top left", topLeftPart)
		fmt.Println("top right", topRight)
		fmt.Println("bottom right", bottomRight)
		fmt.Println("bottom", bottom)
		fmt.Println("bottom left", bottomLeft)
		*/
	}
	return sum
}

func sortString(input string) string {
    s := []rune(input)
    sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
    return string(s)
}



func Day08() (int, int) {
	pwd, _ := os.Getwd()
	data, _ := ioutil.ReadFile(pwd + "/day08/input")
	lines := strings.Split(string(data), "\n")
	digits := getOutput(lines)
	//fmt.Println(digits)
	sum := 0
	for k, v := range countDigits(digits) {
		if k == 1 || k == 7 || k == 4 || k == 8 {
			sum += v
		}
	}
	//fmt.Println("PART 2=====")
	//fmt.Println(part2(digits))
	return sum, part2(digits)

}
