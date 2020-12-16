package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t G m s t 8 2
//  - - - - - - - - - - -
//
//  Test Gmst82 function.
//
//  Called:  Gmst82, vvd
//
//  This revision:  2013 August 7
//
func TestGmst82(t *testing.T) {
	const fname = "Gmst82"
	var theta float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoGmst82},
		{"go", GoGmst82},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		theta = test.fn(2400000.5, 53736.0)

		vvd(t, theta, 1.754174981860675096, 1e-12, tname, "")
	}
}

func BenchmarkGmst82(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoGmst82},
		{"go", GoGmst82},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736.0)
			}
		})
	}
}
