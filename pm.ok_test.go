package sofa

import "testing"

//
//  - - - - - - -
//   T e s t P m
//  - - - - - - -
//
//  Test Pm function.
//
//  Called:  Pm, vvd
//
//  This revision:  2013 August 7
//
func TestPm(t *testing.T) {
	const fname = "Pm"
	var p [3]float64
	p[0] = 0.3
	p[1] = 1.2
	p[2] = -2.5
	tests := []struct {
		ref string
		fn  func([3]float64) float64
	}{
		{"cgo", CgoPm},
		{"go", GoPm},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		r := test.fn(p)
		vvd(t, r, 2.789265136196270604, 1e-12, tname, "")
	}
}

func BenchmarkPm(b *testing.B) {
	var p [3]float64
	p[0] = 0.3
	p[1] = 1.2
	p[2] = -2.5
	tests := []struct {
		ref string
		fn  func([3]float64) float64
	}{
		{"cgo", CgoPm},
		{"go", GoPm},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(p)
			}
		})
	}
}
