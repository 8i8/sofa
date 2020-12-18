package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - -
//   T e s t T p x e s
//  - - - - - - - - - -
//
//  Test Tpxes function.
//
//  Called:  Tpxes, vvd, viv
//
//  This revision:  2017 October 21
//
func TestTpxes(t *testing.T) {
	const fname = "Tpxes"
	var ra, dec, raz, decz, xi, eta float64
	var err en.ErrNum

	ra = 1.3
	dec = 1.55
	raz = 2.3
	decz = 1.5

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (
			c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoTpxes},
		{"go", GoTpxes},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		xi, eta, err = test.fn(ra, dec, raz, decz)

		vvd(t, xi, -0.01753200983236980595, 1e-15, tname, "xi")
		vvd(t, eta, 0.05962940005778712891, 1e-15, tname, "eta")

		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkTpxes(b *testing.B) {
	var ra, dec, raz, decz float64

	ra = 1.3
	dec = 1.55
	raz = 2.3
	decz = 1.5

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (
			c1, c2 float64, c3 en.ErrNum)
	}{
		{"cgo", CgoTpxes},
		{"go", GoTpxes},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(ra, dec, raz, decz)
			}
		})
	}
}
