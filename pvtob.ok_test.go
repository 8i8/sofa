package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t P v t o b
//  - - - - - - - - - -
//
//  Test Pvtob function.
//
//  Called:  Pvtob, vvd
//
//  This revision:  2013 October 2
//
func TestPvtob(t *testing.T) {
	const fname = "Pvtob"
	var elong, phi, hm, xp, yp, sp, theta float64

	elong = 2.0
	phi = 0.5
	hm = 3000.0
	xp = 1e-6
	yp = -0.5e-6
	sp = 1e-8
	theta = 5.0

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4,
			a5, a6, a7 float64) [2][3]float64
	}{
		{"cgo", CgoPvtob},
		{"go", GoPvtob},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		pv := test.fn(elong, phi, hm, xp, yp, sp, theta)

		vvd(t, pv[0][0], 4225081.367071159207, 1e-5,
			tname, "p(1)")
		vvd(t, pv[0][1], 3681943.215856198144, 1e-5,
			tname, "p(2)")
		vvd(t, pv[0][2], 3041149.399241260785, 1e-5,
			tname, "p(3)")
		vvd(t, pv[1][0], -268.4915389365998787, 1e-9,
			tname, "v(1)")
		vvd(t, pv[1][1], 308.0977983288903123, 1e-9,
			tname, "v(2)")
	}
}

func BenchmarkPvtob(b *testing.B) {
	var elong, phi, hm, xp, yp, sp, theta float64

	elong = 2.0
	phi = 0.5
	hm = 3000.0
	xp = 1e-6
	yp = -0.5e-6
	sp = 1e-8
	theta = 5.0

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4,
			a5, a6, a7 float64) [2][3]float64
	}{
		{"cgo", CgoPvtob},
		{"go", GoPvtob},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(elong, phi, hm, xp,
					yp, sp, theta)
			}
		})
	}
}
