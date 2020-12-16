package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t F k 4 5 z
//  - - - - - - - - - -
//
//  Test Fk45z function.
//
//  Called:  Fk45z, vvd
//
//  This revision:  2018 December 6
//
func TestFk45z(t *testing.T) {
	const fname = "Fk45z"
	var r1950, d1950, bepoch, r2000, d2000 float64

	r1950 = 0.01602284975382960982
	d1950 = -0.1164347929099906024
	bepoch = 1954.677617625256806

	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (c1, c2 float64)
	}{
		{"cgo", CgoFk45z},
		{"go", GoFk45z},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref

		r2000, d2000 = test.fn(r1950, d1950, bepoch)

		vvd(t, r2000, 0.02719295911606862303, 1e-15,
			tname, "r2000")
		vvd(t, d2000, -0.1115766001565926892, 1e-13,
			tname, "d2000")
	}
}

func BenchmarkFk45z(b *testing.B) {
	var r1950, d1950, bepoch float64

	r1950 = 0.01602284975382960982
	d1950 = -0.1164347929099906024
	bepoch = 1954.677617625256806

	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (c1, c2 float64)
	}{
		{"cgo", CgoFk45z},
		{"go", GoFk45z},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(r1950, d1950, bepoch)
			}
		})
	}
}
