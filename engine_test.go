package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyBoard(t *testing.T) {
	b, err := boardFromFen("8/8/8/8/8/8/8/8 w KQkq - 0 1")
	if err != nil {
		log.Fatal("Unable to read fen string")
	}
	for i := 2; i < 10; i++ {
		for j := 2; j < 10; j++ {
			assert.Equal(t, EMPTY, b.board[i][j])
		}
	}
}

func TestStartingPos(t *testing.T) {
	b, err := boardFromFen(DEFAULT_POS)
	if err != nil {
		log.Fatal("Unable to read fen string")
	}
	assert.Equal(t, BLACK|ROOK, b.board[2][2])
	assert.Equal(t, BLACK|KNIGHT, b.board[2][3])
	assert.Equal(t, BLACK|BISHOP, b.board[2][4])
	assert.Equal(t, BLACK|QUEEN, b.board[2][5])
	assert.Equal(t, BLACK|KING, b.board[2][6])
	assert.Equal(t, BLACK|BISHOP, b.board[2][7])
	assert.Equal(t, BLACK|KNIGHT, b.board[2][8])
	assert.Equal(t, BLACK|ROOK, b.board[2][9])

	for i := 2; i < 10; i++ {
		assert.Equal(t, BLACK|PAWN, b.board[3][i])
	}

	for i := 4; i < 8; i++ {
		for j := 2; j < 10; j++ {
			assert.Equal(t, EMPTY, b.board[i][j])
		}
	}

	assert.Equal(t, WHITE|ROOK, b.board[9][2])
	assert.Equal(t, WHITE|KNIGHT, b.board[9][3])
	assert.Equal(t, WHITE|BISHOP, b.board[9][4])
	assert.Equal(t, WHITE|QUEEN, b.board[9][5])
	assert.Equal(t, WHITE|KING, b.board[9][6])
	assert.Equal(t, WHITE|BISHOP, b.board[9][7])
	assert.Equal(t, WHITE|KNIGHT, b.board[9][8])
	assert.Equal(t, WHITE|ROOK, b.board[9][9])

	for i := 2; i < 10; i++ {
		assert.Equal(t, WHITE|PAWN, b.board[8][i])
	}
}

func TestCorrectStartingPlayer(t *testing.T) {
	b, err := boardFromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		log.Fatal("Unable to read fen string")
	}
	assert.Equal(t, WHITE, b.toMove)
	b, err = boardFromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1")
	if err != nil {
		log.Fatal("Unable to read fen string")
	}
	assert.Equal(t, BLACK, b.toMove)

}

func TestRandomPos(t *testing.T) {
	b, err := boardFromFen("4R1B1/1kp5/1B1Q4/1P5p/1p2p1pK/8/3pP3/4N1b1 w - - 0 1")
	if err != nil {
		log.Fatal("Unable to read fen string")
	}
	assert.Equal(t, WHITE|ROOK, b.board[2][6])
	assert.Equal(t, WHITE|BISHOP, b.board[2][8])
	assert.Equal(t, BLACK|KING, b.board[3][3])
	assert.Equal(t, BLACK|PAWN, b.board[3][4])
	assert.Equal(t, WHITE|BISHOP, b.board[4][3])
	assert.Equal(t, WHITE|QUEEN, b.board[4][5])
	assert.Equal(t, WHITE|PAWN, b.board[5][3])
	assert.Equal(t, BLACK|PAWN, b.board[5][9])
	assert.Equal(t, BLACK|PAWN, b.board[6][3])
	assert.Equal(t, BLACK|PAWN, b.board[6][6])
	assert.Equal(t, WHITE|KING, b.board[6][9])
	assert.Equal(t, BLACK|PAWN, b.board[8][5])
	assert.Equal(t, WHITE|PAWN, b.board[8][6])
	assert.Equal(t, WHITE|KNIGHT, b.board[9][6])
	assert.Equal(t, BLACK|BISHOP, b.board[9][8])
}

func TestKnightMovesEmptyBoard(t *testing.T) {
	b, err := boardFromFen("8/8/8/8/3N4/8/8/8 w - - 0 1")
	if err != nil {
		log.Fatal("Unable to read fen string")
	}
	var row int8 = 6
	var col int8 = 5
	ret := [][]int8{}
	knightMoves(row, col, WHITE, b, &ret)
	assert.Equal(t, 8, len(ret))
}

func TestKnightMovesCorner(t *testing.T) {
	b, err := boardFromFen("N7/8/8/8/8/8/8/8 w - - 0 1")
	if err != nil {
		log.Fatal("Unable to read fen string")
	}
	var row int8 = 2
	var col int8 = 2
	ret := [][]int8{}
	knightMoves(row, col, WHITE, b, &ret)
	assert.Equal(t, 2, len(ret))
}

func TestKnightMovesWithOtherPiecesWithCapture(t *testing.T) {
	b, err := boardFromFen("8/8/5n2/3NQ3/2K2P2/8/8/8 w - - 0 1")
	if err != nil {
		log.Fatal("Unable to read fen string")
	}
	var row int8 = 5
	var col int8 = 5
	ret := [][]int8{}
	knightMoves(row, col, WHITE, b, &ret)
	assert.Equal(t, 7, len(ret))
}
