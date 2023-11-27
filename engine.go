package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func rookMoves(row int8, col int8, piece uint8, board *Board, moves *[][]int8) {
	rowStart := row + 1
	for isEmpty(board.board[rowStart][col]) {
		*moves = append(*moves, []int8{rowStart, col})
		rowStart++
	}

	if !isOutsideBoard(board.board[rowStart][col]) && piece&COLOR_MASK != board.board[rowStart][col]&COLOR_MASK {
		*moves = append(*moves, []int8{rowStart + 1, col})
	}

	rowStart = row - 1
	for isEmpty(board.board[rowStart][col]) {
		*moves = append(*moves, []int8{rowStart, col})
		rowStart--
	}
	if !isOutsideBoard(board.board[rowStart][col]) && piece&COLOR_MASK != board.board[rowStart][col]&COLOR_MASK {
		*moves = append(*moves, []int8{rowStart - 1, col})
	}

	colStart := col + 1
	for isEmpty(board.board[row][colStart]) {
		*moves = append(*moves, []int8{row, colStart})
		colStart++
	}

	if !isOutsideBoard(board.board[row][colStart]) && piece&COLOR_MASK != board.board[row][colStart]&COLOR_MASK {
		*moves = append(*moves, []int8{row, colStart + 1})
	}

	colStart = col - 1
	for isEmpty(board.board[row][colStart]) {
		*moves = append(*moves, []int8{row, colStart})
		colStart--
	}

	if !isOutsideBoard(board.board[row][colStart]) && piece&COLOR_MASK != board.board[row][colStart]&COLOR_MASK {
		*moves = append(*moves, []int8{row, colStart - 1})
	}
}

func kingMoves(row int8, col int8, piece uint8, board *Board, moves *[][]int8) {
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			r := row + int8(i)
			c := col + int8(j)

			if isOutsideBoard(board.board[r][c]) {
				continue
			}

			if isEmpty(board.board[r][c]) || board.board[r][c]&COLOR_MASK != piece&COLOR_MASK {
				*moves = append(*moves, []int8{r, c})
			}

		}
	}
}

func pawnMoves(row int8, col int8, piece uint8, board *Board, moves *[][]int8) {
	// white pawns move up board
	if isWhite(piece) {
		// check capture
		leftCap := board.board[row-1][col-1]
		rightCap := board.board[row-1][col+1]
		if !isOutsideBoard(leftCap) && isBlack(leftCap) {
			to_cord := [][]int8{{row - 1, col - 1}}
			*moves = append(*moves, to_cord...)
		}
		if !isOutsideBoard(rightCap) && isBlack(rightCap) {
			to_cord := [][]int8{{row - 1, col + 1}}
			*moves = append(*moves, to_cord...)
		}

		// check a normal push
		if isEmpty(board.board[row-1][col]) {
			toCord := []int8{row - 1, col}
			*moves = append(*moves, toCord)
		}

		// check a double push
		if row == 8 && isEmpty(board.board[row-1][col]) && isEmpty(board.board[row-2][col]) {
			*moves = append(*moves, []int8{row - 2, col})
		}
	} else {
		// black to move
		// check capture
		leftCap := board.board[row+1][col+1]
		rightCap := board.board[row+1][col-1]
		if !isOutsideBoard(leftCap) && isWhite(leftCap) {
			*moves = append(*moves, []int8{row + 1, col + 1})
		}
		if !isOutsideBoard(rightCap) && isWhite(rightCap) {
			*moves = append(*moves, []int8{row + 1, col - 1})
		}

		// check a normal push
		if isEmpty(board.board[row+1][col]) {
			*moves = append(*moves, []int8{row + 1, col})
		}

		// check a double push
		if row == 3 && isEmpty(board.board[row+1][col]) && isEmpty(board.board[row+2][col]) {
			*moves = append(*moves, []int8{row + 2, col})
		}
	}
}

func knightMoves(row int8, col int8, piece uint8, board *Board, moves *[][]int8) {
	cords := [][]int8{{1, 2}, {1, -2}, {2, 1}, {2, -1}, {-1, 2}, {-1, -2}, {-2, -1}, {-2, 1}}
	for _, mods := range cords {
		_row := row + mods[0]
		_col := col + mods[1]
		space := board.board[_row][_col]
		if isOutsideBoard(space) {
			continue
		}
		if isEmpty(space) || (space&COLOR_MASK) != piece&COLOR_MASK {
			to_cord := [][]int8{{_row, _col}}
			*moves = append(*moves, to_cord...)
		}
	}
}

func getPieceFromFenStringChar(piece rune) uint8 {
	if piece == 'p' {
		return BLACK | PAWN
	} else if piece == 'n' {
		return BLACK | KNIGHT
	} else if piece == 'b' {
		return BLACK | BISHOP
	} else if piece == 'r' {
		return BLACK | ROOK
	} else if piece == 'q' {
		return BLACK | QUEEN
	} else if piece == 'k' {
		return BLACK | KING
	} else if piece == 'P' {
		return WHITE | PAWN
	} else if piece == 'K' {
		return WHITE | KING
	} else if piece == 'B' {
		return WHITE | BISHOP
	} else if piece == 'N' {
		return WHITE | KNIGHT
	} else if piece == 'R' {
		return WHITE | ROOK
	} else if piece == 'Q' {
		return WHITE | QUEEN
	} else if piece == 'K' {
		return WHITE | KING
	}
	return SENTINEL
}

func boardFromFen(fen string) (*Board, error) {
	b := &[12][12]uint8{}
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			b[i][j] = SENTINEL
		}
	}
	fenConfig := strings.Split(fen, " ")
	if len(fenConfig) != 6 {
		return nil, errors.New("Could not parse fen string: Invalid fen string")
	}
	var toMove uint8
	if fenConfig[1] == "w" {
		toMove = WHITE
	} else {
		toMove = BLACK
	}

	/*
		castlingPrivileges := fenConfig[2]
		enPassant := fenConfig[3]
		halfmoveClock := fenConfig[4]
		fullmoveClock := fenConfig[5]
	*/

	fenRows := strings.Split(fenConfig[0], "/")
	if len(fenRows) != 8 {
		return nil, errors.New("Could not parse fen string: Invalid number of rows provided, 8 expected")
	}
	row := BOARD_START
	col := BOARD_START
	for _, fenRow := range fenRows {
		for _, square := range fenRow {
			if unicode.IsNumber(square) {
				squareSkipCount, err := strconv.Atoi(string(square))
				if err != nil {
					log.Fatal("Unable to convert to integer")
				}
				if squareSkipCount+col > BOARD_END {
					log.Fatal("Could not parse fen string: Index out of bounds")
				}
				for squareSkipCount > 0 {
					b[row][col] = EMPTY
					col++
					squareSkipCount -= 1
				}
			} else {
				piece := getPieceFromFenStringChar(square)
				if piece != SENTINEL {
					b[row][col] = piece
				} else {
					fmt.Println("piece:", piece)
					log.Fatal("Could not parse fen string: Invalid character found")
				}
				col++
			}
		}
		if col != BOARD_END {
			log.Fatal("Could not parse fen string: Complete row was not specified")
		}
		row++
		col = BOARD_START
	}
	return &Board{
		board: *b, toMove: toMove,
	}, nil
}
