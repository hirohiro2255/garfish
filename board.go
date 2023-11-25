package main

import (
	"fmt"
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
	fmt.Println("a b c d e f g h")
	for i := 2; i < 10; i++ {
		for j := 2; j < 10; j++ {
			piece := getPieceCharacter(b.board[i][j])
			fmt.Print(piece, " ")
		}
		fmt.Println(" ", 10-i)
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
