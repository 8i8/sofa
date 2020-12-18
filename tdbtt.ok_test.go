package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - -
//   T e a t T d b t t
//  - - - - - - - - - -
//
//  Test Tdbtt function.
//
//  Called:  Tdbtt, vvd, viv
//
//  This revision:  2013 August 7
//
func TestTdbtt(t *testing.T) {
	const fname = "Tdbtt"
	var t1, t2 float64
	var err en.ErrNum

	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (
			c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoTdbtt},
		{"go", GoTdbtt},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		t1, t2, err = test.fn(2453750.5, 0.892855137, -0.000201)

		vvd(t, t1, 2453750.5, 1e-6, tname, "t1")
		vvd(t, t2, 0.8928551393263888889, 1e-12, tname, "t2")
		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkTdbtt(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (
			c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoTdbtt},
		{"go", GoTdbtt},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(2453750.5, 0.892855137,
				-0.000201)
			}
		})
	}
}
