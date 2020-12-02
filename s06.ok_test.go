package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t S 0 6
//  - - - - - - - -
//
//  Test S06 function.
//
//  Called:  S06, vvd
//
//  This revision:  2013 August 7
//
func TestS06(t *testing.T) {
	const fname = "S06"
	var x, y float64
	x = 0.5791308486706011000e-3
	y = 0.4020579816732961219e-4
	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) float64
	}{
		{"cgo", CgoS06},
		{"go", GoS06},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		s := test.fn(2400000.5, 53736.0, x, y)
		vvd(t, s, -0.1220032213076463117e-7, 1e-18, tname, "")
	}
}

func BenchmarkS06(b *testing.B) {
	var x, y float64
	x = 0.5791308486706011000e-3
	y = 0.4020579816732961219e-4
	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) float64
	}{
		{"cgo", CgoS06},
		{"go", GoS06},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736.0, x, y)
			}
		})
	}
}
