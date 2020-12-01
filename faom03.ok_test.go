package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t F a o m 0 3
//  - - - - - - - - - - -
//
//  Test Faom03 function.
//
//  Called:  iauFaom03, vvd
//
//  This revision:  2013 August 7
//
func TestFaom03(t *testing.T) {
	const fname = "Faom03"
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFaom03},
		{"go", GoFaom03},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		vvd(t, test.fn(0.80), -5.973618440951302183, 1e-12,
			tname, "")
	}
}

func BenchmarkFaom03(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFaom03},
		{"go", GoFaom03},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(0.80)
			}
		})
	}
}
