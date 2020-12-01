package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t S x p
//  - - - - - - - -
//
//  Test Sxp function.
//
//  Called:  Sxp, vvd
//
//  This revision:  2013 August 7
//
func TestSxp(t *testing.T) {
	const fname = "Sxp"
	var s float64
	var p [3]float64

	s = 2.0
	p[0] = 0.3
	p[1] = 1.2
	p[2] = -2.5

	tests := []struct {
		ref string
		fn  func(float64, [3]float64) [3]float64
	}{
		{"cgo", CgoSxp},
		{"go", GoSxp},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref

		sp := test.fn(s, p)

		vvd(t, sp[0], 0.6, 0.0, tname, "1")
		vvd(t, sp[1], 2.4, 0.0, tname, "2")
		vvd(t, sp[2], -5.0, 0.0, tname, "3")
	}
}

func BenchmarkSxp(b *testing.B) {
	var s float64
	var p [3]float64

	s = 2.0
	p[0] = 0.3
	p[1] = 1.2
	p[2] = -2.5

	tests := []struct {
		ref string
		fn  func(float64, [3]float64) [3]float64
	}{
		{"cgo", CgoSxp},
		{"go", GoSxp},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(s, p)
			}
		})
	}
}
