package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - - -
//   T e s t T a i u t c
//  - - - - - - - - - - -
//
//  Test Taiutc function.
//
//  Called:  Taiutc, vvd, viv
//
//  This revision:  2013 October 3
//
func TestTaiutc(t *testing.T) {
	const fname = "Taiutc"
	var u1, u2 float64
	var err en.ErrNum

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoTaiutc},
		{"go", GoTaiutc},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		u1, u2, err = test.fn(2453750.5, 0.892482639)

		vvd(t, u1, 2453750.5, 1e-6, tname, "u1")
		vvd(t, u2, 0.8921006945555555556, 1e-12, tname, "u2")
		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkTaiutc(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoTaiutc},
		{"go", GoTaiutc},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = test.fn(2453750.5, 0.892482639)
			}
		})
	}
}
