package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - -
//   T e s t T t t a i
//  - - - - - - - - - -
//
//  Test Tttai function.
//
//  Called:  Tttai, vvd, viv
//
//  This revision:  2013 August 7
//
func TestTttai(t *testing.T) {
	const fname = "Tttai"
	var a1, a2 float64
	var err en.ErrNum

	tests := []struct {
		ref string
		fn  func(a1,a2 float64) (c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoTttai},
		{"go", GoTttai},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		a1, a2, err = test.fn(2453750.5, 0.892482639)

		vvd(t, a1, 2453750.5, 1e-6, tname, "a1")
		vvd(t, a2, 0.892110139, 1e-12, tname, "a2")
		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkTttai(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1,a2 float64) (c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoTttai},
		{"go", GoTttai},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(2453750.5, 0.892482639)
			}
		})
	}
}
