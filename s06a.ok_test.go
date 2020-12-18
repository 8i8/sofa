package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t S 0 6 a
//  - - - - - - - - -
//
//  Test S06a function.
//
//  Called:  S06a, vvd
//
//  This revision:  2013 August 7
//
func TestS06a(t *testing.T) {
	const fname = "S06a"
	var s float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoS06a},
		{"go", GoS06a},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		s = test.fn(2400000.5, 52541.0)

		vvd(t, s, -0.1340680437291812383e-7, 1e-18, tname, "")
	}
}

func BenchmarkS06a(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoS06a},
		{"go", GoS06a},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 52541.0)
			}
		})
	}
}
