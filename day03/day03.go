package day03

import (
    "io/ioutil"
    "os"
    "strings"
    "strconv"
)

func countMostCommonNums(lines []string, numColumns int) []int {
    counts := make([]int, numColumns)
    for _, val := range lines {
        for i, singleBit := range val {
            bitValue, _ := strconv.Atoi(string(singleBit))
            counts[i] += bitValue
        }
    }
    return counts
}


func filterBits(lines []string, column int, mostCommon bool) int64 {
    var newLines []string
    numLines := len(lines)
    numColumns := len(lines[0])
    var countCols = countMostCommonNums(lines, numColumns)
    for _, val := range lines {
        bitValue, _ := strconv.Atoi(string([]rune(val)[column]))
        compValue := countCols[column]
        if compValue * 2 == numLines {
            if mostCommon && bitValue == 1 {
                newLines = append(newLines, val)
            } else if !mostCommon && bitValue == 0 {
                newLines = append(newLines, val)
            }
        } else if compValue * 2 > numLines {
            if mostCommon && bitValue == 1 {
                newLines = append(newLines, val)
            }
            if !mostCommon && bitValue == 0 {
                newLines = append(newLines, val)
            }
        } else {
            if mostCommon && bitValue == 0 {
                newLines = append(newLines, val)
            }
            if !mostCommon && bitValue == 1 {
                newLines = append(newLines, val)
            }
        }
    }
    if len(newLines) == 1 {
        ret, _ := strconv.ParseInt(newLines[0], 2, 64)  
        return ret
    } else {
        return filterBits(newLines, column + 1, mostCommon)
    }
}

func calculateGammaEpsilon(lines []string) int64 {
    numLines := len(lines)
    numColumns := len(lines[0])
    countCols := make([]int, numColumns)
    for _, val := range lines {
        for i, singleBit := range val {
            bitValue, _ := strconv.Atoi(string(singleBit))
            countCols[i] += bitValue
        }
    }
    var gammaBuilder, epsilonBuilder strings.Builder
    for _, val := range countCols {
        if val * 2 > numLines {
            gammaBuilder.WriteString("1")
            epsilonBuilder.WriteString("0")

        } else {
            gammaBuilder.WriteString("0")
            epsilonBuilder.WriteString("1")
        }
    }
    gamma, _ := strconv.ParseInt(gammaBuilder.String(), 2, 64)  
    epsilon, _ := strconv.ParseInt(epsilonBuilder.String(), 2, 64)  
    return epsilon * gamma
}

func Day03() (int, int) {
    pwd, _ := os.Getwd()
    data, _ := ioutil.ReadFile(pwd + "/day03/input")
    file_string := string(data)
    lines := strings.Split(file_string, "\n")
    answer := calculateGammaEpsilon(lines)
    var ox, co2 int64
    ox = filterBits(lines, 0, true)
    co2 = filterBits(lines, 0, false)
    return int(answer), int(ox* co2)

}
