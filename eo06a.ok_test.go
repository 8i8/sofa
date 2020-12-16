package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t E o 0 6 a
//  - - - - - - - - - -
//
//  Test Eo06a function.
//
//  Called:  iauEo06a, vvd
//
//  This revision:  2013 August 7
//
func TestEo06a(t *testing.T) {
	const fname = "Eo06a"
	var eo float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoEo06a},
		{"go", GoEo06a},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		eo = test.fn(2400000.5, 53736.0)

		vvd(t, eo, -0.1332882371941833644e-2, 1e-15, tname, "")
	}
}

func BenchmarkEo06a(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoEo06a},
		{"go", GoEo06a},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736.0)
			}
		})
	}
}
