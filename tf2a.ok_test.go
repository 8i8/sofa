package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - -
//   T e s t T f 2 a
//  - - - - - - - - -
//
//  Test Tf2a function.
//
//  Called:  Tf2a, vvd, viv
//
//  This revision:  2013 August 7
//
func TestTf2a(t *testing.T) {
	const fname = "Tf2a"
	var a float64
	var err en.ErrNum

	tests := []struct {
		ref string
		fn  func(a1 byte, a2, a3 int, a4 float64) (
			float64, en.ErrNum)
	}{
		{"cgo", CgoTf2a},
		{"go", GoTf2a},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		a, err = test.fn('+', 4, 58, 20.2)

		vvd(t, a, 1.301739278189537429, 1e-12, tname, "a")
		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkTf2a(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1 byte, a2, a3 int, a4 float64) (
			float64, en.ErrNum)
	}{
		{"cgo", CgoTf2a},
		{"go", GoTf2a},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn('+', 4, 58, 20.2)
			}
		})
	}
}
