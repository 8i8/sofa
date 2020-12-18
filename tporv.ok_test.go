package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - -
//   T e s t T p o r v
//  - - - - - - - - - -
//
//  Test Tporv function.
//
//  Called:  Tporv, iauS2c, vvd, viv
//
//  This revision:  2017 October 21
//
func TestTporv(t *testing.T) {
	const fname = "Tporv"
	var xi, eta, ra, dec float64
	var v, vz1, vz2 [3]float64
	var n en.ErrNum

	xi = -0.03
	eta = 0.07
	ra = 1.3
	dec = 1.5

	tests := []struct {
		ref string
		fn  func(a1, a2 float64, a3 [3]float64) (
			c1, c2 [3]float64, c4 en.ErrNum)
		fnAssist func(a1,a2 float64) [3]float64
	}{
		{"cgo", CgoTporv, CgoS2c},
		{"go", GoTporv, GoS2c},
	}

	for _, test := range tests {

		v = test.fnAssist(ra, dec)

		tname := fname + " " + test.ref

		vz1, vz2, n = test.fn(xi, eta, v)

		vvd(t, vz1[0], -0.02206252822366888610, 1e-15,
			tname, "x1")
		vvd(t, vz1[1], 0.1318251060359645016, 1e-14,
			tname, "y1")
		vvd(t, vz1[2], 0.9910274397144543895, 1e-14,
			tname, "z1")

		vvd(t, vz2[0], -0.003712211763801968173, 1e-16,
			tname, "x2")
		vvd(t, vz2[1], -0.004341519956299836813, 1e-16,
			tname, "y2")
		vvd(t, vz2[2], 0.9999836852110587012, 1e-14,
			tname, "z2")

		errEN(t, 2, n, tname, "n")
	}
}

func BenchmarkTporv(b *testing.B) {
	var xi, eta, ra, dec float64
	var v [3]float64

	xi = -0.03
	eta = 0.07
	ra = 1.3
	dec = 1.5

	v = GoS2c(ra, dec)

	tests := []struct {
		ref string
		fn  func(a1, a2 float64, a3 [3]float64) (
			c1, c2 [3]float64, c4 en.ErrNum)
	}{
		{"cgo", CgoTporv},
		{"go", GoTporv},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(xi, eta, v)
			}
		})
	}
}
