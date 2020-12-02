package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t F a v e 0 3
//  - - - - - - - - - - -
//
//  Test Fave03 function.
//
//  Called:  Fave03, vvd
//
//  This revision:  2013 August 7
//
func TestFave03(t *testing.T) {
	const fname = "Fave03"
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFave03},
		{"go", GoFave03},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		vvd(t, test.fn(0.80), 3.424900460533758000, 1e-12,
			tname, "")
	}
}

func BenchmarkFave03(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFave03},
		{"go", GoFave03},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(0.80)
			}
		})
	}
}
