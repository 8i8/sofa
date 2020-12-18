package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t T p s t v
//  - - - - - - - - - -
//
//  Test Tpstv function.
//
//  Called:  Tpstv, iauS2c, vvd
//
//  This revision:  2017 October 21
//
func TestTpstv(t *testing.T) {
	const fname = "Tpstv"
	var xi, eta, raz, decz float64
	var vz, v [3]float64

	xi = -0.03
	eta = 0.07
	raz = 2.3
	decz = 1.5

	tests := []struct {
		ref      string
		fn       func(a1, a2 float64, a3 [3]float64) [3]float64
		fnAssist func(a1, a2 float64) [3]float64
	}{
		{"cgo", CgoTpstv, CgoS2c},
		{"go", GoTpstv, GoS2c},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		vz = test.fnAssist(raz, decz)

		v = test.fn(xi, eta, vz)

		vvd(t, v[0], 0.02170030454907376677, 1e-15, tname, "x")
		vvd(t, v[1], 0.02060909590535367447, 1e-15, tname, "y")
		vvd(t, v[2], 0.9995520806583523804, 1e-14, tname, "z")
	}
}

func BenchmarkTpstv(b *testing.B) {
	var xi, eta, raz, decz float64
	var vz [3]float64

	xi = -0.03
	eta = 0.07
	raz = 2.3
	decz = 1.5

	tests := []struct {
		ref      string
		fn       func(a1, a2 float64, a3 [3]float64) [3]float64
		fnAssist func(a1, a2 float64) [3]float64
	}{
		{"cgo", CgoTpstv, CgoS2c},
		{"go", GoTpstv, GoS2c},
	}

	for _, test := range tests {
		vz = test.fnAssist(raz, decz)
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(xi, eta, vz)
			}
		})
	}
}
