package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t O b l 0 6
//  - - - - - - - - - -
//
//  Test Obl06 function.
//
//  Called:  Obl06, vvd
//
//  This revision:  2013 August 7
//
func TestObl06(t *testing.T) {
	const fname = "Olb06"
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoObl06},
		{"go", GoObl06},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		vvd(t, test.fn(2400000.5, 54388.0), 0.4090749229387258204, 1e-14,
			tname, "")
	}
}

func BenchmarkObl06(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoObl06},
		{"go", GoObl06},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 54388.0)
			}
		})
	}
}
