package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t F k 5 2 h
//  - - - - - - - - - -
//
//  Test Fk52h function.
//
//  Called:  Fk52h, vvd
//
//  This revision:  2017 January 3
//
func TestFk52h(t *testing.T) {
	const fname = "Fk52h"
	var r5, d5, dr5, dd5, px5, rv5,
		rh, dh, drh, ddh, pxh, rvh float64

	r5 = 1.76779433
	d5 = -0.2917517103
	dr5 = -1.91851572e-7
	dd5 = -5.8468475e-6
	px5 = 0.379210
	rv5 = -7.6

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6 float64) (
			c1, c2, c3, c4, c5, c6 float64)
	}{
		{"cgo", CgoFk52h},
		{"go", GoFk52h},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rh, dh, drh, ddh, pxh, rvh = test.fn(r5, d5, dr5, dd5,
			px5, rv5)

		vvd(t, rh, 1.767794226299947632, 1e-14, tname, "ra")
		vvd(t, dh, -0.2917516070530391757, 1e-14, tname, "dec")
		vvd(t, drh, -0.19618741256057224e-6, 1e-19,
			tname, "dr5")
		vvd(t, ddh, -0.58459905176693911e-5, 1e-19,
			tname, "dd5")
		vvd(t, pxh, 0.37921, 1e-14, tname, "px")
		vvd(t, rvh, -7.6000000940000254, 1e-11, tname, "rv")
	}
}

func BenchmarkFk52h(b *testing.B) {
	var r5, d5, dr5, dd5, px5, rv5 float64

	r5 = 1.76779433
	d5 = -0.2917517103
	dr5 = -1.91851572e-7
	dd5 = -5.8468475e-6
	px5 = 0.379210
	rv5 = -7.6

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6 float64) (
			c1, c2, c3, c4, c5, c6 float64)
	}{
		{"cgo", CgoFk52h},
		{"go", GoFk52h},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(r5, d5, dr5, dd5, px5, rv5)
			}
		})
	}
}
