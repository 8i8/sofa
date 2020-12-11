package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t C 2 t 0 6 a
//  - - - - - - - - - - -
//
//  Test C2t06a function.
//
//  Called:  C2t06a, vvd
//
//  This revision:  2013 August 7
//
func TestC2t06a(t *testing.T) {
	const fname = "C2t06a"
	var tta, ttb, uta, utb, xp, yp float64
	var rc2t [3][3]float64

	tta = 2400000.5
	uta = 2400000.5
	ttb = 53736.0
	utb = 53736.0
	xp = 2.55060238e-7
	yp = 1.860359247e-6

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6 float64) [3][3]float64
	}{
		{"cgo", CgoC2t06a},
		{"go", GoC2t06a},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rc2t = test.fn(tta, ttb, uta, utb, xp, yp)

		vvd(t, rc2t[0][0], -0.1810332128305897282, 1e-12,
			tname, "11")
		vvd(t, rc2t[0][1], 0.9834769806938592296, 1e-12,
			tname, "12")
		vvd(t, rc2t[0][2], 0.6555550962998436505e-4, 1e-12,
			tname, "13")

		vvd(t, rc2t[1][0], -0.9834768134136214897, 1e-12,
			tname, "21")
		vvd(t, rc2t[1][1], -0.1810332203649130832, 1e-12,
			tname, "22")
		vvd(t, rc2t[1][2], 0.5749800844905594110e-3, 1e-12,
			tname, "23")

		vvd(t, rc2t[2][0], 0.5773474024748545878e-3, 1e-12,
			tname, "31")
		vvd(t, rc2t[2][1], 0.3961816829632690581e-4, 1e-12,
			tname, "32")
		vvd(t, rc2t[2][2], 0.9999998325501747785, 1e-12,
			tname, "33")
	}
}

func BenchmarkC2t06a(b *testing.B) {
	var tta, ttb, uta, utb, xp, yp float64

	tta = 2400000.5
	uta = 2400000.5
	ttb = 53736.0
	utb = 53736.0
	xp = 2.55060238e-7
	yp = 1.860359247e-6

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6 float64) [3][3]float64
	}{
		{"cgo", CgoC2t06a},
		{"go", GoC2t06a},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(tta, ttb, uta, utb, xp, yp)
			}
		})
	}
}
