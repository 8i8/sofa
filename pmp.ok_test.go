package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t P m p
//  - - - - - - - -
//
//  Test Pmp function.
//
//  Called:  Pmp, vvd
//
//  This revision:  2013 August 7
//
func TestPmp(t *testing.T) {
	const fname = "Pmp"
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
		{"cgo", CgoPmp},
		{"go", GoPmp},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		amb := test.fn(a, b)

		vvd(t, amb[0], 1.0, 1e-12, tname, "0")
		vvd(t, amb[1], -1.0, 1e-12, tname, "1")
		vvd(t, amb[2], -1.0, 1e-12, tname, "2")
	}
}

func BenchmarkPmp(b *testing.B) {
	var c, d [3]float64

	c[0] = 2.0
	c[1] = 2.0
	c[2] = 3.0

	d[0] = 1.0
	d[1] = 3.0
	d[2] = 4.0

	tests := []struct {
		ref string
		fn  func(a1, a2 [3]float64) [3]float64
	}{
		{"cgo", CgoPmp},
		{"go", GoPmp},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(c, d)
			}
		})
	}
}
