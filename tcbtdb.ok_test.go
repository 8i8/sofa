package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - - -
//   T e s t T c b t d b
//  - - - - - - - - - - -
//
//  Test Tcbtdb function.
//
//  Called:  Tcbtdb, vvd, viv
//
//  This revision:  2013 August 7
//
func TestTcbtdb(t *testing.T) {
	const fname = "Tcbtdb"
	var b1, b2 float64
	var err en.ErrNum

	tests := []struct {
		ref string
		fn  func(a1,a2 float64) (c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoTcbtdb},
		{"go", GoTcbtdb},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		b1, b2, err = test.fn(2453750.5, 0.893019599)

		vvd(t, b1, 2453750.5, 1e-6, tname, "b1")
		vvd(t, b2, 0.8928551362746343397, 1e-12, tname, "b2")
		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkTcbtdb(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1,a2 float64) (c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoTcbtdb},
		{"go", GoTcbtdb},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(2453750.5, 0.893019599)
			}
		})
	}
}
