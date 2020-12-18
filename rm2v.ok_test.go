package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t R m 2 v
//  - - - - - - - - -
//
//  Test Rm2v function.
//
//  Called:  Rm2v, vvd
//
//  This revision:  2013 August 7
//
func TestRm2v(t *testing.T) {
	const fname = "Rm2v"
	var w [3]float64
	var r [3][3]float64

	r[0][0] = 0.00
	r[0][1] = -0.80
	r[0][2] = -0.60

	r[1][0] = 0.80
	r[1][1] = -0.36
	r[1][2] = 0.48

	r[2][0] = 0.60
	r[2][1] = 0.48
	r[2][2] = -0.64

	tests := []struct {
		ref string
		fn  func([3][3]float64) [3]float64
	}{
		{"cgo", CgoRm2v},
		{"go", GoRm2v},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		w = test.fn(r)

		vvd(t, w[0], 0.0, 1e-12, tname, "1")
		vvd(t, w[1], 1.413716694115406957, 1e-12, tname, "2")
		vvd(t, w[2], -1.884955592153875943, 1e-12, tname, "3")
	}
}

func BenchmarkRm2v(b *testing.B) {
	var r [3][3]float64

	r[0][0] = 0.00
	r[0][1] = -0.80
	r[0][2] = -0.60

	r[1][0] = 0.80
	r[1][1] = -0.36
	r[1][2] = 0.48

	r[2][0] = 0.60
	r[2][1] = 0.48
	r[2][2] = -0.64

	tests := []struct {
		ref string
		fn  func([3][3]float64) [3]float64
	}{
		{"cgo", CgoRm2v},
		{"go", GoRm2v},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(r)
			}
		})
	}
}
