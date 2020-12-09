package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t C 2 s
//  - - - - - - - -
//
//  Test C2s function.
//
//  Called:  C2s, vvd
//
//  This revision:  2013 August 7
//
func TestC2s(t *testing.T) {
	const fname = "C2s"
	var p [3]float64

	p[0] = 100.0
	p[1] = -50.0
	p[2] = 25.0

	tests := []struct {
		ref string
		fn  func([3]float64) (b1, b2 float64)
	}{
		{"cgo", CgoC2s},
		{"go", GoC2s},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		theta, phi := test.fn(p)

		vvd(t, theta, -0.4636476090008061162, 1e-14, tname, "theta")
		vvd(t, phi, 0.2199879773954594463, 1e-14, tname, "phi")
	}
}

func BenchmarkC2s(b *testing.B) {
	var p [3]float64

	p[0] = 100.0
	p[1] = -50.0
	p[2] = 25.0

	tests := []struct {
		ref string
		fn  func([3]float64) (b1, b2 float64)
	}{
		{"cgo", CgoC2s},
		{"go", GoC2s},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(p)
			}
		})
	}
}
