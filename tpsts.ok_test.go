package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t T p s t s
//  - - - - - - - - - -
//
//  Test Tpsts function.
//
//  Called:  Tpsts, vvd
//
//  This revision:  2017 October 21
//
func TestTpsts(t *testing.T) {
	const fname = "Tpsts"
	var xi, eta, raz, decz, ra, dec float64

	xi = -0.03
	eta = 0.07
	raz = 2.3
	decz = 1.5

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (c1, c2 float64)
	}{
		{"cgo", CgoTpsts},
		{"go", GoTpsts},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		ra, dec = test.fn(xi, eta, raz, decz)

		vvd(t, ra, 0.7596127167359629775, 1e-14, tname, "ra")
		vvd(t, dec, 1.540864645109263028, 1e-13, tname, "dec")
	}
}

func BenchmarkTpsts(b *testing.B) {
	var xi, eta, raz, decz float64

	xi = -0.03
	eta = 0.07
	raz = 2.3
	decz = 1.5

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (c1, c2 float64)
	}{
		{"cgo", CgoTpsts},
		{"go", GoTpsts},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(xi, eta, raz, decz)
			}
		})
	}
}
