package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t F a d 0 3
//  - - - - - - - - - -
//
//  Test Fad03 function.
//
//  Called:  Fad03, vvd
//
//  This revision:  2013 August 7
//
func TestFad03(t *testing.T) {
	const fname = "Fad03"

	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFad03},
		{"go", GoFad03},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		vvd(t, test.fn(0.80), 1.946709205396925672, 1e-12, tname, "")
	}
}

func BenchmarkFad03(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFad03},
		{"go", GoFad03},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(0.80)
			}
		})
	}
}
