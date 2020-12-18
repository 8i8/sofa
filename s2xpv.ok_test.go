package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t S 2 x p v
//  - - - - - - - - - -
//
//  Test S2xpv function.
//
//  Called:  S2xpv, vvd
//
//  This revision:  2013 August 7
//
func TestS2xpv(t *testing.T) {
	const fname = "S2xpv"
	var s1, s2 float64
	var pv, spv [2][3]float64

	s1 = 2.0
	s2 = 3.0

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = 0.5
	pv[1][1] = 2.3
	pv[1][2] = -0.4

	tests := []struct {
		ref string
		fn  func(a1, a2 float64, a3 [2][3]float64) [2][3]float64
	}{
		{"cgo", CgoS2xpv},
		{"go", GoS2xpv},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		spv = test.fn(s1, s2, pv)

		vvd(t, spv[0][0], 0.6, 1e-12, tname, "p1")
		vvd(t, spv[0][1], 2.4, 1e-12, tname, "p2")
		vvd(t, spv[0][2], -5.0, 1e-12, tname, "p3")

		vvd(t, spv[1][0], 1.5, 1e-12, tname, "v1")
		vvd(t, spv[1][1], 6.9, 1e-12, tname, "v2")
		vvd(t, spv[1][2], -1.2, 1e-12, tname, "v3")
	}
}

func BenchmarkS2xpv(b *testing.B) {
	var s1, s2 float64
	var pv [2][3]float64

	s1 = 2.0
	s2 = 3.0

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = 0.5
	pv[1][1] = 2.3
	pv[1][2] = -0.4

	tests := []struct {
		ref string
		fn  func(a1, a2 float64, a3 [2][3]float64) [2][3]float64
	}{
		{"cgo", CgoS2xpv},
		{"go", GoS2xpv},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(s1, s2, pv)
			}
		})
	}
}
