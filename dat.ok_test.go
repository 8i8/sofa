package sofa

import (
	"testing"
	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - -
//   T e s t D a t
//  - - - - - - - -
//
//  Test Dat function.
//
//  Called:  Dat, vvd, viv
//
//  This revision:  2016 July 11
//
func TestDat(t *testing.T) {
	const fname = "Dat"
	tests := []struct {
		ref string
		fn  func(a1, a2, a3 int, a4 float64) (
			b1 float64, b2 en.ErrNum)
	}{
		{"cgo", CgoDat},
		{"go", GoDat},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		deltat, err := test.fn(2003, 6, 1, 0.0)
		vvd(t, deltat, 32.0, 0.0, tname, "d1")
		errT(t, nil, err, tname, "j1")

		deltat, err = test.fn(2008, 1, 17, 0.0)
		vvd(t, deltat, 33.0, 0.0, tname, "d2")
		errT(t, nil, err, tname, "j2")

		deltat, err = test.fn(2017, 9, 1, 0.0)
		vvd(t, deltat, 37.0, 0.0, tname, "d3")
		errT(t, nil, err, tname, "j3")
	}
}

func BenchmarkDat(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2, a3 int, a4 float64) (
			b1 float64, b2 en.ErrNum)
	}{
		{"cgo", CgoDat},
		{"go", GoDat},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(2003, 6, 1, 0.0)
				test.fn(2008, 1, 17, 0.0)
				test.fn(2017, 9, 1, 0.0)
			}
		})
	}
}
