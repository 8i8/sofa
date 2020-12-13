package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t C 2 t p e
//  - - - - - - - - - -
//
//  Test C2tpe function.
//
//  Called:  C2tpe, vvd
//
//  This revision:  2013 August 7
//
func TestC2tpe(t *testing.T) {
	const fname = "C2tpe"
	var tta, ttb, uta, utb, dpsi, deps, xp, yp float64
	var rc2t [3][3]float64

	tta = 2400000.5
	uta = 2400000.5
	ttb = 53736.0
	utb = 53736.0
	deps = 0.4090789763356509900
	dpsi = -0.9630909107115582393e-5
	xp = 2.55060238e-7
	yp = 1.860359247e-6

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6, a7,
			a8 float64) [3][3]float64
	}{
		{"cgo", CgoC2tpe},
		{"go", GoC2tpe},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rc2t = test.fn(tta, ttb, uta, utb, dpsi, deps, xp, yp)

		vvd(t, rc2t[0][0], -0.1813677995763029394, 1e-12,
			tname, "11")
		vvd(t, rc2t[0][1], 0.9023482206891683275, 1e-12,
			tname, "12")
		vvd(t, rc2t[0][2], -0.3909902938641085751, 1e-12,
			tname, "13")

		vvd(t, rc2t[1][0], -0.9834147641476804807, 1e-12,
			tname, "21")
		vvd(t, rc2t[1][1], -0.1659883635434995121, 1e-12,
			tname, "22")
		vvd(t, rc2t[1][2], 0.7309763898042819705e-1, 1e-12,
			tname, "23")

		vvd(t, rc2t[2][0], 0.1059685430673215247e-2, 1e-12,
			tname, "31")
		vvd(t, rc2t[2][1], 0.3977631855605078674, 1e-12,
			tname, "32")
		vvd(t, rc2t[2][2], 0.9174875068792735362, 1e-12,
			tname, "33")
	}
}

func BenchmarkC2tpe(b *testing.B) {
	var tta, ttb, uta, utb, dpsi, deps, xp, yp float64

	tta = 2400000.5
	uta = 2400000.5
	ttb = 53736.0
	utb = 53736.0
	deps = 0.4090789763356509900
	dpsi = -0.9630909107115582393e-5
	xp = 2.55060238e-7
	yp = 1.860359247e-6

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6, a7,
			a8 float64) [3][3]float64
	}{
		{"cgo", CgoC2tpe},
		{"go", GoC2tpe},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(tta, ttb, uta, utb, dpsi,
					deps, xp, yp)
			}
		})
	}
}
