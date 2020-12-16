package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t G m s t 0 6
//  - - - - - - - - - - -
//
//  Test Gmst06 function.
//
//  Called:  Gmst06, vvd
//
//  This revision:  2013 August 7
//
func TestGmst06(t *testing.T) {
	const fname = "Gmst06"
	var theta float64

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) float64
	}{
		{"cgo", CgoGmst06},
		{"go", GoGmst06},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		theta = test.fn(2400000.5, 53736.0, 2400000.5, 53736.0)

		vvd(t, theta, 1.754174971870091203, 1e-12, tname, "")
	}
}

func BenchmarkGmst06(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) float64
	}{
		{"cgo", CgoGmst06},
		{"go", GoGmst06},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736.0, 2400000.5, 53736.0)
			}
		})
	}
}
