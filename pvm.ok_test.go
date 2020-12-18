package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t P v m
//  - - - - - - - -
//
//  Test Pvm function.
//
//  Called:  Pvm, vvd
//
//  This revision:  2013 August 7
//
func TestPvm(t *testing.T) {
	const fname = "Pvm"
	var r, s float64
	var pv [2][3]float64

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = 0.45
	pv[1][1] = -0.25
	pv[1][2] = 1.1

	tests := []struct {
		ref string
		fn  func([2][3]float64) (c1,c2 float64)
	}{
		{"cgo", CgoPvm},
		{"go", GoPvm},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		r, s = test.fn(pv)

		vvd(t, r, 2.789265136196270604, 1e-12, tname, "r")
		vvd(t, s, 1.214495780149111922, 1e-12, tname, "s")
	}
}

func BenchmarkPvm(b *testing.B) {
	var pv [2][3]float64

	pv[0][0] = 0.3
	pv[0][1] = 1.2
	pv[0][2] = -2.5

	pv[1][0] = 0.45
	pv[1][1] = -0.25
	pv[1][2] = 1.1

	tests := []struct {
		ref string
		fn  func([2][3]float64) (c1,c2 float64)
	}{
		{"cgo", CgoPvm},
		{"go", GoPvm},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(pv)
			}
		})
	}
}
