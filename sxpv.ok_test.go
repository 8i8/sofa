package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t S x p v
//  - - - - - - - - -
//
//  Test Sxpv function.
//
//  Called:  Sxpv, vvd
//
//  This revision:  2013 August 7
//
func TestSxpv(t *testing.T) {
	const fname = "Sxpv"
	var s float64
	var pv, spv [2][3]float64

	s = 2.0

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = 0.5
	pv[1][1] = 3.2
	pv[1][2] = -0.7

	tests := []struct {
		ref string
		fn  func(float64, [2][3]float64) [2][3]float64
	}{
		{"cgo", CgoSxpv},
		{"go", GoSxpv},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		spv = test.fn(s, pv)

		vvd(t, spv[0][0], 0.6, 0.0, tname, "p1")
		vvd(t, spv[0][1], 2.4, 0.0, tname, "p2")
		vvd(t, spv[0][2], -5.0, 0.0, tname, "p3")

		vvd(t, spv[1][0], 1.0, 0.0, tname, "v1")
		vvd(t, spv[1][1], 6.4, 0.0, tname, "v2")
		vvd(t, spv[1][2], -1.4, 0.0, tname, "v3")
	}
}

func BenchmarkSxpv(b *testing.B) {
	var s float64
	var pv [2][3]float64

	s = 2.0

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = 0.5
	pv[1][1] = 3.2
	pv[1][2] = -0.7

	tests := []struct {
		ref string
		fn  func(float64, [2][3]float64) [2][3]float64
	}{
		{"cgo", CgoSxpv},
		{"go", GoSxpv},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(s, pv)
			}
		})
	}
}
