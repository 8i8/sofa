package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - - -
//   T e s t U t 1 t a i
//  - - - - - - - - - - -
//
//  Test Ut1tai function.
//
//  Called:  Ut1tai, vvd, viv
//
//  This revision:  2013 August 7
//
func TestUt1tai(t *testing.T) {
	const fname = "Ut1tai"
	var a1, a2 float64
	var err en.ErrNum

	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoUt1tai},
		{"go", GoUt1tai},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		a1, a2, err = test.fn(2453750.5, 0.892104561, -32.6659)

		vvd(t, a1, 2453750.5, 1e-6, tname, "a1")
		vvd(t, a2, 0.8924826385462962963, 1e-12, tname, "a2")
		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkUt1tai(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoUt1tai},
		{"go", GoUt1tai},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(2453750.5, 0.892104561, -32.6659)
			}
		})
	}
}
