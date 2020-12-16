package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t L t e c e q
//  - - - - - - - - - - -
//
//  Test Lteceq function.
//
//  Called:  Lteceq, vvd
//
//  This revision:  2016 March 12
//
func TestLteceq(t *testing.T) {
	const fname = "Lteceq"
	var epj, dl, db, dr, dd float64

	epj = 2500.0
	dl = 1.5
	db = 0.6

	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (c1, c2 float64)
	}{
		{"cgo", CgoLteceq},
		{"go", GoLteceq},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		dr, dd = test.fn(epj, dl, db)

		vvd(t, dr, 1.275156021861921167, 1e-14, tname, "dr")
		vvd(t, dd, 0.9966573543519204791, 1e-14, tname, "dd")
	}
}

func BenchmarkLteceq(b *testing.B) {
	var epj, dl, db float64

	epj = 2500.0
	dl = 1.5
	db = 0.6

	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (c1, c2 float64)
	}{
		{"cgo", CgoLteceq},
		{"go", GoLteceq},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(epj, dl, db)
			}
		})
	}
}
