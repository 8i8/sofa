package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - - -
//   T e s t J d c a l f
//  - - - - - - - - - - -
//
//  Test Jdcalf function.
//
//  Called:  Jdcalf, viv
//
//  This revision:  2013 August 7
//
func TestJdcalf(t *testing.T) {
	const fname = "Jdcalf"
	var iydmf [4]int
	var err en.ErrNum
	var dj1, dj2 float64

	dj1 = 2400000.5
	dj2 = 50123.9999

	tests := []struct {
		ref string
		fn  func(a1 int, a2, a3 float64) (
			[4]int, en.ErrNum)
	}{
		{"cgo", CgoJdcalf},
		{"go", GoJdcalf},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		iydmf, err = test.fn(4, dj1, dj2)

		viv(t, iydmf[0], 1996, tname, "y")
		viv(t, iydmf[1], 2, tname, "m")
		viv(t, iydmf[2], 10, tname, "d")
		viv(t, iydmf[3], 9999, tname, "f")

		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkJdcalf(b *testing.B) {
	var dj1, dj2 float64

	dj1 = 2400000.5
	dj2 = 50123.9999

	tests := []struct {
		ref string
		fn  func(a1 int, a2, a3 float64) (
			[4]int, en.ErrNum)
	}{
		{"cgo", CgoJdcalf},
		{"go", GoJdcalf},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(4, dj1, dj2)
			}
		})
	}
}
