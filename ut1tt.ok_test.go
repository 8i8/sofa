package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - -
//   T e s t U t 1 t t
//  - - - - - - - - - -
//
//  Test Ut1tt function.
//
//  Called:  Ut1tt, vvd, viv
//
//  This revision:  2013 October 3
//
func TestUt1tt(t *testing.T) {
	const fname = "Ut1tt"
	var t1, t2 float64
	var err en.ErrNum

	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (
			c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoUt1tt},
		{"go", GoUt1tt},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		t1, t2, err = test.fn(2453750.5, 0.892104561, 64.8499)

		vvd(t, t1, 2453750.5, 1e-6, tname, "t1")
		vvd(t, t2, 0.8928551385462962963, 1e-12, tname, "t2")
		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkUt1tt(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (
			c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoUt1tt},
		{"go", GoUt1tt},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(2453750.5, 0.892104561, 64.8499)
			}
		})
	}
}
