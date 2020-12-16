package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t G 2 i c r s
//  - - - - - - - - - - -
//
//  Test G2icrs function.
//
//  Called:  G2icrs, vvd
//
//  This revision:  2015 January 30
//
func TestG2icrs(t *testing.T) {
	const fname = "G2icrs"
	var dl, db, dr, dd float64

	dl = 5.5850536063818546461558105
	db = -0.7853981633974483096156608

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2 float64)
	}{
		{"cgo", CgoG2icrs},
		{"go", GoG2icrs},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		dr, dd = test.fn(dl, db)

		vvd(t, dr, 5.9338074302227188048671, 1e-14, tname, "R")
		vvd(t, dd, -1.1784870613579944551541, 1e-14, tname, "D")
	}
}

func BenchmarkG2icrs(b *testing.B) {
	var dl, db float64

	dl = 5.5850536063818546461558105
	db = -0.7853981633974483096156608

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2 float64)
	}{
		{"cgo", CgoG2icrs},
		{"go", GoG2icrs},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(dl, db)
			}
		})
	}
}
