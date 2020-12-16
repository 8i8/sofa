package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t G s t 0 0 b
//  - - - - - - - - - - -
//
//  Test Gst00b function.
//
//  Called:  Gst00b, vvd
//
//  This revision:  2013 August 7
//
func TestGst00b(t *testing.T) {
	const fname = "Gst00b"
	var theta float64

	tests := []struct {
		ref string
		fn  func(a1,a2 float64) float64
	}{
		{"cgo", CgoGst00b},
		{"go", GoGst00b},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		theta = test.fn(2400000.5, 53736.0)

		vvd(t, theta, 1.754166136510680589, 1e-12, tname, "")
	}
}

func BenchmarkGst00b(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1,a2 float64) float64
	}{
		{"cgo", CgoGst00b},
		{"go", GoGst00b},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736.0)
			}
		})
	}
}
