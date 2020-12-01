package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t F a e 0 3
//  - - - - - - - - - -
//
//  Test iauFae03 function.
//
//  Called:  Fae03, vvd
//
//  This revision:  2013 August 7
//
func TestFae03(t *testing.T) {
	const fname = "Fae03"

	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", Fae03},
		{"go", goFae03},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		vvd(t, test.fn(0.80), 1.744713738913081846, 1e-12,
			tname, "")
	}
}

func BenchmarkFae03(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", Fae03},
		{"go", goFae03},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(0.80)
			}
		})
	}
}
