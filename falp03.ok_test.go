package sofa

import "testing"

//
//  - - - - - - - - -
//   t _ f a l p 0 3
//  - - - - - - - - -
//
//  Test Falp03 function.
//
//  Called:  Falp03, vvd
//
//  This revision:  2013 August 7
//
func TestFalp03(t *testing.T) {
	const fname = "Falp03"
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFalp03},
		{"go", GoFalp03},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		vvd(t, test.fn(0.80), 6.226797973505507345, 1e-12,
			tname, "")
	}
}

func BenchmarkFalp03(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFalp03},
		{"go", GoFalp03},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(0.80)
			}
		})
	}
}
