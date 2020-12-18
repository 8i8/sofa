package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t P r e c 7 6
//  - - - - - - - - - - -
//
//  Test Prec76 function.
//
//  Called:  Prec76, vvd
//
//  This revision:  2013 August 7
//
func TestPrec76(t *testing.T) {
	const fname = "Prec76"
	var ep01, ep02, ep11, ep12, zeta, z, theta float64

	ep01 = 2400000.5
	ep02 = 33282.0
	ep11 = 2400000.5
	ep12 = 51544.0

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (c1, c2, c3 float64)
	}{
		{"cgo", CgoPrec76},
		{"go", GoPrec76},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		zeta, z, theta = test.fn(ep01, ep02, ep11, ep12)

		vvd(t, zeta, 0.5588961642000161243e-2, 1e-12,
			tname, "zeta")
		vvd(t, z, 0.5589922365870680624e-2, 1e-12,
			tname, "z")
		vvd(t, theta, 0.4858945471687296760e-2, 1e-12,
			tname, "theta")
	}
}

func BenchmarkPrec76(b *testing.B) {
	var ep01, ep02, ep11, ep12 float64
	ep01 = 2400000.5
	ep02 = 33282.0
	ep11 = 2400000.5
	ep12 = 51544.0

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (c1, c2, c3 float64)
	}{
		{"cgo", CgoPrec76},
		{"go", GoPrec76},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = test.fn(ep01, ep02, ep11, ep12)
			}
		})
	}
}
