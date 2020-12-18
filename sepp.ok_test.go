package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t S e p p
//  - - - - - - - - -
//
//  Test Sepp function.
//
//  Called:  Sepp, vvd
//
//  This revision:  2013 August 7
//
func TestSepp(t *testing.T) {
	const fname = "Sepp"
	var s float64
	var a, b [3]float64

	a[0] = 1.0
	a[1] = 0.1
	a[2] = 0.2

	b[0] = -3.0
	b[1] = 1e-3
	b[2] = 0.2

	tests := []struct {
		ref string
		fn  func(a1, a2 [3]float64) float64
	}{
		{"cgo", CgoSepp},
		{"go", GoSepp},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		s = test.fn(a, b)

		vvd(t, s, 2.860391919024660768, 1e-12, tname, "")
	}
}

func BenchmarkSepp(bm *testing.B) {
	var a, b [3]float64

	a[0] = 1.0
	a[1] = 0.1
	a[2] = 0.2

	b[0] = -3.0
	b[1] = 1e-3
	b[2] = 0.2

	tests := []struct {
		ref string
		fn  func(a1, a2 [3]float64) float64
	}{
		{"cgo", CgoSepp},
		{"go", GoSepp},
	}

	for _, test := range tests {
		bm.Run(test.ref, func(bm *testing.B) {
			for i := 0; i < bm.N; i++ {
				_ = test.fn(a, b)
			}
		})
	}
}
