package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t F w 2 m
//  - - - - - - - - -
//
//  Test Fw2m function.
//
//  Called:  Fw2m, vvd
//
//  This revision:  2013 August 7
//
func TestFw2m(t *testing.T) {
	const fname = "Fw2m"
	var gamb, phib, psi, eps float64

	gamb = -0.2243387670997992368e-5
	phib = 0.4091014602391312982
	psi = -0.9501954178013015092e-3
	eps = 0.4091014316587367472

	tests := []struct {
		ref string
		fn  func(a, b, c, d float64) [3][3]float64
	}{
		{"cgo", CgoFw2m},
		{"go", GoFw2m},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref

		r := test.fn(gamb, phib, psi, eps)

		vvd(t, r[0][0], 0.9999995505176007047, 1e-12,
			tname, "11")
		vvd(t, r[0][1], 0.8695404617348192957e-3, 1e-12,
			tname, "12")
		vvd(t, r[0][2], 0.3779735201865582571e-3, 1e-12,
			tname, "13")

		vvd(t, r[1][0], -0.8695404723772016038e-3, 1e-12,
			tname, "21")
		vvd(t, r[1][1], 0.9999996219496027161, 1e-12,
			tname, "22")
		vvd(t, r[1][2], -0.1361752496887100026e-6, 1e-12,
			tname, "23")

		vvd(t, r[2][0], -0.3779734957034082790e-3, 1e-12,
			tname, "31")
		vvd(t, r[2][1], -0.1924880848087615651e-6, 1e-12,
			tname, "32")
		vvd(t, r[2][2], 0.9999999285679971958, 1e-12,
			tname, "33")
	}
}

func BenchmarkFw2m(b *testing.B) {
	var gamb, phib, psi, eps float64

	gamb = -0.2243387670997992368e-5
	phib = 0.4091014602391312982
	psi = -0.9501954178013015092e-3
	eps = 0.4091014316587367472

	tests := []struct {
		ref string
		fn  func(a, b, c, d float64) [3][3]float64
	}{
		{"cgo", CgoFw2m},
		{"go", GoFw2m},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(gamb, phib, psi, eps)
			}
		})
	}
}
