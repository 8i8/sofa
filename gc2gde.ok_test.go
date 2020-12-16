package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - - -
//   T e s t G c 2 g d e
//  - - - - - - - - - - -
//
//  Test Gc2gde function.
//
//  Called:  Gc2gde, viv, vvd
//
//  This revision:  2016 March 12
//
func TestGc2gde(t *testing.T) {
	const fname = "Gc2gde"
	var e, p, h float64
	var err en.ErrNum
	var a, f = 6378136.0, 0.0033528
	var xyz = [...]float64{2e6, 3e6, 5.244e6}

	tests := []struct {
		ref string
		fn  func(a1, a2 float64, a3 [3]float64) (
			c1, c2, c3 float64, c4 en.ErrNum)
	}{
		{"cgo", CgoGc2gde},
		{"go", GoGc2gde},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		e, p, h, err = test.fn(a, f, xyz)

		errT(t, nil, err, tname, "0")
		vvd(t, e, 0.9827937232473290680, 1e-14, tname, "e")
		vvd(t, p, 0.9716018377570411532, 1e-14, tname, "p")
		vvd(t, h, 332.36862495764397, 1e-8, tname, "h")
	}
}

func BenchmarkGc2gde(b *testing.B) {
	var a, f = 6378136.0, 0.0033528
	var xyz = [...]float64{2e6, 3e6, 5.244e6}

	tests := []struct {
		ref string
		fn  func(a1, a2 float64, a3 [3]float64) (
			c1, c2, c3 float64, c4 en.ErrNum)
	}{
		{"cgo", CgoGc2gde},
		{"go", GoGc2gde},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _, _ = test.fn(a, f, xyz)
			}
		})
	}
}
