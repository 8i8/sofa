package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t D 2 d t f
//  - - - - - - - - - -
//
//  Test D2dtf function.
//
//  Called:  D2dtf, viv
//
//  This revision:  2013 August 7
//
func TestD2dtf(t *testing.T) {
	const fname = "D2dtf"
	var iy, im, id int
	var ihmsf [4]int
	var err error

	tests := []struct {
		ref string
		fn  func(string, int, float64, float64) (
			int, int, int, [4]int, error)
	}{
		{"cgo", CgoD2dtf},
		{"go", GoD2dtf},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		iy, im, id, ihmsf, err = test.fn(
			"UTC", 5, 2400000.5, 49533.99999)

		viv(t, iy, 1994, tname, "y")
		viv(t, im, 6, tname, "mo")
		viv(t, id, 30, tname, "d")
		viv(t, ihmsf[0], 23, tname, "h")
		viv(t, ihmsf[1], 59, tname, "m")
		viv(t, ihmsf[2], 60, tname, "s")
		viv(t, ihmsf[3], 13599, tname, "f")
		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkD2dtf(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(string, int, float64, float64) (
			int, int, int, [4]int, error)
	}{
		{"cgo", CgoD2dtf},
		{"go", GoD2dtf},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn("UTC", 5, 2400000.5, 49533.99999)
			}
		})
	}
}
