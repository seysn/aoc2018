package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountOverlapped(t *testing.T) {
	var board Board

	board.PutRectangle(ParseLine("#1 @ 1,3: 4x4"))
	board.PutRectangle(ParseLine("#2 @ 3,1: 4x4"))
	board.PutRectangle(ParseLine("#3 @ 5,5: 2x2"))

	assert.Equal(t, 4, board.CountOverlapped())
}

func TestParseLine(t *testing.T) {
	line := "#123 @ 3,2: 5x4"

	r := ParseLine(line)
	assert.Equal(t, 123, r.id)
	assert.Equal(t, 3, r.x)
	assert.Equal(t, 2, r.y)
	assert.Equal(t, 5, r.w)
	assert.Equal(t, 4, r.h)
}
