package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t P v 2 s
//  - - - - - - - - -
//
//  Test Pv2s function.
//
//  Called:  Pv2s, vvd
//
//  This revision:  2013 August 7
//
func TestPv2s(t *testing.T) {
	const fname = "Pv2s"
	var theta, phi, r, td, pd, rd float64
	var pv [2][3]float64

	pv[0][0] = -0.4514964673880165
	pv[0][1] = 0.03093394277342585
	pv[0][2] = 0.05594668105108779

	pv[1][0] = 1.292270850663260e-5
	pv[1][1] = 2.652814182060692e-6
	pv[1][2] = 2.568431853930293e-6

	tests := []struct {
		ref string
		fn  func([2][3]float64) (c1, c2, c3, c4, c5, c6 float64)
	}{
		{"cgo", CgoPv2s},
		{"go", GoPv2s},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		theta, phi, r, td, pd, rd = test.fn(pv)

		vvd(t, theta, 3.073185307179586515, 1e-12, tname, "theta")
		vvd(t, phi, 0.1229999999999999992, 1e-12, tname, "phi")
		vvd(t, r, 0.4559999999999999757, 1e-12, tname, "r")
		vvd(t, td, -0.7800000000000000364e-5, 1e-16, tname, "td")
		vvd(t, pd, 0.9010000000000001639e-5, 1e-16, tname, "pd")
		vvd(t, rd, -0.1229999999999999832e-4, 1e-16, tname, "rd")
	}
}

func BenchmarkPv2s(b *testing.B) {
	var pv [2][3]float64

	pv[0][0] = -0.4514964673880165
	pv[0][1] = 0.03093394277342585
	pv[0][2] = 0.05594668105108779

	pv[1][0] = 1.292270850663260e-5
	pv[1][1] = 2.652814182060692e-6
	pv[1][2] = 2.568431853930293e-6

	tests := []struct {
		ref string
		fn  func([2][3]float64) (c1, c2, c3, c4, c5, c6 float64)
	}{
		{"cgo", CgoPv2s},
		{"go", GoPv2s},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(pv)
			}
		})
	}
}
