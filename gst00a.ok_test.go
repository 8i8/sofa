package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t G s t 0 0 a
//  - - - - - - - - - - -
//
//  Test Gst00a function.
//
//  Called:  Gst00a, vvd
//
//  This revision:  2013 August 7
//
func TestGst00a(t *testing.T) {
	const fname = "Gst00a"
	var theta float64

	tests := []struct {
		ref string
		fn  func(a1,a2,a3,a4 float64) float64
	}{
		{"cgo", CgoGst00a},
		{"go", GoGst00a},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		theta = test.fn(2400000.5, 53736.0, 2400000.5, 53736.0)

		vvd(t, theta, 1.754166138018281369, 1e-12, tname, "")
	}
}

func BenchmarkGst00a(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1,a2,a3,a4 float64) float64
	}{
		{"cgo", CgoGst00a},
		{"go", GoGst00a},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736.0,
					2400000.5, 53736.0)
			}
		})
	}
}
