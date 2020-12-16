package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t L t p e q u
//  - - - - - - - - - - -
//
//  Test Ltpequ function.
//
//  Called:  Ltpequ, vvd
//
//  This revision:  2016 March 12
//
func TestLtpequ(t *testing.T) {
	const fname = "Ltpequ"
	var veq [3]float64
	var epj float64

	epj = -2500.0

	tests := []struct {
		ref string
		fn  func(float64) [3]float64
	}{
		{"cgo", CgoLtpequ},
		{"go", GoLtpequ},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		veq = test.fn(epj)

		vvd(t, veq[0], -0.3586652560237326659, 1e-14,
			tname, "veq1")
		vvd(t, veq[1], -0.1996978910771128475, 1e-14,
			tname, "veq2")
		vvd(t, veq[2], 0.9118552442250819624, 1e-14,
			tname, "veq3")
	}
}

func BenchmarkLtpequ(b *testing.B) {
	var epj float64

	epj = -2500.0

	tests := []struct {
		ref string
		fn  func(float64) [3]float64
	}{
		{"cgo", CgoLtpequ},
		{"go", GoLtpequ},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(epj)
			}
		})
	}
}
