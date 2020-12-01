package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t F a m a 0 3
//  - - - - - - - - - - -
//
//  Test Fama03 function.
//
//  Called:  Fama03, vvd
//
//  This revision:  2013 August 7
//
func TestFama03(t *testing.T) {
	const fname = "Fama03"
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFama03},
		{"go", GoFama03},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		vvd(t, test.fn(0.80), 3.275506840277781492, 1e-12,
			tname, "")
	}
}

func BenchmarkFama03(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFama03},
		{"go", GoFama03},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(0.80)
			}
		})
	}
}
