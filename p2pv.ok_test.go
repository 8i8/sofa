package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t p 2 p v
//  - - - - - - - - -
//
//  Test P2pv function.
//
//  Called:  P2pv, vvd
//
//  This revision:  2013 August 7
//
func TestP2pv(t *testing.T) {
	const fname = "P2pv"

	var p [3]float64
	var pv [2][3]float64

	p[0] = 0.25
	p[1] = 1.2
	p[2] = 3.0

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = -0.5
	pv[1][1] = 3.1
	pv[1][2] = 0.9

	tests := []struct {
		ref string
		fn  func([3]float64) [2][3]float64
	}{
		{"cgo", CgoP2pv},
		{"go", GoP2pv},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		pv = test.fn(p)

		vvd(t, pv[0][0], 0.25, 0.0, tname, "p1")
		vvd(t, pv[0][1], 1.2, 0.0, tname, "p2")
		vvd(t, pv[0][2], 3.0, 0.0, tname, "p3")

		vvd(t, pv[1][0], 0.0, 0.0, tname, "v1")
		vvd(t, pv[1][1], 0.0, 0.0, tname, "v2")
		vvd(t, pv[1][2], 0.0, 0.0, tname, "v3")
	}
}

func BenchmarkP2pv(b *testing.B) {
	var p [3]float64

	p[0] = 0.25
	p[1] = 1.2
	p[2] = 3.0

	tests := []struct {
		ref string
		fn  func([3]float64) [2][3]float64
	}{
		{"cgo", CgoP2pv},
		{"go", GoP2pv},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(p)
			}
		})
	}
}
