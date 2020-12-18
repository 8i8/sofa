package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - -
//   T e s t T t t c g
//  - - - - - - - - - -
//
//  Test Tttcg function.
//
//  Called:  Tttcg, vvd, viv
//
//  This revision:  2013 August 7
//
func TestTttcg(t *testing.T) {
	const fname = "Tttcg"
	var g1, g2 float64
	var err en.ErrNum

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoTttcg},
		{"go", GoTttcg},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		g1, g2, err = test.fn(2453750.5, 0.892482639)

		vvd(t, g1, 2453750.5, 1e-6, tname, "g1")
		vvd(t, g2, 0.8924900312508587113, 1e-12, tname, "g2")
		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkTttcg(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoTttcg},
		{"go", GoTttcg},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(2453750.5, 0.892482639)
			}
		})
	}
}
