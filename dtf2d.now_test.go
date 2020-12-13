package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - -
//   T e s t D t f 2 d
//  - - - - - - - - - -
//
//  Test Dtf2d function.
//
//  Called:  Dtf2d, vvd, viv
//
//  This revision:  2013 August 7
//
func TestDtf2d(t *testing.T) {
	const fname = "Dtf2d"

	tests := []struct {
		ref string
		fn  func(a1 string, a2, a3, a4, a5, a6 int, a7 float64) (
			c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoDtf2d},
		{"go", GoDtf2d},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		u1, u2, err := test.fn("UTC", 1994, 6, 30, 23, 59, 60.13599)

		vvd(t, u1+u2, 2449534.49999, 1e-6, tname, "u")
		// viv(t, j, 0, "iauDtf2d", "j")
		errT(t, nil, err, tname)
	}
}

func BenchmarkDtf2d(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1 string, a2, a3, a4, a5, a6 int, a7 float64) (
			c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoDtf2d},
		{"go", GoDtf2d},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn("UTC", 1994, 6, 30, 23, 59, 60.13599)
			}
		})
	}
}
