package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t A e 2 h d
//  - - - - - - - - - -
//
//  Test Ae2hd function.
//
//  Called:  iauAe2hd and vvd
//
//  This revision:  2017 October 21
//
func TestAe2hd(t *testing.T) {
	const fname = "Ae2hd"
	var a, e, p, h, d float64

	a = 5.5
	e = 1.1
	p = 0.7

	tests := []struct {
		ref string
		fn  func(a, b, c float64) (e, f float64)
	}{
		{"cgo", CgoAe2hd},
		{"go", GoAe2hd},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		h, d = test.fn(a, e, p)

		vvd(t, h, 0.5933291115507309663, 1e-14, tname, "h")
		vvd(t, d, 0.9613934761647817620, 1e-14, tname, "d")
	}
}

func BenchmarkAe2hd(b *testing.B) {
	var a, e, p float64
	a = 5.5
	e = 1.1
	p = 0.7
	tests := []struct {
		ref string
		fn  func(a, b, c float64) (e, f float64)
	}{
		{"cgo", CgoAe2hd},
		{"go", GoAe2hd},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(a, e, p)
			}
		})
	}
}
