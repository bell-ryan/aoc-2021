package bingo

import (
	"aoc/inputs"
	"aoc/utils"
	"strconv"
	"strings"
)

type board struct {
	winner bool
	matrix [5][5]int // create a matrix with all 0
	actual [5][5]string
	Score  int
}

func (b *board) checkWinner() {
	for row := 0; row < 5; row++ {
		rowTotal := 0
		colTotal := 0
		for col := 0; col < 5; col++ {
			rowTotal += b.matrix[row][col]
			colTotal += b.matrix[col][row]
		}
		if rowTotal == 5 || colTotal == 5 {
			b.winner = true
			return
		}
	}

}
func (b *board) checkNumber(num string) {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if num == b.actual[row][col] {
				n, err := strconv.Atoi(num)
				utils.ErrorCheck(err)

				b.Score -= n
				b.matrix[row][col] = 1

			}
		}
	}

}

func buildBoards(data string) ([]string, []board) {
	numbersAndBoards := strings.Split(data, "\n\n")
	numbers := strings.Split(numbersAndBoards[0], ",")
	boards := numbersAndBoards[1:]

	boardList := make([]board, len(boards))

	for bi, boardData := range boards {
		b := board{}
		rows := strings.Split(boardData, "\n")
		for ri, row := range rows {
			row = strings.ReplaceAll(row, "  ", " ")
			row = strings.TrimLeft(row, " ")
			cols := strings.Split(row, " ")
			for ci, col := range cols {
				b.actual[ri][ci] = col
				colInt, err := strconv.Atoi(col)
				utils.ErrorCheck(err)

				b.Score += colInt // calculate max possible Score of a board
				boardList[bi] = b

			}
		}

	}
	return numbers, boardList
}

func Part1() board {
	data := inputs.GetInputData(4)
	numbers, boardList := buildBoards(data)

	var winner board

out:
	for _, num := range numbers {
		for i := range boardList {
			board := &boardList[i]
			board.checkNumber(num)
			board.checkWinner()
			if board.winner {
				n, err := strconv.Atoi(num)
				utils.ErrorCheck(err)
				board.Score *= n
				winner = *board
				break out
			}
		}

	}

	return winner
}
