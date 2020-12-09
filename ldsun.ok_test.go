package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t L d s u n
//  - - - - - - - - - -
//
//  Test Ldsun function.
//
//  Called:  Ldsun, vvd
//
//  This revision:  2013 October 2
//
func TestLdsun(t *testing.T) {
	const fname = "Ldsun"
	var p, e [3]float64
	var em float64

	p[0] = -0.763276255
	p[1] = -0.608633767
	p[2] = -0.216735543
	e[0] = -0.973644023
	e[1] = -0.20925523
	e[2] = -0.0907169552
	em = 0.999809214

	tests := []struct {
		ref string
		fn  func(a1, a2 [3]float64, a3 float64) [3]float64
	}{
		{"cgo", CgoLdsun},
		{"go", GoLdsun},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		p1 := test.fn(p, e, em)

		vvd(t, p1[0], -0.7632762580731413169, 1e-12,
			tname, "1")
		vvd(t, p1[1], -0.6086337635262647900, 1e-12,
			tname, "2")
		vvd(t, p1[2], -0.2167355419322321302, 1e-12,
			tname, "3")
	}
}

func BenchmarkLdsun(b *testing.B) {
	var p, e [3]float64
	var em float64

	p[0] = -0.763276255
	p[1] = -0.608633767
	p[2] = -0.216735543
	e[0] = -0.973644023
	e[1] = -0.20925523
	e[2] = -0.0907169552
	em = 0.999809214

	tests := []struct {
		ref string
		fn  func(a1, a2 [3]float64, a3 float64) [3]float64
	}{
		{"cgo", CgoLdsun},
		{"go", GoLdsun},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(p, e, em)
			}
		})
	}
}
