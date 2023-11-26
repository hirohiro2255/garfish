package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func knightMoves(row int8, col int8, color uint8, board *Board, moves *[][]int8) {
	cords := [][]int8{{1, 2}, {1, -2}, {2, 1}, {2, -1}, {-1, 2}, {-1, -2}, {-2, -1}, {-2, 1}}
	for _, mods := range cords {
		_row := row + mods[0]
		_col := col + mods[1]
		space := board.board[_row][_col]
		if isEmpty(space) || (space&COLOR_MASK) != color {
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
