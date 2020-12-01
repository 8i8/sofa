package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t F a j u 0 3
//  - - - - - - - - - - -
//
//  Test Faju03 function.
//
//  Called:  Faju03, vvd
//
//  This revision:  2013 August 7
//
func TestFaju03(t *testing.T) {
	const fname = "Faju03"
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFaju03},
		{"go", GoFaju03},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		vvd(t, test.fn(0.80), 5.275711665202481138, 1e-12,
			tname, "")
	}
}

func BenchmarkFaju03(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(float64) float64
	}{
		{"cgo", CgoFaju03},
		{"go", GoFaju03},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(0.80)
			}
		})
	}
}
