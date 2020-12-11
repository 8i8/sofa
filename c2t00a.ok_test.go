package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t C 2 t 0 0 a
//  - - - - - - - - - - -
//
//  Test C2t00a function.
//
//  Called:  C2t00a, vvd
//
//  This revision:  2013 August 7
//
func TestC2t00a(t *testing.T) {
	const fname = "C2t00a"
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
		{"cgo", CgoC2t00a},
		{"go", GoC2t00a},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rc2t = test.fn(tta, ttb, uta, utb, xp, yp)

		vvd(t, rc2t[0][0], -0.1810332128307182668, 1e-12,
			tname, "11")
		vvd(t, rc2t[0][1], 0.9834769806938457836, 1e-12,
			tname, "12")
		vvd(t, rc2t[0][2], 0.6555535638688341725e-4, 1e-12,
			tname, "13")

		vvd(t, rc2t[1][0], -0.9834768134135984552, 1e-12,
			tname, "21")
		vvd(t, rc2t[1][1], -0.1810332203649520727, 1e-12,
			tname, "22")
		vvd(t, rc2t[1][2], 0.5749801116141056317e-3, 1e-12,
			tname, "23")

		vvd(t, rc2t[2][0], 0.5773474014081406921e-3, 1e-12,
			tname, "31")
		vvd(t, rc2t[2][1], 0.3961832391770163647e-4, 1e-12,
			tname, "32")
		vvd(t, rc2t[2][2], 0.9999998325501692289, 1e-12,
			tname, "33")
	}
}

func BenchmarkC2t00a(b *testing.B) {
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
		{"cgo", CgoC2t00a},
		{"go", GoC2t00a},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(tta, ttb, uta, utb, xp, yp)
			}
		})
	}
}
