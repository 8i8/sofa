package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t E p j
//  - - - - - - - -
//
//  Test Epj function.
//
//  Called:  Epj, vvd
//
//  This revision:  2013 August 7
//
func TestEpj(t *testing.T) {
	const fname = "Epj"
	var epj float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoEpj},
		{"go", GoEpj},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		epj = test.fn(2451545, -7392.5)

		vvd(t, epj, 1979.760438056125941, 1e-12, tname, "")
	}
}

func BenchmarkEpj(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoEpj},
		{"go", GoEpj},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2451545, -7392.5)
			}
		})
	}
}
