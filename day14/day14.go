package day14

import (
	"io/ioutil"
	"os"
	//"fmt"
	//"strconv"
	"strings"
	//"sort"
	"math"
)


func getPairs(lines []string) map[string]string {
	mappings := make(map[string]string)
	for _, line := range lines {
		if len(line) < 2 {
			continue
		}
		split := strings.Split(line, "->")
		mappings[strings.TrimSpace(split[0])] = strings.TrimSpace(split[1])
	}
	return mappings
}

func incrementMapValueByOne(m map[string]int64, k string){
	if _, ok := m[k]; !ok {
		m[k] = 1
	} else {
		m[k] += 1
	}
}

func copyMap(m map[string]int64) map[string]int64 {
	newMap := make(map[string]int64)
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

func incrementMapValueByN(m map[string]int64, to map[string]int64, k string, n int64) {
	if _, ok := to[k]; !ok {
		//fmt.Println(k, n)
		to[k] = n
	} else {
		to[k] += n
	}
}

func decrementMapValueByN(m map[string]int64, to map[string]int64, k string, n int64) {
	if _, ok := to[k]; !ok {
		to[k] = n
	} else {
		to[k] -= n
	}
}
func growPolymer(seq string, mappings map[string]string, steps int64) map[string]int64 {
	var newPolymer string
	characterCounters := make(map[string]int64)
	for i, v := range seq {
		current := string(v)
		incrementMapValueByOne(characterCounters, current)
		newPolymer += current
		if i < len(seq) -1 {
			next := string(seq[i + 1])
			key := current + next
			inBetween, ok := mappings[key]; if ok {
				newPolymer += inBetween
				incrementMapValueByOne(characterCounters, inBetween)
			}
		} else {
			continue
		}
	}
	if steps > 0 {
		return growPolymer(newPolymer, mappings, steps - 1)
	}
	return characterCounters
}

func polymerMapFromString(seq string) map[string]int64 {
	polymerMap := make(map[string]int64)
	for i, v := range seq {
		current := string(v)
		if i < len(seq) -1 {
			next := string(seq[i + 1])
			key := current + next
			incrementMapValueByOne(polymerMap, key)
		} else {
			continue
		}
	}
	return polymerMap
}

func growPolymerMap(polymerMap map[string]int64, mappings map[string]string, steps int) map[string]int64 {
	characterCounters := make(map[string]int64)
	for i := 0; i < steps; i++ {
		tMap := copyMap(polymerMap)
		for k,v := range polymerMap {
			newChar := mappings[k]
			newPairBeg := string(k[0]) + newChar
			newPairEnd := newChar + string(k[1])
			decrementMapValueByN(polymerMap, tMap, k, v)
			incrementMapValueByN(polymerMap, tMap, newPairBeg, v)
			incrementMapValueByN(polymerMap, tMap, newPairEnd, v)
			incrementMapValueByN(characterCounters, characterCounters, newChar, v)
		}
		polymerMap = copyMap(tMap)
	}
	return characterCounters
}


func min(m map[string]int64) string {
	var min int64
	min = math.MaxInt64
	minK := ""
	for k,v := range m {
		if v < min {
			min = v
			minK = k
		}
	}
	return minK
}
func max(m map[string]int64) string {
	var max int64
	max = 0
	maxK := ""
	for k,v := range m {
		if v > max {
			max = v
			maxK = k
		}
	}
	return maxK
}

func addSeqToCounts(seq string, m map[string]int64) map[string]int64 {
	for _, v := range seq {
		incrementMapValueByOne(m, string(v))
	}
	return m
}

func Day14() (int, int) {
	pwd, _ := os.Getwd()
	data, _ := ioutil.ReadFile(pwd + "/day14/input")
	parts := strings.Split(string(data), "\n\n")
	seq := parts[0]
	lines := strings.Split(string(parts[1]), "\n")
	mappings := getPairs(lines)
	counts := growPolymerMap(polymerMapFromString(seq), mappings, 10)
	smallestCount := min(counts)
	biggestCount := max(counts)
	polyMap := growPolymerMap(polymerMapFromString(seq), mappings, 40)
	polyMap = addSeqToCounts(seq, polyMap)
	smallestCountF := min(polyMap)
	biggestCountF := max(polyMap)
	return int(counts[biggestCount] - counts[smallestCount]), int(polyMap[biggestCountF] - polyMap[smallestCountF])
}
