package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t R x p v
//  - - - - - - - - -
//
//  Test Rxpv function.
//
//  Called:  Rxpv, vvd
//
//  This revision:  2013 August 7
//
func TestTrxpv(t *testing.T) {
	const fname = "Trxpv"
	var r [3][3]float64
	var pv [2][3]float64

	r[0][0] = 2.0
	r[0][1] = 3.0
	r[0][2] = 2.0

	r[1][0] = 3.0
	r[1][1] = 2.0
	r[1][2] = 3.0

	r[2][0] = 3.0
	r[2][1] = 4.0
	r[2][2] = 5.0

	pv[0][0] = 0.2
	pv[0][1] = 1.5
	pv[0][2] = 0.1

	pv[1][0] = 1.5
	pv[1][1] = 0.2
	pv[1][2] = 0.1

	tests := []struct {
		ref string
		fn  func([3][3]float64, [2][3]float64) [2][3]float64
	}{
		{"cgo", CgoTrxpv},
		{"go", GoTrxpv},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		trpv := test.fn(r, pv)

		vvd(t, trpv[0][0], 5.2, 1e-12, tname, "p1")
		vvd(t, trpv[0][1], 4.0, 1e-12, tname, "p1")
		vvd(t, trpv[0][2], 5.4, 1e-12, tname, "p1")

		vvd(t, trpv[1][0], 3.9, 1e-12, tname, "v1")
		vvd(t, trpv[1][1], 5.3, 1e-12, tname, "v2")
		vvd(t, trpv[1][2], 4.1, 1e-12, tname, "v3")
	}
}

func BenchmarkTrxpv(b *testing.B) {
	var r [3][3]float64
	var pv [2][3]float64

	r[0][0] = 2.0
	r[0][1] = 3.0
	r[0][2] = 2.0

	r[1][0] = 3.0
	r[1][1] = 2.0
	r[1][2] = 3.0

	r[2][0] = 3.0
	r[2][1] = 4.0
	r[2][2] = 5.0

	pv[0][0] = 0.2
	pv[0][1] = 1.5
	pv[0][2] = 0.1

	pv[1][0] = 1.5
	pv[1][1] = 0.2
	pv[1][2] = 0.1

	tests := []struct {
		ref string
		fn  func([3][3]float64, [2][3]float64) [2][3]float64
	}{
		{"cgo", CgoTrxpv},
		{"go", GoTrxpv},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(r, pv)
			}
		})
	}
}
