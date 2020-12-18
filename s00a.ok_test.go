package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t S 0 0 a
//  - - - - - - - - -
//
//  Test S00a function.
//
//  Called:  S00a, vvd
//
//  This revision:  2013 August 7
//
func TestS00a(t *testing.T) {
	const fname = "S00a"
	var s float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoS00a},
		{"go", GoS00a},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		s = test.fn(2400000.5, 52541.0)

		vvd(t, s, -0.1340684448919163584e-7, 1e-18, tname, "")
	}
}

func BenchmarkS00a(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoS00a},
		{"go", GoS00a},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 52541.0)
			}
		})
	}
}
