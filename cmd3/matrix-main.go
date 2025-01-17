package main

import (
	"os"
)

type IMatrix interface {
	Get(x int, y int) rune
	Set(x int, y int, r rune)
	Width() int
	Height() int
}
type Matrix struct {
	width, height int
	runes         []rune
}

func NewMatrix(w, h int) IMatrix {
	matrix := Matrix{width: w, height: h, runes: make([]rune, w*h)}
	Fill(matrix, ' ')
	return matrix
}
func (m Matrix) Get(x int, y int) rune {
	return m.runes[x+y*m.width]
}
func (m Matrix) Set(x int, y int, r rune) {
	m.runes[x+y*m.width] = r
}
func (m Matrix) Width() int {
	return m.width
}
func (m Matrix) Height() int {
	return m.height
}

type SubMatrix struct {
	width, height int
	dx, dy        int
	m             IMatrix
}

func NewSubMatrix(width, height int, m IMatrix, dx, dy int) IMatrix {
	sm := SubMatrix{width: width, height: height, m: m, dx: dx, dy: dy}
	return sm
}
func (sm SubMatrix) Get(x int, y int) rune {
	matrix := sm.m
	return matrix.Get(x+sm.dx, y+sm.dy)
}
func (sm SubMatrix) Set(x int, y int, r rune) {
	matrix := sm.m
	matrix.Set(x+sm.dx, y+sm.dy, r)
}
func (sm SubMatrix) Width() int {
	return sm.width
}
func (sm SubMatrix) Height() int {
	return sm.height
}

func Fill(matrix IMatrix, r rune) {
	for y := 0; y < matrix.Height(); y++ {
		for x := 0; x < matrix.Width(); x++ {
			matrix.Set(x, y, r)
		}
	}
}

func WriteTo(matrix IMatrix, f *os.File) {
	for y := 0; y < matrix.Height(); y++ {
		for x := 0; x < matrix.Width(); x++ {
			r := matrix.Get(x, y)
			f.WriteString(string(r))
		}
		f.WriteString("↵\n")
	}
}
func main() {
	m := NewMatrix(20, 10)
	Fill(m, 'a')
	m.Set(0, 0, 'A')
	WriteTo(m, os.Stdout)
	println()

	sm := NewSubMatrix(18, 8, m, 1, 1)
	Fill(sm, 'b')
	sm.Set(0, 0, 'B')
	WriteTo(sm, os.Stdout)
	println()

	ssm := NewSubMatrix(16, 6, sm, 1, 1)
	Fill(ssm, 'c')
	ssm.Set(0, 0, 'C')
	WriteTo(ssm, os.Stdout)
	println()
}
