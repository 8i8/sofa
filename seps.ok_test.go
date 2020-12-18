package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t S e p s
//  - - - - - - - - -
//
//  Test Seps function.
//
//  Called:  Seps, vvd
//
//  This revision:  2013 August 7
//
func TestSeps(t *testing.T) {
	const fname = "Seps"
	var al, ap, bl, bp, s float64

	al = 1.0
	ap = 0.1

	bl = 0.2
	bp = -3.0

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) float64
	}{
		{"cgo", CgoSeps},
		{"go", GoSeps},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		s = test.fn(al, ap, bl, bp)

		vvd(t, s, 2.346722016996998842, 1e-14, tname, "")
	}
}

func BenchmarkSeps(b *testing.B) {
	var al, ap, bl, bp float64

	al = 1.0
	ap = 0.1

	bl = 0.2
	bp = -3.0

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) float64
	}{
		{"cgo", CgoSeps},
		{"go", GoSeps},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(al, ap, bl, bp)
			}
		})
	}
}
