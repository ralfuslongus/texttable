package texttable

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestFill(t *testing.T) {
	// Wie bekomme ich denn Ausgaben auf stdout im testmodus hin???
	fmt.Println("hallo")
	fmt.Println("hallo")
	fmt.Println("hallo")
	fmt.Println("hallo")

}

func TestNewRuneMatrix(t *testing.T) {
	type args struct {
		w int
		h int
	}
	tests := []struct {
		name string
		args args
		want RuneMatrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRuneMatrix(tt.args.w, tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRuneMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneMatrix_Fill(t *testing.T) {
	type fields struct {
		w     int
		h     int
		runes []rune
	}
	type args struct {
		x      int
		y      int
		width  int
		height int
		r      rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &RuneMatrix{
				w:     tt.fields.w,
				h:     tt.fields.h,
				Runes: tt.fields.runes,
			}
			m.Fill(tt.args.x, tt.args.y, tt.args.width, tt.args.height, tt.args.r)
		})
	}
}

func TestRuneMatrix_SmoothOpenCrossEnds(t *testing.T) {
	type fields struct {
		w     int
		h     int
		runes []rune
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &RuneMatrix{
				w:     tt.fields.w,
				h:     tt.fields.h,
				Runes: tt.fields.runes,
			}
			m.SmoothOpenCrossEnds()
		})
	}
}

func TestRuneMatrix_HasFrameAt(t *testing.T) {
	type fields struct {
		w     int
		h     int
		runes []rune
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &RuneMatrix{
				w:     tt.fields.w,
				h:     tt.fields.h,
				Runes: tt.fields.runes,
			}
			if got := m.HasFrameAt(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("RuneMatrix.HasFrameAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneMatrix_SmoothOpenCrossEnd(t *testing.T) {
	type fields struct {
		w     int
		h     int
		runes []rune
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &RuneMatrix{
				w:     tt.fields.w,
				h:     tt.fields.h,
				Runes: tt.fields.runes,
			}
			if got := m.SmoothOpenCrossEnd(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("RuneMatrix.SmoothOpenCrossEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneMatrix_Set(t *testing.T) {
	type fields struct {
		w     int
		h     int
		runes []rune
	}
	type args struct {
		x int
		y int
		r rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &RuneMatrix{
				w:     tt.fields.w,
				h:     tt.fields.h,
				Runes: tt.fields.runes,
			}
			m.Set(tt.args.x, tt.args.y, tt.args.r)
		})
	}
}

func TestRuneMatrix_Get(t *testing.T) {
	type fields struct {
		w     int
		h     int
		runes []rune
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   rune
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &RuneMatrix{
				w:     tt.fields.w,
				h:     tt.fields.h,
				Runes: tt.fields.runes,
			}
			if got := m.Get(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("RuneMatrix.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneMatrix_HorizontalLineAt(t *testing.T) {
	type fields struct {
		w     int
		h     int
		runes []rune
	}
	type args struct {
		y int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &RuneMatrix{
				w:     tt.fields.w,
				h:     tt.fields.h,
				Runes: tt.fields.runes,
			}
			m.HorizontalLineAt(tt.args.y)
		})
	}
}

func TestRuneMatrix_VerticalLineAt(t *testing.T) {
	type fields struct {
		w     int
		h     int
		runes []rune
	}
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &RuneMatrix{
				w:     tt.fields.w,
				h:     tt.fields.h,
				Runes: tt.fields.runes,
			}
			m.VerticalLineAt(tt.args.x)
		})
	}
}

func TestWriteString(t *testing.T) {
	type args struct {
		m *RuneMatrix
		x int
		y int
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := WriteString(tt.args.m, tt.args.x, tt.args.y, tt.args.s)
			if got != tt.want {
				t.Errorf("WriteString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("WriteString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRuneMatrix_RenderTo(t *testing.T) {
	type fields struct {
		w     int
		h     int
		runes []rune
	}
	type args struct {
		f *os.File
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &RuneMatrix{
				w:     tt.fields.w,
				h:     tt.fields.h,
				Runes: tt.fields.runes,
			}
			m.RenderTo(tt.args.f)
		})
	}
}
