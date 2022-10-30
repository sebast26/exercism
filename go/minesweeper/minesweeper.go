package minesweeper

import (
	"fmt"
	"math"
	"strings"
)

type Board struct {
	r, c  int
	mines positions
	hints positions
}

func NewBoard(r, c int) Board {
	return Board{r: r, c: c}
}

func (b *Board) Mine(r, c int) {
	b.mines = append(b.mines, bPos{r: r, c: c})
	h := validBoardPos(b.r, b.c, r, c)
	b.hints = append(b.hints, h...)
}

func (b *Board) Annotate() []string {
	var out []string
	minesMap := b.mines.toMap()
	hintMap := b.hints.toMap()
	for ri := 0; ri < b.r; ri++ {
		var sb strings.Builder
		for ci := 0; ci < b.c; ci++ {
			_, ok := minesMap[bPos{r: ri, c: ci}]
			if ok {
				sb.WriteString("*")
				continue
			}

			v, ok := hintMap[bPos{r: ri, c: ci}]
			if !ok {
				sb.WriteString(" ")
			} else {
				sb.WriteString(fmt.Sprintf("%d", v))
			}
		}
		out = append(out, sb.String())
	}
	return out
}

type positions []bPos

type bPos struct {
	r, c int
}

func validBoardPos(bR, bC, r, c int) []bPos {
	minR := int(math.Max(0, float64(r-1)))
	minC := int(math.Max(0, float64(c-1)))
	maxR := int(math.Min(float64(bR), float64(r+1)))
	maxC := int(math.Min(float64(bC), float64(c+1)))
	var p positions
	for ri := minR; ri <= maxR; ri++ {
		for ci := minC; ci <= maxC; ci++ {
			if ri == r && ci == c {
				continue
			}
			p = append(p, bPos{r: ri, c: ci})
		}
	}
	return p
}

func (p positions) toMap() map[bPos]int {
	var m = map[bPos]int{}
	for _, hint := range p {
		_, ok := m[hint]
		if !ok {
			m[hint] = 0
		}
		m[hint]++
	}
	return m
}

// Annotate returns an annotated board
func Annotate(board []string) []string {
	if len(board) == 0 {
		return board
	}
	for _, l := range board {
		if l == "" {
			return board
		}
	}

	b := NewBoard(len(board), len(board[0]))
	for r := range board {
		for c, v := range board[r] {
			if v == '*' {
				b.Mine(r, c)
			}
		}
	}

	return b.Annotate()
}
