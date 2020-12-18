package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - -
//   T e s t T p o r s
//  - - - - - - - - - -
//
//  Test Tpors function.
//
//  Called:  Tpors, vvd, viv
//
//  This revision:  2017 October 21
//
func TestTpors(t *testing.T) {
	const fname = "Tpors"
	var xi, eta, ra, dec, az1, bz1, az2, bz2 float64
	var n en.ErrNum

	xi = -0.03
	eta = 0.07
	ra = 1.3
	dec = 1.5

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (
			c1, c2, c3, c4 float64, c5 en.ErrNum)
	}{
		{"cgo", CgoTpors},
		{"go", GoTpors},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		az1, bz1, az2, bz2, n = test.fn(xi, eta, ra, dec)

		vvd(t, az1, 1.736621577783208748, 1e-13, tname, "az1")
		vvd(t, bz1, 1.436736561844090323, 1e-13, tname, "bz1")

		vvd(t, az2, 4.004971075806584490, 1e-13, tname, "az2")
		vvd(t, bz2, 1.565084088476417917, 1e-13, tname, "bz2")

		errEN(t, 2, n, tname, "n")
	}
}

func BenchmarkTpors(b *testing.B) {
	var xi, eta, ra, dec float64

	xi = -0.03
	eta = 0.07
	ra = 1.3
	dec = 1.5

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (
			c1, c2, c3, c4 float64, c5 en.ErrNum)
	}{
		{"cgo", CgoTpors},
		{"go", GoTpors},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(xi, eta, ra, dec)
			}
		})
	}
}
