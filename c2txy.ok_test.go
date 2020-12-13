package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t C 2 t x y
//  - - - - - - - - - -
//
//  Test C2txy function.
//
//  Called:  C2txy, vvd
//
//  This revision:  2013 August 7
//
func TestC2txy(t *testing.T) {
	const fname = "C2txy"
	var tta, ttb, uta, utb, x, y, xp, yp float64
	var rc2t [3][3]float64

	tta = 2400000.5
	uta = 2400000.5
	ttb = 53736.0
	utb = 53736.0
	x = 0.5791308486706011000e-3
	y = 0.4020579816732961219e-4
	xp = 2.55060238e-7
	yp = 1.860359247e-6

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6, a7,
			a8 float64) [3][3]float64
	}{
		{"cgo", CgoC2txy},
		{"go", GoC2txy},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		rc2t = test.fn(tta, ttb, uta, utb, x, y, xp, yp)

		vvd(t, rc2t[0][0], -0.1810332128306279253, 1e-12,
			tname, "11")
		vvd(t, rc2t[0][1], 0.9834769806938520084, 1e-12,
			tname, "12")
		vvd(t, rc2t[0][2], 0.6555551248057665829e-4, 1e-12,
			tname, "13")

		vvd(t, rc2t[1][0], -0.9834768134136142314, 1e-12,
			tname, "21")
		vvd(t, rc2t[1][1], -0.1810332203649529312, 1e-12,
			tname, "22")
		vvd(t, rc2t[1][2], 0.5749800843594139912e-3, 1e-12,
			tname, "23")

		vvd(t, rc2t[2][0], 0.5773474028619264494e-3, 1e-12,
			tname, "31")
		vvd(t, rc2t[2][1], 0.3961816546911624260e-4, 1e-12,
			tname, "32")
		vvd(t, rc2t[2][2], 0.9999998325501746670, 1e-12,
			tname, "33")
	}
}

func BenchmarkC2txy(b *testing.B) {
	var tta, ttb, uta, utb, x, y, xp, yp float64

	tta = 2400000.5
	uta = 2400000.5
	ttb = 53736.0
	utb = 53736.0
	x = 0.5791308486706011000e-3
	y = 0.4020579816732961219e-4
	xp = 2.55060238e-7
	yp = 1.860359247e-6

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6, a7,
			a8 float64) [3][3]float64
	}{
		{"cgo", CgoC2txy},
		{"go", GoC2txy},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(tta, ttb, uta, utb, x, y,
					xp, yp)
			}
		})
	}
}
