package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t F a u r 0 3
//  - - - - - - - - - - -
//
//  Test Faur03 function.
//
//  Called:  Faur03, vvd
//
//  This revision:  2013 August 7
//
func TestFaur03(t *testing.T) {
	const fname = "Faur03"
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFaur03},
		{"go", GoFaur03},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		vvd(t, test.fn(0.80), 5.180636450180413523, 1e-12,
			tname, "")
	}
}

func BenchmarkFaur03(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFaur03},
		{"go", GoFaur03},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(0.80)
			}
		})
	}
}
