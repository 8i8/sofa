package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - -
//   T e s t T c g t t
//  - - - - - - - - - -
//
//  Test Tcgtt function.
//
//  Called:  Tcgtt, vvd, viv
//
//  This revision:  2013 August 7
//
func TestTcgtt(t *testing.T) {
	const fname = "Tcgtt"
	var t1, t2 float64
	var err en.ErrNum

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoTcgtt},
		{"go", GoTcgtt},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		t1, t2, err = test.fn(2453750.5, 0.892862531)

		vvd(t, t1, 2453750.5, 1e-6, tname, "t1")
		vvd(t, t2, 0.8928551387488816828, 1e-12, tname, "t2")
		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkTcgtt(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoTcgtt},
		{"go", GoTcgtt},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(2453750.5, 0.892862531)
			}
		})
	}
}
