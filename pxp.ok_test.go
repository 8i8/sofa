package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t P x p
//  - - - - - - - -
//
//  Test Pxp function.
//
//  Called:  Pxp, vvd
//
//  This revision:  2013 August 7
//
func TestPxp(t *testing.T) {
	const fname = "Pxp"
	var a, b [3]float64

	a[0] = 2.0
	a[1] = 2.0
	a[2] = 3.0

	b[0] = 1.0
	b[1] = 3.0
	b[2] = 4.0

	tests := []struct {
		ref string
		fn  func(a1, a2 [3]float64) [3]float64
	}{
		{"cgo", CgoPxp},
		{"go", GoPxp},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		axb := test.fn(a, b)

		vvd(t, axb[0], -1.0, 1e-12, tname, "1")
		vvd(t, axb[1], -5.0, 1e-12, tname, "2")
		vvd(t, axb[2], 4.0, 1e-12, tname, "3")
	}
}

func BenchmarkPxp(b *testing.B) {
	var c, d [3]float64

	c[0] = 2.0
	c[1] = 2.0
	c[2] = 3.0

	c[0] = 1.0
	c[1] = 3.0
	c[2] = 4.0

	tests := []struct {
		ref string
		fn  func(a1, a2 [3]float64) [3]float64
	}{
		{"cgo", CgoPxp},
		{"go", GoPxp},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(c, d)
			}
		})
	}
}
