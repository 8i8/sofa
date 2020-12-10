package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t S 2 c
//  - - - - - - - -
//
//  Test S2c function.
//
//  Called:  S2c, vvd
//
//  This revision:  2013 August 7
//
func TestS2c(t *testing.T) {
	const fname = "S2c"

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3]float64
	}{
		{"cgo", CgoS2c},
		{"go", GoS2c},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		c := test.fn(3.0123, -0.999)

		vvd(t, c[0], -0.5366267667260523906, 1e-12, tname, "1")
		vvd(t, c[1], 0.0697711109765145365, 1e-12, tname, "2")
		vvd(t, c[2], -0.8409302618566214041, 1e-12, tname, "3")
	}
}

func BenchmarkS2c(b *testing.B) {

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3]float64
	}{
		{"cgo", CgoS2c},
		{"go", GoS2c},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(3.0123, -0.999)
			}
		})
	}
}
