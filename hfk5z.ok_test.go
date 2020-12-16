package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t H f k 5 z
//  - - - - - - - - - -
//
//  Test Hfk5z function.
//
//  Called:  Hfk5z, vvd
//
//  This revision:  2013 August 7
//
func TestHfk5z(t *testing.T) {
	const fname = "Hfk5z"
	var rh, dh, r5, d5, dr5, dd5 float64

	rh = 1.767794352
	dh = -0.2917512594

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (
			c1, c2, c3, c4 float64)
	}{
		{"cgo", CgoHfk5z},
		{"go", GoHfk5z},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		r5, d5, dr5, dd5 = test.fn(rh, dh, 2400000.5, 54479.0)

		vvd(t, r5, 1.767794490535581026, 1e-13,
			tname, "ra")
		vvd(t, d5, -0.2917513695320114258, 1e-14,
			tname, "dec")
		vvd(t, dr5, 0.4335890983539243029e-8, 1e-22,
			tname, "dr5")
		vvd(t, dd5, -0.8569648841237745902e-9, 1e-23,
			tname, "dd5")
	}
}

func BenchmarkHfk5z(b *testing.B) {
	var rh, dh float64

	rh = 1.767794352
	dh = -0.2917512594

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (
			c1, c2, c3, c4 float64)
	}{
		{"cgo", CgoHfk5z},
		{"go", GoHfk5z},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(rh, dh, 2400000.5, 54479.0)
			}
		})
	}
}
