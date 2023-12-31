package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func getMoves(row int8, col int8, piece uint8, board *Board, moves *[][]int8) {
	targetPiece := piece & PIECE_MASK
	if targetPiece == PAWN {
		pawnMoves(row, col, piece, board, moves)
	} else if targetPiece == ROOK {
		rookMoves(row, col, piece, board, moves)
	} else if targetPiece == BISHOP {
		bishopMoves(row, col, piece, board, moves)
	} else if targetPiece == KNIGHT {
		knightMoves(row, col, piece, board, moves)
	} else if targetPiece == KING {
		kingMoves(row, col, piece, board, moves)
	} else if targetPiece == QUEEN {
		queenMoves(row, col, piece, board, moves)
	} else {
		panic("Unrecognized piece")
	}
}

func queenMoves(row int8, col int8, piece uint8, board *Board, moves *[][]int8) {
	rookMoves(row, col, piece, board, moves)
	bishopMoves(row, col, piece, board, moves)
}

func bishopMoves(row int8, col int8, piece uint8, board *Board, moves *[][]int8) {
	mods := [2]int8{1, -1}
	for _, i := range mods {
		for _, j := range mods {
			var multiplier int8 = 1
			_row := row + i
			_col := col + j
			square := board.board[_row][_col]
			for isEmpty(square) {
				*moves = append(*moves, []int8{_row, _col})
				multiplier++
				_row = row + (i * multiplier)
				_col = col + (j * multiplier)
				square = board.board[_row][_col]
			}

			if !isOutsideBoard(square) && piece&COLOR_MASK != square&COLOR_MASK {
				*moves = append(*moves, []int8{_row, _col})
			}

		}
	}
}

func rookMoves(row int8, col int8, piece uint8, board *Board, moves *[][]int8) {
	mods := [][]int8{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, m := range mods {
		multiplier := 1
		_row := row + m[0]
		_col := col + m[1]
		square := board.board[_row][_col]
		for isEmpty(square) {
			*moves = append(*moves, []int8{_row, _col})
			multiplier++
			_row = row + (m[0] * int8(multiplier))
			_col = col + (m[1] * int8(multiplier))
			square = board.board[_row][_col]
		}

		if !isOutsideBoard(square) && piece&COLOR_MASK != square&COLOR_MASK {
			*moves = append(*moves, []int8{_row, _col})
		}
	}
}

func kingMoves(row int8, col int8, piece uint8, board *Board, moves *[][]int8) {
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			_row := row + int8(i)
			_col := col + int8(j)

			if isOutsideBoard(board.board[_row][_col]) {
				continue
			}

			if isEmpty(board.board[_row][_col]) || board.board[_row][_col]&COLOR_MASK != piece&COLOR_MASK {
				*moves = append(*moves, []int8{_row, _col})
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
		square := board.board[_row][_col]
		if isOutsideBoard(square) {
			continue
		}
		if isEmpty(square) || (square&COLOR_MASK) != piece&COLOR_MASK {
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

	whiteKingLocation := [2]int{0, 0}
	blackKingLocation := [2]int{0, 0}

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

				if isKing(b[row][col]) {
					if isWhite(b[row][col]) {
						whiteKingLocation = [2]int{row, col}
					} else {
						blackKingLocation = [2]int{row, col}
					}
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
		whiteKingLocation: whiteKingLocation, blackKingLocation: blackKingLocation,
	}, nil
}
