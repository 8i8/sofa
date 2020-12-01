package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t A n p
//  - - - - - - - -
//
//  Test Anp function.
//
//  Called:  Anp, vvd
//
//  This revision:  2013 August 7
//
func TestAnp(t *testing.T) {
	const fname = "Anp"
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoAnp},
		{"go", GoAnp},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		vvd(t, test.fn(-0.1), 6.183185307179586477,
			1e-12, tname, "")
	}
}

func BenchmarkAnp(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoAnp},
		{"go", GoAnp},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(-0.1)
			}
		})
	}
}
