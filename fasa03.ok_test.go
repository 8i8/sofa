package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t F a s a 0 3
//  - - - - - - - - - - -
//
//  Test Fasa03 function.
//
//  Called:  Fasa03, vvd
//
//  This revision:  2013 August 7
//
func TestFasa03(t *testing.T) {
	const fname = "Fasa03"
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFasa03},
		{"go", GoFasa03},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		vvd(t, test.fn(0.80), 5.371574539440827046, 1e-12,
			tname, "")
	}
}

func BenchmarkFasa03(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFasa03},
		{"go", GoFasa03},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(0.80)
			}
		})
	}
}
