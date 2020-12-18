package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t S 2 p
//  - - - - - - - -
//
//  Test S2p function.
//
//  Called:  S2p, vvd
//
//  This revision:  2013 August 7
//
func TestS2p(t *testing.T) {
	const fname = "S2p"
	var p [3]float64

	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) [3]float64
	}{
		{"cgo", CgoS2p},
		{"go", GoS2p},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		p = test.fn(-3.21, 0.123, 0.456)

		vvd(t, p[0], -0.4514964673880165228, 1e-12, tname, "x")
		vvd(t, p[1], 0.0309339427734258688, 1e-12, tname, "y")
		vvd(t, p[2], 0.0559466810510877933, 1e-12, tname, "z")
	}
}

func BenchmarkS2p(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) [3]float64
	}{
		{"cgo", CgoS2p},
		{"go", GoS2p},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(-3.21, 0.123, 0.456)
			}
		})
	}
}
