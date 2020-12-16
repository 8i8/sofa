package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t E q e c 0 6
//  - - - - - - - - - - -
//
//  Test Eqec06 function.
//
//  Called:  Eqec06, vvd
//
//  This revision:  2016 March 12
//
func TestEqec06(t *testing.T) {
	const fname = "Eqec06"
	var date1, date2, dr, dd, dl, db float64

	date1 = 1234.5
	date2 = 2440000.5
	dr = 1.234
	dd = 0.987

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (c1, c2 float64)
	}{
		{"cgo", CgoEqec06},
		{"go", GoEqec06},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		dl, db = test.fn(date1, date2, dr, dd)

		vvd(t, dl, 1.342509918994654619, 1e-14, tname, "dl")
		vvd(t, db, 0.5926215259704608132, 1e-14, tname, "db")
	}
}

func BenchmarkEqec06(b *testing.B) {
	var date1, date2, dr, dd float64

	date1 = 1234.5
	date2 = 2440000.5
	dr = 1.234
	dd = 0.987

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (c1, c2 float64)
	}{
		{"cgo", CgoEqec06},
		{"go", GoEqec06},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(date1, date2, dr, dd)
			}
		})
	}
}
