package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t F a p a 0 3
//  - - - - - - - - - - -
//
//  Test Fapa03 function.
//
//  Called:  Fapa03, vvd
//
//  This revision:  2013 August 7
//
func TestFapa03(t *testing.T) {
	const fname = "Fapa03"
	tests := []struct {
		ref string
		fn  func(t float64) float64
	}{
		{"cgo", CgoFapa03},
		{"go", GoFapa03},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		vvd(t, test.fn(0.80), 0.1950884762240000000e-1, 1e-12,
			tname, "")
	}
}

func BenchmarkFapa03(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(t float64) float64
	}{
		{"cgo", CgoFapa03},
		{"go", GoFapa03},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(0.80)
			}
		})
	}
}
