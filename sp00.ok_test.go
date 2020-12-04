package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t S p 0 0
//  - - - - - - - - -
//
//  Test Sp00 function.
//
//  Called:  Sp00, vvd
//
//  This revision:  2013 August 7
//
func TestSp00(t *testing.T) {
	const fname = "Sp00"
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoSp00},
		{"go", GoSp00},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		vvd(t, test.fn(2400000.5, 52541.0),
			-0.6216698469981019309e-11, 1e-12, tname, "")
	}
}

func BenchmarkSp00(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoSp00},
		{"go", GoSp00},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(2400000.5, 52541.0)
			}
		})
	}
}
