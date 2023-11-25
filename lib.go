package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

const COLOR_MASK uint8 = 0b10000000
const WHITE uint8 = 0b10000000
const BLACK uint8 = 0b00000000

const PIECE_MASK uint8 = 0b00000111
const PAWN uint8 = 0b00000001
const KNIGHT uint8 = 0b00000010
const BISHOP uint8 = 0b00000011
const ROOK uint8 = 0b00000100
const QUEEN uint8 = 0b00000101
const KING uint8 = 0b00000110

const EMPTY uint8 = 0
const SENTINEL uint8 = 0b11111111

func isWhite(square uint8) bool {
	return square&COLOR_MASK == WHITE
}

func isBlack(square uint8) bool {
	return !isWhite(square)
}

func isPawn(square uint8) bool {
	return square&PIECE_MASK == PAWN
}

func isKnight(square uint8) bool {
	return square&PIECE_MASK == KNIGHT
}

func isBishop(square uint8) bool {
	return square&PIECE_MASK == BISHOP
}

func isRook(square uint8) bool {
	return square&PIECE_MASK == ROOK
}

func isQueen(square uint8) bool {
	return square&PIECE_MASK == QUEEN
}

func isKing(square uint8) bool {
	return square&PIECE_MASK == KING
}

func isEmpty(square uint8) bool {
	return square == EMPTY
}

func isOutsideBoard(square uint8) bool {
	return square == SENTINEL
}

type Board struct {
	board  [12][12]uint8
	toMove uint8
}

func (b *Board) printBoard() {
	for i := 2; i < 10; i++ {
		for j := 2; j < 10; j++ {
			piece := getPieceCharacter(b.board[i][j])
			fmt.Print(piece, " ")
		}
		fmt.Println()
	}
}

func newBoard() Board {
	b := [12][12]uint8{}
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			b[i][j] = SENTINEL
		}
	}

	b[2][2] = WHITE | ROOK
	b[2][3] = WHITE | KNIGHT
	b[2][4] = WHITE | BISHOP
	b[2][5] = WHITE | KING
	b[2][6] = WHITE | QUEEN
	b[2][7] = WHITE | BISHOP
	b[2][8] = WHITE | KNIGHT
	b[2][9] = WHITE | ROOK

	for i := 2; i < 10; i++ {
		b[3][i] = WHITE | PAWN
	}

	for i := 4; i < 8; i++ {
		for j := 2; j < 10; j++ {
			b[i][j] = EMPTY
		}
	}

	b[9][2] = BLACK | ROOK
	b[9][3] = BLACK | KNIGHT
	b[9][4] = BLACK | BISHOP
	b[9][5] = BLACK | KING
	b[9][6] = BLACK | QUEEN
	b[9][7] = BLACK | BISHOP
	b[9][8] = BLACK | KNIGHT
	b[9][9] = BLACK | ROOK

	for i := 2; i < 10; i++ {
		b[8][i] = BLACK | PAWN
	}
	return Board{
		board: b, toMove: WHITE,
	}
}

func getPieceCharacter(piece uint8) string {
	if piece == WHITE|PAWN {
		return "♙"
	} else if piece == WHITE|KNIGHT {
		return "♘"
	} else if piece == WHITE|BISHOP {
		return "♗"
	} else if piece == WHITE|ROOK {
		return "♖"
	} else if piece == WHITE|QUEEN {
		return "♕"
	} else if piece == WHITE|KING {
		return "♔"
	} else if piece == BLACK|PAWN {
		return "♟︎"
	} else if piece == BLACK|KNIGHT {
		return "♞"
	} else if piece == BLACK|BISHOP {
		return "♝"
	} else if piece == BLACK|ROOK {
		return "♜"
	} else if piece == BLACK|QUEEN {
		return "♛"
	} else if piece == BLACK|KING {
		return "♚"
	}
	return "."
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
	row := 2
	col := 2
	for _, fenRow := range fenRows {
		for _, square := range fenRow {
			if unicode.IsNumber(square) {
				squareSkipCount, err := strconv.Atoi(string(square))
				if err != nil {
					log.Fatal("Unable to convert to integer")
				}
				if squareSkipCount+col > 10 {
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
		if col != 10 {
			log.Fatal("Could not parse fen string: Complete row was not specified")
		}
		row++
		col = 2
	}
	return &Board{
		board: *b, toMove: toMove,
	}, nil
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
		return WHITE | KNIGHT
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
