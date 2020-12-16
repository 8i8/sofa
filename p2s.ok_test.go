package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t P 2 s
//  - - - - - - - -
//
//  Test P2s function.
//
//  Called:  P2s, vvd
//
//  This revision:  2013 August 7
//
func TestP2s(t *testing.T) {
	const fname = "P2s"
	var theta, phi, r float64
	var p [3]float64

	p[0] = 100.0
	p[1] = -50.0
	p[2] = 25.0

	tests := []struct {
		ref string
		fn  func(p [3]float64) (c1, c2, c3 float64)
	}{
		{"cgo", CgoP2s},
		{"go", GoP2s},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		theta, phi, r = test.fn(p)

		vvd(t, theta, -0.4636476090008061162, 1e-12, tname, "theta")
		vvd(t, phi, 0.2199879773954594463, 1e-12, tname, "phi")
		vvd(t, r, 114.5643923738960002, 1e-9, tname, "r")
	}
}

func BenchmarkP2s(b *testing.B) {
	var p [3]float64

	p[0] = 100.0
	p[1] = -50.0
	p[2] = 25.0

	tests := []struct {
		ref string
		fn  func(p [3]float64) (c1, c2, c3 float64)
	}{
		{"cgo", CgoP2s},
		{"go", GoP2s},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(p)
			}
		})
	}
}
