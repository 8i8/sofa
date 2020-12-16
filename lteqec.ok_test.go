package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t L t e q e c
//  - - - - - - - - - - -
//
//  Test Lteqec function.
//
//  Called:  Lteqec, vvd
//
//  This revision:  2016 March 12
//
func TestLteqec(t *testing.T) {
	const fname = "Lteqec"
	var epj, dr, dd, dl, db float64

	epj = -1500.0
	dr = 1.234
	dd = 0.987

	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (c1, c2 float64)
	}{
		{"cgo", CgoLteqec},
		{"go", GoLteqec},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		dl, db = test.fn(epj, dr, dd)

		vvd(t, dl, 0.5039483649047114859, 1e-14, tname, "dl")
		vvd(t, db, 0.5848534459726224882, 1e-14, tname, "db")
	}
}

func BenchmarkLteqec(b *testing.B) {
	var epj, dr, dd float64

	epj = -1500.0
	dr = 1.234
	dd = 0.987

	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (c1, c2 float64)
	}{
		{"cgo", CgoLteqec},
		{"go", GoLteqec},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(epj, dr, dd)
			}
		})
	}
}
