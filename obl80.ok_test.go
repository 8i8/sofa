package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t O b l 8 0
//  - - - - - - - - - -
//
//  Test Obl80 function.
//
//  Called:  Obl80, vvd
//
//  This revision:  2013 August 7
//
func TestObl80(t *testing.T) {
	const fname = "Obl80"
	var eps0 float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoObl80},
		{"go", GoObl80},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		eps0 = test.fn(2400000.5, 54388.0)

		vvd(t, eps0, 0.4090751347643816218, 1e-14, tname, "")
	}
}

func BenchmarkObl80(b *testing.B) {

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoObl80},
		{"go", GoObl80},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 54388.0)
			}
		})
	}
}
