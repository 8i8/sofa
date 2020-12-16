package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t S 2 p v
//  - - - - - - - - -
//
//  Test S2pv function.
//
//  Called:  S2pv, vvd
//
//  This revision:  2013 August 7
//
func TestS2pv(t *testing.T) {
	const fname = "S2pv"
	var pv [2][3]float64

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6 float64) [2][3]float64
	}{
		{"cgo", CgoS2pv},
		{"go", GoS2pv},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		pv = test.fn(-3.21, 0.123, 0.456, -7.8e-6, 9.01e-6,
			-1.23e-5)

		vvd(t, pv[0][0], -0.4514964673880165228, 1e-12,
			tname, "x")
		vvd(t, pv[0][1], 0.0309339427734258688, 1e-12,
			tname, "y")
		vvd(t, pv[0][2], 0.0559466810510877933, 1e-12,
			tname, "z")

		vvd(t, pv[1][0], 0.1292270850663260170e-4, 1e-16,
			tname, "vx")
		vvd(t, pv[1][1], 0.2652814182060691422e-5, 1e-16,
			tname, "vy")
		vvd(t, pv[1][2], 0.2568431853930292259e-5, 1e-16,
			tname, "vz")
	}
}

func BenchmarkS2pv(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6 float64) [2][3]float64
	}{
		{"cgo", CgoS2pv},
		{"go", GoS2pv},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(-3.21, 0.123, 0.456,
				-7.8e-6, 9.01e-6, -1.23e-5)
			}
		})
	}
}
