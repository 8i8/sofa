package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - -
//   T e s t T t u t 1
//  - - - - - - - - - -
//
//  Test Ttut1 function.
//
//  Called:  Ttut1, vvd, viv
//
//  This revision:  2013 August 7
//
func TestTtut1(t *testing.T) {
	const fname = "Ttut1"
	var u1, u2 float64
	var err en.ErrNum

	tests := []struct {
		ref string
		fn  func(a1,a2,a3 float64) (c1,c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoTtut1},
		{"go", GoTtut1},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		u1, u2, err = test.fn(2453750.5, 0.892855139, 64.8499)

		vvd(t, u1, 2453750.5, 1e-6, tname, "u1")
		vvd(t, u2, 0.8921045614537037037, 1e-12, tname, "u2")
		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkTtut1(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1,a2,a3 float64) (c1,c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoTtut1},
		{"go", GoTtut1},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(2453750.5, 0.892855139, 64.8499)
			}
		})
	}
}
