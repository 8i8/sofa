package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t P f w 0 6
//  - - - - - - - - - -
//
//  Test Pfw06 function.
//
//  Called:  Pfw06, vvd
//
//  This revision:  2013 August 7
//
func TestPfw06(t *testing.T) {
	const fname = "Pfw06"
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (b1, b2, b3, b4 float64)
	}{
		{"cgo", CgoPfw06},
		{"go", GoPfw06},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		gamb, phib, psib, epsa := test.fn(2400000.5, 50123.9999)

		vvd(t, gamb, -0.2243387670997995690e-5, 1e-16,
			tname, "gamb")
		vvd(t, phib, 0.4091014602391312808, 1e-12,
			tname, "phib")
		vvd(t, psib, -0.9501954178013031895e-3, 1e-14,
			tname, "psib")
		vvd(t, epsa, 0.4091014316587367491, 1e-12,
			tname, "epsa")
	}
}

func BenchmarkPfw06(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (b1, b2, b3, b4 float64)
	}{
		{"cgo", CgoPfw06},
		{"go", GoPfw06},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(2400000.5, 50123.9999)
			}
		})
	}
}
