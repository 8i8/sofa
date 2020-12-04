package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t R x p
//  - - - - - - - -
//
//  Test Rxp function.
//
//  Returned:
//     status    int         FALSE = success, TRUE = fail
//
//  Called:  Rxp, vvd
//
//  This revision:  2013 August 7
//
func TestRxp(t *testing.T) {
	const fname = "Rxp"
	var r [3][3]float64
	var p [3]float64

	r[0][0] = 2.0
	r[0][1] = 3.0
	r[0][2] = 2.0

	r[1][0] = 3.0
	r[1][1] = 2.0
	r[1][2] = 3.0

	r[2][0] = 3.0
	r[2][1] = 4.0
	r[2][2] = 5.0

	p[0] = 0.2
	p[1] = 1.5
	p[2] = 0.1

	tests := []struct {
		ref string
		fn  func([3][3]float64, [3]float64) [3]float64
	}{
		{"cgo", CgoRxp},
		{"go", GoRxp},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		rp := test.fn(r, p)

		vvd(t, rp[0], 5.1, 1e-12, tname, "1")
		vvd(t, rp[1], 3.9, 1e-12, tname, "2")
		vvd(t, rp[2], 7.1, 1e-12, tname, "3")
	}
}

func BenchmarkRxp(b *testing.B) {
	var r [3][3]float64
	var p [3]float64

	r[0][0] = 2.0
	r[0][1] = 3.0
	r[0][2] = 2.0

	r[1][0] = 3.0
	r[1][1] = 2.0
	r[1][2] = 3.0

	r[2][0] = 3.0
	r[2][1] = 4.0
	r[2][2] = 5.0

	p[0] = 0.2
	p[1] = 1.5
	p[2] = 0.1

	tests := []struct {
		ref string
		fn  func([3][3]float64, [3]float64) [3]float64
	}{
		{"cgo", CgoRxp},
		{"go", GoRxp},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(r, p)
			}
		})
	}
}
