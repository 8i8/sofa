package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t R x r
//  - - - - - - - -
//
//  Test iauRxr function.
//
//  Called:  iauRxr, vvd
//
//  This revision:  2013 August 7
//
func TestRxr(t *testing.T) {
	const fname = "Rxr"
	var a, b [3][3]float64

	a[0][0] = 2.0
	a[0][1] = 3.0
	a[0][2] = 2.0

	a[1][0] = 3.0
	a[1][1] = 2.0
	a[1][2] = 3.0

	a[2][0] = 3.0
	a[2][1] = 4.0
	a[2][2] = 5.0

	b[0][0] = 1.0
	b[0][1] = 2.0
	b[0][2] = 2.0

	b[1][0] = 4.0
	b[1][1] = 1.0
	b[1][2] = 1.0

	b[2][0] = 3.0
	b[2][1] = 0.0
	b[2][2] = 1.0

	tests := []struct {
		ref string
		fn  func(c, d [3][3]float64) [3][3]float64
	}{
		{"cgo", CgoRxr},
		{"go", GoRxr},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		atb := test.fn(a, b)

		vvd(t, atb[0][0], 20.0, 1e-12, tname, "11")
		vvd(t, atb[0][1], 7.0, 1e-12, tname, "12")
		vvd(t, atb[0][2], 9.0, 1e-12, tname, "13")

		vvd(t, atb[1][0], 20.0, 1e-12, tname, "21")
		vvd(t, atb[1][1], 8.0, 1e-12, tname, "22")
		vvd(t, atb[1][2], 11.0, 1e-12, tname, "23")

		vvd(t, atb[2][0], 34.0, 1e-12, tname, "31")
		vvd(t, atb[2][1], 10.0, 1e-12, tname, "32")
		vvd(t, atb[2][2], 15.0, 1e-12, tname, "33")
	}
}

func BenchmarkRxr(b *testing.B) {
	var a1, b1 [3][3]float64

	a1[0][0] = 2.0
	a1[0][1] = 3.0
	a1[0][2] = 2.0

	a1[1][0] = 3.0
	a1[1][1] = 2.0
	a1[1][2] = 3.0

	a1[2][0] = 3.0
	a1[2][1] = 4.0
	a1[2][2] = 5.0

	b1[0][0] = 1.0
	b1[0][1] = 2.0
	b1[0][2] = 2.0

	b1[1][0] = 4.0
	b1[1][1] = 1.0
	b1[1][2] = 1.0

	b1[2][0] = 3.0
	b1[2][1] = 0.0
	b1[2][2] = 1.0

	tests := []struct {
		ref string
		fn  func(c, d [3][3]float64) [3][3]float64
	}{
		{"cgo", CgoRxr},
		{"go", GoRxr},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(a1, b1)
			}
		})
	}
}
