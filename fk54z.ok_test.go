package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t F k 5 4 z
//  - - - - - - - - - -
//
//  Test Fk54z function.
//
//  Called:  Fk54z, vvd
//
//  This revision:  2018 December 6
//
func TestFk54z(t *testing.T) {
	const fname = "Fk54z"
	var r2000, d2000, bepoch, r1950, d1950, dr1950, dd1950 float64

	r2000 = 0.02719026625066316119
	d2000 = -0.1115815170738754813
	bepoch = 1954.677308160316374

	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (c1, c2, c3, c4 float64)
	}{
		{"cgo", CgoFk54z},
		{"go", GoFk54z},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		r1950, d1950, dr1950, dd1950 = test.fn(r2000, d2000, bepoch)

		vvd(t, r1950, 0.01602015588390065476, 1e-14,
			tname, "r1950")
		vvd(t, d1950, -0.1164397101110765346, 1e-13,
			tname, "d1950")
		vvd(t, dr1950, -0.1175712648471090704e-7, 1e-20,
			tname, "dr1950")
		vvd(t, dd1950, 0.2108109051316431056e-7, 1e-20,
			tname, "dd1950")
	}
}

func BenchmarkFk54z(b *testing.B) {
	var r2000, d2000, bepoch float64

	r2000 = 0.02719026625066316119
	d2000 = -0.1115815170738754813
	bepoch = 1954.677308160316374

	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (c1, c2, c3, c4 float64)
	}{
		{"cgo", CgoFk54z},
		{"go", GoFk54z},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(r2000, d2000, bepoch)
			}
		})
	}
}
