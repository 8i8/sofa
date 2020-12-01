package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t F a f 0 3
//  - - - - - - - - - -
//
//  Test Faf03 function.
//
//  Called:  Faf03, vvd
//
//  This revision:  2013 August 7
//
func TestFaf03(t *testing.T) {
	const fname = "Faf03"

	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFaf03},
		{"go", GoFaf03},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		vvd(t, test.fn(0.80), 0.2597711366745499518, 1e-12,
			tname, "")
	}
}

func BenchmarkFaf03(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFaf03},
		{"go", GoFaf03},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(0.80)
			}
		})
	}
}
