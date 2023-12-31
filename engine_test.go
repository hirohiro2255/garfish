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
	knightMoves(row, col, WHITE|KNIGHT, b, &ret)
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
	knightMoves(row, col, WHITE|KNIGHT, b, &ret)
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
	knightMoves(row, col, WHITE|KNIGHT, b, &ret)
	assert.Equal(t, 7, len(ret))
}

func TestWhitePawnStart(t *testing.T) {
	b, err := boardFromFen("8/8/8/8/8/8/P7/8 w - - 0 1")
	if err != nil {
		log.Fatal("Unable to read fen string")
	}
	var row int8 = 8
	var col int8 = 2
	ret := [][]int8{}
	pawnMoves(row, col, WHITE|PAWN, b, &ret)
	assert.Equal(t, 2, len(ret))
}

func TestWhitePawnhasMoved(t *testing.T) {
	b, err := boardFromFen("8/8/8/8/8/3P4/8/8 w - - 0 1")
	if err != nil {
		log.Fatal("Unable to read fen string")
	}
	var row int8 = 7
	var col int8 = 5
	ret := [][]int8{}
	pawnMoves(row, col, WHITE|PAWN, b, &ret)
	assert.Equal(t, 1, len(ret))
}

func TestWhitePawnCantMoveBlackPieceBlock(t *testing.T) {
	b, _ := boardFromFen("8/8/8/8/3r4/3P4/8/8 w - - 0 1")
	var row int8 = 7
	var col int8 = 5
	ret := [][]int8{}
	pawnMoves(row, col, WHITE|PAWN, b, &ret)
	assert.Equal(t, 0, len(ret))
}

func TestWhitePawnCantMoveWhitePieceBlock(t *testing.T) {
	b, _ := boardFromFen("8/8/8/8/3K4/3P4/8/8 w - - 0 1")
	var row int8 = 7
	var col int8 = 5
	ret := [][]int8{}
	pawnMoves(row, col, WHITE|PAWN, b, &ret)
	assert.Equal(t, 0, len(ret))
}

func TestWhitePawnWithTwoCapturesAndStart(t *testing.T) {
	b, _ := boardFromFen("8/8/8/8/8/n1q5/1P6/8 w - - 0 1")
	var row int8 = 8
	var col int8 = 3
	ret := [][]int8{}
	pawnMoves(row, col, WHITE|PAWN, b, &ret)
	assert.Equal(t, 4, len(ret))
}

func TestWhitePawnWithOneCapture(t *testing.T) {
	b, _ := boardFromFen("8/8/Q1b5/1P6/8/8/8/8 w - - 0 1")
	var row int8 = 5
	var col int8 = 3
	ret := [][]int8{}
	pawnMoves(row, col, WHITE|PAWN, b, &ret)
	assert.Equal(t, 2, len(ret))

}

func TestWhitePawnDoublePushPieceInFront(t *testing.T) {
	b, _ := boardFromFen("8/8/8/8/8/b7/P7/8 w - - 0 1")
	var row int8 = 8
	var col int8 = 2
	ret := [][]int8{}
	pawnMoves(row, col, WHITE|PAWN, b, &ret)
	assert.Equal(t, 0, len(ret))
}

func TestBlackPawnDoublePush(t *testing.T) {
	b, _ := boardFromFen("8/p7/8/8/8/8/8/8 w - - 0 1")
	var row int8 = 3
	var col int8 = 2
	ret := [][]int8{}
	pawnMoves(row, col, BLACK|PAWN, b, &ret)
	assert.Equal(t, 2, len(ret))
}

func TestBlackPawnHasMoved(t *testing.T) {
	b, _ := boardFromFen("8/8/8/3p4/8/8/8/8 w - - 0 1")
	var row int8 = 5
	var col int8 = 5
	ret := [][]int8{}
	pawnMoves(row, col, BLACK|PAWN, b, &ret)
	assert.Equal(t, 1, len(ret))
}

func TestBlackPawnCantMoveWhitePieceBlock(t *testing.T) {
	b, _ := boardFromFen("8/3p4/3R4/8/8/8/8/8 w - - 0 1")
	var row int8 = 3
	var col int8 = 5
	ret := [][]int8{}
	pawnMoves(row, col, BLACK|PAWN, b, &ret)
	assert.Equal(t, 0, len(ret))
}

func TestBlackPawnWithTwoCapturesAndStart(t *testing.T) {
	b, _ := boardFromFen("8/3p4/2R1R3/8/8/8/8/8 w - - 0 1")
	var row int8 = 3
	var col int8 = 5
	ret := [][]int8{}
	pawnMoves(row, col, BLACK|PAWN, b, &ret)
	assert.Equal(t, 4, len(ret))
}

func TestBlackPawnWithOneCapture(t *testing.T) {
	b, _ := boardFromFen("8/3p4/3qR3/8/8/8/8/8 w - - 0 1")
	var row int8 = 3
	var col int8 = 5
	ret := [][]int8{}
	pawnMoves(row, col, BLACK|PAWN, b, &ret)
	assert.Equal(t, 1, len(ret))
}

func TestKingEmptyBoardCenter(t *testing.T) {
	b, _ := boardFromFen("8/8/8/8/3K4/8/8/8 w - - 0 1")
	var row int8 = 5
	var col int8 = 6
	ret := [][]int8{}
	kingMoves(row, col, WHITE|KING, b, &ret)
	assert.Equal(t, 8, len(ret))
}

func TestKingStartPos(t *testing.T) {
	b, _ := boardFromFen("8/8/8/8/8/8/8/4K3 w - - 0 1")
	var row int8 = 9
	var col int8 = 6
	ret := [][]int8{}
	kingMoves(row, col, WHITE|KING, b, &ret)
	assert.Equal(t, 5, len(ret))
}

func TestStartPosOtherPieces(t *testing.T) {
	b, _ := boardFromFen("8/8/8/8/8/8/3Pn3/3QKB2 w - - 0 1")
	var row int8 = 9
	var col int8 = 6
	ret := [][]int8{}
	kingMoves(row, col, WHITE|KING, b, &ret)
	assert.Equal(t, 2, len(ret))
}

func TestKingBlackOtherPieces(t *testing.T) {
	b, _ := boardFromFen("8/8/8/8/8/3Pn3/3QkB2/3R1q2 w - - 0 1")
	var row int8 = 8
	var col int8 = 6
	ret := [][]int8{}
	kingMoves(row, col, BLACK|KING, b, &ret)
	assert.Equal(t, 6, len(ret))
}

func TestRookCenterOfEmptyBoard(t *testing.T) {
	b, _ := boardFromFen("8/8/8/8/3R4/8/8/8 w - - 0 1")
	var row int8 = 6
	var col int8 = 5
	ret := [][]int8{}
	rookMoves(row, col, WHITE|ROOK, b, &ret)
	assert.Equal(t, 14, len(ret))
}

func TestRookCenterOfBoard(t *testing.T) {
	b, _ := boardFromFen("8/8/8/3q4/2kRp3/3b4/8/8 w - - 0 1")
	var row int8 = 6
	var col int8 = 5
	ret := [][]int8{}
	rookMoves(row, col, WHITE|ROOK, b, &ret)
	assert.Equal(t, 4, len(ret))
}

func TestRookCenterOfBoardWithWhitePieces(t *testing.T) {
	b, _ := boardFromFen("7p/3N4/8/4n3/2kR4/3b4/8/8 w - - 0 1")
	var row int8 = 6
	var col int8 = 5
	ret := [][]int8{}
	rookMoves(row, col, WHITE|ROOK, b, &ret)
	assert.Equal(t, 8, len(ret))
}

func TestRookCorner(t *testing.T) {
	b, _ := boardFromFen("7p/3N4/8/4n3/2kR4/3b4/8/8 w - - 0 1")
	var row int8 = 9
	var col int8 = 9
	ret := [][]int8{}
	rookMoves(row, col, WHITE|ROOK, b, &ret)
	assert.Equal(t, 14, len(ret))
}

func TestBlackRookCenterOfBoardWithWhitePieces(t *testing.T) {
	b, _ := boardFromFen("7p/3N4/8/4n3/2kR4/3b4/8/8 w - - 0 1")
	var row int8 = 6
	var col int8 = 5
	ret := [][]int8{}
	rookMoves(row, col, BLACK|ROOK, b, &ret)
	assert.Equal(t, 7, len(ret))
}

func TestBlackBishopCenterEmptyBoard(t *testing.T) {
	b, _ := boardFromFen("8/8/8/3b4/8/8/8/8 w - - 0 1")
	var row int8 = 5
	var col int8 = 5
	ret := [][]int8{}
	bishopMoves(row, col, BLACK|BISHOP, b, &ret)
	assert.Equal(t, 13, len(ret))
}

func TestBlackBishopCenterWithCaptures(t *testing.T) {
	b, _ := boardFromFen("6P1/8/8/3b4/8/1R6/8/3Q4 w - - 0 1")
	var row int8 = 5
	var col int8 = 5
	ret := [][]int8{}
	bishopMoves(row, col, BLACK|BISHOP, b, &ret)
	assert.Equal(t, 12, len(ret))
}

func TestBlackBishopCenterWithCapturesAndPieces(t *testing.T) {
	b, _ := boardFromFen("6P1/8/2Q5/3b4/2k1n3/1R6/8/b2Q4 w - - 0 1")
	var row int8 = 5
	var col int8 = 5
	ret := [][]int8{}
	bishopMoves(row, col, BLACK|BISHOP, b, &ret)
	assert.Equal(t, 4, len(ret))
}

func TestWhiteBishopCenterWithCapturesAndWhitePieces(t *testing.T) {
	b, _ := boardFromFen("8/8/8/4r3/5B2/8/3Q4/8 w - - 0 1")
	var row int8 = 6
	var col int8 = 7
	ret := [][]int8{}
	bishopMoves(row, col, WHITE|BISHOP, b, &ret)
	assert.Equal(t, 6, len(ret))
}

func TestWhiteQueenEmptyBoard(t *testing.T) {
	b, _ := boardFromFen("8/8/8/8/3Q4/8/8/8 w - - 0 1")
	var row int8 = 6
	var col int8 = 5
	ret := [][]int8{}
	queenMoves(row, col, WHITE|QUEEN, b, &ret)
	assert.Equal(t, 27, len(ret))
}

func TestWhiteQueenCantMove(t *testing.T) {
	b, _ := boardFromFen("8/8/8/2NBR3/2PQR3/2RRR3/8/8 w - - 0 1")
	var row int8 = 6
	var col int8 = 5
	ret := [][]int8{}
	queenMoves(row, col, WHITE|QUEEN, b, &ret)
	assert.Equal(t, 0, len(ret))
}

func TestWhiteQueenWithOtherPiece(t *testing.T) {
	b, _ := boardFromFen("8/6r1/8/8/3Q4/5N2/8/6P1 w - - 0 1")
	var row int8 = 6
	var col int8 = 5
	ret := [][]int8{}
	queenMoves(row, col, WHITE|QUEEN, b, &ret)
	assert.Equal(t, 25, len(ret))
}

func TestPerftDepthOne(t *testing.T) {
	b, _ := boardFromFen(DEFAULT_POS)
	moves := [][]int8{}
	for i := BOARD_START; i < BOARD_END; i++ {
		for j := BOARD_START; j < BOARD_END; j++ {
			if isWhite(b.board[i][j]) {
				getMoves(int8(i), int8(j), b.board[i][j], b, &moves)
			}
		}
	}
	assert.Equal(t, 20, len(moves))
}

func TestCorrectKingLocation(t *testing.T) {
	b, _ := boardFromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	assert.Equal(t, b.whiteKingLocation[0], 9)
	assert.Equal(t, b.whiteKingLocation[1], 6)
	assert.Equal(t, b.blackKingLocation[0], 2)
	assert.Equal(t, b.blackKingLocation[1], 6)
}

func TestCorrectKingLocationTwo(t *testing.T) {
	b, _ := boardFromFen("6rk/1b4np/5pp1/1p6/8/1P3NP1/1B3P1P/5RK1 w KQkq - 0 1")
	assert.Equal(t, b.whiteKingLocation[0], 9)
	assert.Equal(t, b.whiteKingLocation[1], 8)
	assert.Equal(t, b.blackKingLocation[0], 2)
	assert.Equal(t, b.blackKingLocation[1], 9)
}
