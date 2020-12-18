package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - -
//   T e s t T p x e v
//  - - - - - - - - - -
//
//  Test Tpxev function.
//
//  Called:  Tpxev, iauS2c, vvd
//
//  This revision:  2017 October 21
//
func TestTpxev(t *testing.T) {
	const fname = "Tpxev"
	var ra, dec, raz, decz, xi, eta float64
	var v, vz [3]float64
	var err en.ErrNum

	ra = 1.3
	dec = 1.55
	raz = 2.3
	decz = 1.5

	tests := []struct {
		ref string
		fn  func(a1, a2 [3]float64) (
			c1, c2 float64, c3 en.ErrNum)
		fnAssist func(a1, a2 float64) [3]float64
	}{
		{"cgo", CgoTpxev, CgoS2c},
		{"go", GoTpxev, GoS2c},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		v = test.fnAssist(ra, dec)
		vz = test.fnAssist(raz, decz)

		xi, eta, err = test.fn(v, vz)

		vvd(t, xi, -0.01753200983236980595, 1e-15, tname, "xi")
		vvd(t, eta, 0.05962940005778712891, 1e-15, tname, "eta")

		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkTpxev(b *testing.B) {
	var ra, dec, raz, decz float64
	var v, vz [3]float64

	ra = 1.3
	dec = 1.55
	raz = 2.3
	decz = 1.5

	tests := []struct {
		ref string
		fn  func(a1, a2 [3]float64) (
			c1, c2 float64, c3 en.ErrNum)
		fnAssist func(a1, a2 float64) [3]float64
	}{
		{"cgo", CgoTpxev, CgoS2c},
		{"go", GoTpxev, GoS2c},
	}

	for _, test := range tests {
		v = test.fnAssist(ra, dec)
		vz = test.fnAssist(raz, decz)
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(v, vz)
			}
		})
	}
}
