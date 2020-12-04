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
func TestRxpv(t *testing.T) {
	const fname = "Rxpv"
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
		{"cgo", CgoRxpv},
		{"go", GoRxpv},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		rpv := test.fn(r, pv)

		vvd(t, rpv[0][0], 5.1, 1e-12, tname, "11")
		vvd(t, rpv[1][0], 3.8, 1e-12, tname, "12")

		vvd(t, rpv[0][1], 3.9, 1e-12, tname, "21")
		vvd(t, rpv[1][1], 5.2, 1e-12, tname, "22")

		vvd(t, rpv[0][2], 7.1, 1e-12, tname, "31")
		vvd(t, rpv[1][2], 5.8, 1e-12, tname, "32")
	}
}

func BenchmarkRxpv(b *testing.B) {
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
		{"cgo", CgoRxpv},
		{"go", GoRxpv},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(r, pv)
			}
		})
	}
}
