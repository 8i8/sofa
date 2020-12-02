package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t F a n e 0 3
//  - - - - - - - - - - -
//
//  Test Fane03 function.
//
//  Called:  Fane03, vvd
//
//  This revision:  2013 August 7
//
func TestFane03(t *testing.T) {
	const fname = "Fane03"
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFane03},
		{"go", GoFane03},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		vvd(t, test.fn(0.80), 2.079343830860413523, 1e-12,
			tname, "")
	}
}

func BenchmarkFane03(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFane03},
		{"go", GoFane03},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(0.80)
			}
		})
	}
}
