package sofa

import "testing"

//
//  - - - - - - -
//   T e s t P n
//  - - - - - - -
//
//  Test Pn function.
//
//  Called:  Pn, vvd
//
//  This revision:  2013 August 7
//
func TestPn(t *testing.T) {
	const fname = "Pn"
	var p, u [3]float64
	var r float64

	p[0] = 0.3
	p[1] = 1.2
	p[2] = -2.5

	tests := []struct {
		ref string
		fn  func([3]float64) (float64, [3]float64)
	}{
		{"cgo", CgoPn},
		{"go", GoPn},
	}

	for _, test := range tests {

		tname := fname + " " + test.ref
		r, u = test.fn(p)

		vvd(t, r, 2.789265136196270604, 1e-12, tname, "r")

		vvd(t, u[0], 0.1075552109073112058, 1e-12, tname, "u1")
		vvd(t, u[1], 0.4302208436292448232, 1e-12, tname, "u2")
		vvd(t, u[2], -0.8962934242275933816, 1e-12, tname, "u3")
	}
}

func BenchmarkPn(b *testing.B) {
	var p [3]float64
	p[0] = 0.3
	p[1] = 1.2
	p[2] = -2.5
	tests := []struct {
		ref string
		fn  func([3]float64) (float64, [3]float64)
	}{
		{"cgo", CgoPn},
		{"go", GoPn},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(p)
			}
		})
	}
}
