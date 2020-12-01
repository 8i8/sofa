package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t F a l 0 3
//  - - - - - - - - - -
//
//  Test Fal03 function.
//
//  Called:  Fal03, vvd
//
//  This revision:  2013 August 7
//
func TestFal03(t *testing.T) {
	const fname = "Fal03"
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFal03},
		{"go", GoFal03},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		vvd(t, test.fn(0.80), 5.132369751108684150, 1e-12,
			tname, "")
	}
}

func BenchmarkFal03(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFal03},
		{"go", GoFal03},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(0.80)
			}
		})
	}
}
