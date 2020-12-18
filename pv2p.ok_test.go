package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t P v 2 p
//  - - - - - - - - -
//
//  Test Pv2p function.
//
//  Called:  Pv2p, vvd
//
//  This revision:  2013 August 7
//
func TestPv2p(t *testing.T) {
	const fname = "Pv2p"
	var p [3]float64
	var pv [2][3]float64

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = -0.5
	pv[1][1] = 3.1
	pv[1][2] = 0.9

	tests := []struct {
		ref string
		fn  func([2][3]float64) [3]float64
	}{
		{"cgo", CgoPv2p},
		{"go", GoPv2p},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		p = test.fn(pv)

		vvd(t, p[0], 0.3, 0.0, tname, "1")
		vvd(t, p[1], 1.2, 0.0, tname, "2")
		vvd(t, p[2], -2.5, 0.0, tname, "3")
	}
}

func BenchmarkPv2p(b *testing.B) {
	var pv [2][3]float64

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = -0.5
	pv[1][1] = 3.1
	pv[1][2] = 0.9

	tests := []struct {
		ref string
		fn  func([2][3]float64) [3]float64
	}{
		{"cgo", CgoPv2p},
		{"go", GoPv2p},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(pv)
			}
		})
	}
}
