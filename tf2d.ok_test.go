package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - -
//   T e s t T f 2 d
//  - - - - - - - - -
//
//  Test Tf2d function.
//
//  Called:  Tf2d, vvd, viv
//
//  This revision:  2013 August 7
//
func TestTf2d(t *testing.T) {
	const fname = "Tf2d"
	var d float64
	var err en.ErrNum

	tests := []struct {
		ref string
		fn  func(a1 byte, a2, a3 int, a4 float64) (
			float64, en.ErrNum)
	}{
		{"cgo", CgoTf2d},
		{"go", GoTf2d},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		d, err = test.fn(' ', 23, 55, 10.9)

		vvd(t, d, 0.9966539351851851852, 1e-12, tname, "d")
		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkTf2d(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1 byte, a2, a3 int, a4 float64) (
			float64, en.ErrNum)
	}{
		{"cgo", CgoTf2d},
		{"go", GoTf2d},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(' ', 23, 55, 10.9)
			}
		})
	}
}
