package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t G s t 9 4
//  - - - - - - - - - -
//
//  Test Gst94 function.
//
//  Called:  Gst94, vvd
//
//  This revision:  2013 August 7
//
func TestGst94(t *testing.T) {
	const fname = "Gst94"
	var theta float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoGst94},
		{"go", GoGst94},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		theta = test.fn(2400000.5, 53736.0)

		vvd(t, theta, 1.754166136020645203, 1e-12, tname, "")
	}
}

func BenchmarkGst94(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoGst94},
		{"go", GoGst94},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736.0)
			}
		})
	}
}
