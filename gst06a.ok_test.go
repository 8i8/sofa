package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t G s t 0 6 a
//  - - - - - - - - - - -
//
//  Test Gst06a function.
//
//  Called:  Gst06a, vvd
//
//  This revision:  2013 August 7
//
func TestGst06a(t *testing.T) {
	const fname = "Gst06a"
	var theta float64

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) float64
	}{
		{"cgo", CgoGst06a},
		{"go", GoGst06a},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		theta = test.fn(2400000.5, 53736.0, 2400000.5, 53736.0)

		vvd(t, theta, 1.754166137675019159, 1e-12, tname, "")
	}
}

func BenchmarkGst06a(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) float64
	}{
		{"cgo", CgoGst06a},
		{"go", GoGst06a},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736.0, 2400000.5, 53736.0)
			}
		})
	}
}
