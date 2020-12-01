package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t A n p m
//  - - - - - - - - -
//
//  Test Anpm function.
//
//  Called:  Anpm, vvd
//
//  This revision:  2013 August 7
//
func TestAnpm(t *testing.T) {
	const fname = "Anpm"
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoAnpm},
		{"go", GoAnpm},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		vvd(t, test.fn(-4.0), 2.283185307179586477,
			1e-12, tname, "")
	}
}

func BenchmarkAnpm(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoAnpm},
		{"go", GoAnpm},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(-4.0)
			}
		})
	}
}
