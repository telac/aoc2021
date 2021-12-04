package day04

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
    "strconv"
)
const nx = 5

type boardPos struct {
    value int
    marked bool
}

type winStats struct {
    board []boardPos
    isWinner bool
    lastMove int
}

func readMoves(moves string) []int {
    var stringMoveList = strings.Split(moves, ",")
    moveList := make([]int, len(stringMoveList))
    for i, v := range stringMoveList {
        var convValue, _ = strconv.Atoi(v)
        moveList[i] = convValue
    }
    return moveList
}

func readBoards(lines []string) [][]int {
    boards := [][]int{}
    var  boardC int
    boardC = -1
    board := make([]int, 25)
    var rowCount int
    rowCount = 0
    for i, row := range lines {
        if i > 0 {
            if len(row) > 2 {
                boardRow := strings.Fields(row)
                for j, val := range boardRow {
                    var convValue, _ = strconv.Atoi(val)
                    board[j + nx * rowCount] = convValue
                }
                rowCount += 1
                if rowCount == 5 {
                    tmp := make([]int, len(board))
                    copy(tmp, board)
                    boards = append(boards, tmp)

                }
            } else {
                boardC += 1
                rowCount = 0
            }
        }
    }
    return boards
}


func isWinner(markedNumbers []int, board []int) winStats {
    markedTiles := make([]boardPos, 25)
    for i, tile := range board {
        for _, move := range markedNumbers {
            if  !markedTiles[i].marked {
                var newTile = boardPos{
                    value: tile,
                    marked: move == tile,
                }
                markedTiles[i] = newTile
            }
        }
    }
    var col_counts [5]int
    var row_counts [5]int
    for i, val := range markedTiles {
        if val.marked {
            col_counts[i % nx] += 1
            row_counts[i / nx] += 1 
        }
    }
    var isWinningBoard = false
    for _, val := range col_counts {
        if val == 5 {
            isWinningBoard = true
        }
    }
    for _, val := range row_counts {
        if val == 5 {
            isWinningBoard = true
        }
    }

    var stats = winStats{
        board: markedTiles,
        isWinner: isWinningBoard,
        lastMove: markedNumbers[len(markedNumbers) - 1],
    }
    return stats
}

func getPoints(winningBoard []boardPos, move int) int {
    score := 0
    for _, val := range winningBoard {
        if !val.marked {
            score += val.value
        }
    }
    return score * move

}

func getWinner(moves []int, boards [][]int) int {
    var markedNumbers []int
    var winningPoints int
    for _, move := range moves {
        markedNumbers = append(markedNumbers, move)
        for _, board := range boards {
            stats := isWinner(markedNumbers, board)
            if stats.isWinner {
                winningPoints = getPoints(stats.board, move)
                return winningPoints
            }
        }
    }
    return winningPoints
}

func getWinningBoards(moves []int, boards [][]int) []winStats {
    var markedNumbers []int
    var winningBoards []winStats
    boardNums := make(map[int]bool)
    for _, move := range moves {
        markedNumbers = append(markedNumbers, move)
        for i, board := range boards {
            stats := isWinner(markedNumbers, board)
            if stats.isWinner {
                if _, ok := boardNums[i]; ok {
                    continue
                }
                boardNums[i] = true
                winningBoards = append(winningBoards, stats)
            }
        }
    }
    return winningBoards
}

func Day04() {
    pwd, _ := os.Getwd()
    data, _ := ioutil.ReadFile(pwd + "/day04/input")
    file_string := string(data)
    lines := strings.Split(file_string, "\n")
    moves := readMoves(lines[0])
    boards := readBoards(lines)
    val := getWinner(moves, boards)
    fmt.Println(val)

    winningBoards := getWinningBoards(moves, boards)
    fmt.Println(
        getPoints(
            winningBoards[len(winningBoards)-1].board, 
            winningBoards[len(winningBoards)-1].lastMove),
        )
}
