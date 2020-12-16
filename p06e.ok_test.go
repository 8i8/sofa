package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t P 0 6 e
//  - - - - - - - - -
//
//  Test P06e function.
//
//  Called:  P06e, vvd
//
//  This revision:  2020 May 30
//
func TestP06e(t *testing.T) {
	const fname = "Po6e"
	var eps0, psia, oma, bpa, bqa, pia, bpia,
		epsa, chia, za, zetaa, thetaa, pa, gam, phi, psi float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (
			c1, c2, c3, c4, c5, c6, c7, c8,
			c9, c10, c11, c12, c13, c14, c15, c16 float64)
	}{
		{"cgo", CgoP06e},
		{"go", GoP06e},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		eps0, psia, oma, bpa, bqa, pia, bpia, epsa, chia, za,
			zetaa, thetaa, pa, gam, phi, psi = test.fn(
			2400000.5, 52541.0)

		vvd(t, eps0, 0.4090926006005828715, 1e-14,
			tname, "eps0")
		vvd(t, psia, 0.6664369630191613431e-3, 1e-14,
			tname, "psia")
		vvd(t, oma, 0.4090925973783255982, 1e-14,
			tname, "oma")
		vvd(t, bpa, 0.5561149371265209445e-6, 1e-14,
			tname, "bpa")
		vvd(t, bqa, -0.6191517193290621270e-5, 1e-14,
			tname, "bqa")
		vvd(t, pia, 0.6216441751884382923e-5, 1e-14,
			tname, "pia")
		vvd(t, bpia, 3.052014180023779882, 1e-14,
			tname, "bpia")
		vvd(t, epsa, 0.4090864054922431688, 1e-14,
			tname, "epsa")
		vvd(t, chia, 0.1387703379530915364e-5, 1e-14,
			tname, "chia")
		vvd(t, za, 0.2921789846651790546e-3, 1e-14,
			tname, "za")
		vvd(t, zetaa, 0.3178773290332009310e-3, 1e-14,
			tname, "zetaa")
		vvd(t, thetaa, 0.2650932701657497181e-3, 1e-14,
			tname, "thetaa")
		vvd(t, pa, 0.6651637681381016288e-3, 1e-14,
			tname, "pa")
		vvd(t, gam, 0.1398077115963754987e-5, 1e-14,
			tname, "gam")
		vvd(t, phi, 0.4090864090837462602, 1e-14,
			tname, "phi")
		vvd(t, psi, 0.6664464807480920325e-3, 1e-14,
			tname, "psi")
	}
}

func BenchmarkP06e(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (
			c1, c2, c3, c4, c5, c6, c7, c8,
			c9, c10, c11, c12, c13, c14, c15, c16 float64)
	}{
		{"cgo", CgoP06e},
		{"go", GoP06e},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(2400000.5, 52541.0)
			}
		})
	}
}
