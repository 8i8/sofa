package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t P v x p v
//  - - - - - - - - - -
//
//  Test Pvxpv function.
//
//  Called:  Pvxpv, vvd
//
//  This revision:  2013 August 7
//
func TestPvxpv(t *testing.T) {
	const fname = "Pvxpv"
	var axb [2][3]float64
	var a, b [2][3]float64

	a[0][0] = 2.0
	a[0][1] = 2.0
	a[0][2] = 3.0

	a[1][0] = 6.0
	a[1][1] = 0.0
	a[1][2] = 4.0

	b[0][0] = 1.0
	b[0][1] = 3.0
	b[0][2] = 4.0

	b[1][0] = 0.0
	b[1][1] = 2.0
	b[1][2] = 8.0

	tests := []struct {
		ref string
		fn  func(a1,a2 [2][3]float64) [2][3]float64
	}{
		{"cgo", CgoPvxpv},
		{"go", GoPvxpv},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		axb = test.fn(a, b)

		vvd(t, axb[0][0], -1.0, 1e-12, tname, "p1")
		vvd(t, axb[0][1], -5.0, 1e-12, tname, "p2")
		vvd(t, axb[0][2], 4.0, 1e-12, tname, "p3")

		vvd(t, axb[1][0], -2.0, 1e-12, tname, "v1")
		vvd(t, axb[1][1], -36.0, 1e-12, tname, "v2")
		vvd(t, axb[1][2], 22.0, 1e-12, tname, "v3")
	}
}

func BenchmarkPvxpv(bm *testing.B) {
	var a, b [2][3]float64

	a[0][0] = 2.0
	a[0][1] = 2.0
	a[0][2] = 3.0

	a[1][0] = 6.0
	a[1][1] = 0.0
	a[1][2] = 4.0

	b[0][0] = 1.0
	b[0][1] = 3.0
	b[0][2] = 4.0

	b[1][0] = 0.0
	b[1][1] = 2.0
	b[1][2] = 8.0

	tests := []struct {
		ref string
		fn  func(a1,a2 [2][3]float64) [2][3]float64
	}{
		{"cgo", CgoPvxpv},
		{"go", GoPvxpv},
	}

	for _, test := range tests {
		bm.Run(test.ref, func(bm *testing.B) {
			for i := 0; i < bm.N; i++ {
				_ = test.fn(a, b)
			}
		})
	}
}
