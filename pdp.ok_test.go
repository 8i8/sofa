package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t P d p
//  - - - - - - - -
//
//  Test Pdp function.
//
//  Called:  Pdp, vvd
//
//  This revision:  2013 August 7
//
func TestPdp(t *testing.T) {
	const fname = "Pdp"
	var a, b [3]float64

	a[0] = 2.0
	a[1] = 2.0
	a[2] = 3.0

	b[0] = 1.0
	b[1] = 3.0
	b[2] = 4.0

	tests := []struct {
		ref string
		fn  func(a, b [3]float64) float64
	}{
		{"cgo", CgoPdp},
		{"go", GoPdp},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		adb := test.fn(a, b)
		vvd(t, adb, 20, 1e-12, tname, "")
	}
}

func BenchmarkPdp(b *testing.B) {
	var a1, b1 [3]float64

	a1[0] = 2.0
	a1[1] = 2.0
	a1[2] = 3.0

	b1[0] = 1.0
	b1[1] = 3.0
	b1[2] = 4.0

	tests := []struct {
		ref string
		fn  func(c, d [3]float64) float64
	}{
		{"cgo", CgoPdp},
		{"go", GoPdp},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(a1, b1)
			}
		})
	}
}
