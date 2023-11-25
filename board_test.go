package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPieceRecognized(t *testing.T) {
	assert.True(t, isWhite(WHITE|BISHOP))
	assert.True(t, isWhite(WHITE|ROOK))
	assert.True(t, isWhite(WHITE|KING))
	assert.True(t, isWhite(WHITE|PAWN))

	assert.True(t, isBlack(BLACK|BISHOP))
	assert.True(t, isBlack(BLACK|ROOK))
	assert.True(t, isBlack(BLACK|KING))
	assert.True(t, isBlack(BLACK|PAWN))

	assert.True(t, isPawn(WHITE|PAWN))
	assert.True(t, isPawn(BLACK|PAWN))
	assert.True(t, !isPawn(BLACK|ROOK))

	assert.True(t, isKnight(WHITE|KNIGHT))
	assert.True(t, isKnight(BLACK|KNIGHT))
	assert.True(t, !isKnight(WHITE|QUEEN))

	assert.True(t, isBishop(WHITE|BISHOP))
	assert.True(t, isBishop(BLACK|BISHOP))
	assert.True(t, !isBishop(BLACK|ROOK))

	assert.True(t, isQueen(WHITE|QUEEN))
	assert.True(t, isQueen(BLACK|QUEEN))
	assert.True(t, !isQueen(BLACK|ROOK))

	assert.True(t, isKing(WHITE|KING))
	assert.True(t, isKing(BLACK|KING))
	assert.True(t, !isKing(BLACK|QUEEN))

	assert.True(t, isEmpty(EMPTY))
	assert.True(t, !isEmpty(WHITE|KING))

	assert.True(t, isOutsideBoard(SENTINEL))
	assert.True(t, !isOutsideBoard(EMPTY))
	assert.True(t, !isOutsideBoard(WHITE|KING))

}
