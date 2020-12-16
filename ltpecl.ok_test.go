package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t L t p e c l
//  - - - - - - - - - - -
//
//  Test Ltpecl function.
//
//  Called:  Ltpecl, vvd
//
//  This revision:  2016 March 12
//
func TestLtpecl(t *testing.T) {
	const fname = "Ltpecl"
	var vec [3]float64
	var epj float64

	epj = -1500.0

	tests := []struct {
		ref string
		fn  func(float64) [3]float64
	}{
		{"cgo", CgoLtpecl},
		{"go", GoLtpecl},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		vec = test.fn(epj)

		vvd(t, vec[0], 0.4768625676477096525e-3, 1e-14,
			tname, "vec1")
		vvd(t, vec[1], -0.4052259533091875112, 1e-14,
			tname, "vec2")
		vvd(t, vec[2], 0.9142164401096448012, 1e-14,
			tname, "vec3")
	}
}

func BenchmarkLtpecl(b *testing.B) {
	var epj float64

	epj = -1500.0

	tests := []struct {
		ref string
		fn  func(float64) [3]float64
	}{
		{"cgo", CgoLtpecl},
		{"go", GoLtpecl},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(epj)
			}
		})
	}
}
