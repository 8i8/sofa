package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t A p e r
//  - - - - - - - - -
//
//  Test Aper function.
//
//  Called:  Aper, vvd
//
//  This revision:  2013 October 3
//
func TestAper(t *testing.T) {
	const fname = "Aper"
	var theta float64
	var astrom ASTROM

	astrom.along = 1.234
	theta = 5.678

	tests := []struct {
		ref string
		fn  func(float64, ASTROM) ASTROM
	}{
		{"cgo", CgoAper},
		{"go", GoAper},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		astrom2 := test.fn(theta, astrom)

		vvd(t, astrom2.eral, 6.912000000000000000, 1e-12,
			tname, "pmt")
	}
}

func BenchmarkAper(b *testing.B) {
	var theta float64
	var astrom ASTROM

	astrom.along = 1.234
	theta = 5.678

	tests := []struct {
		ref string
		fn  func(float64, ASTROM) ASTROM
	}{
		{"cgo", CgoAper},
		{"go", GoAper},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(theta, astrom)
			}
		})
	}
}
