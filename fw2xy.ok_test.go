package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t F w 2 x y
//  - - - - - - - - - -
//
//  Test Fw2xy function.
//
//  Called:  Fw2xy, vvd
//
//  This revision:  2013 August 7
//
func TestFw2xy(t *testing.T) {
	const fname = "Fw2xy"
	var gamb, phib, psi, eps, x, y float64

	gamb = -0.2243387670997992368e-5
	phib = 0.4091014602391312982
	psi = -0.9501954178013015092e-3
	eps = 0.4091014316587367472

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (c1, c2 float64)
	}{
		{"cgo", CgoFw2xy},
		{"go", GoFw2xy},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		x, y = test.fn(gamb, phib, psi, eps)

		vvd(t, x, -0.3779734957034082790e-3, 1e-14, tname, "x")
		vvd(t, y, -0.1924880848087615651e-6, 1e-14, tname, "y")
	}
}

func BenchmarkFw2xy(b *testing.B) {
	var gamb, phib, psi, eps float64

	gamb = -0.2243387670997992368e-5
	phib = 0.4091014602391312982
	psi = -0.9501954178013015092e-3
	eps = 0.4091014316587367472

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (c1, c2 float64)
	}{
		{"cgo", CgoFw2xy},
		{"go", GoFw2xy},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(gamb, phib, psi, eps)
			}
		})
	}
}
