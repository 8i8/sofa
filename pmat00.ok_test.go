package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t P m a t 0 0
//  - - - - - - - - - - -
//
//  Test Pmat00 function.
//
//  Called:  Pmat00, vvd
//
//  This revision:  2013 August 7
//
func TestPmat00(t *testing.T) {
	const fname = "Pmat00"
	var rbp [3][3]float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoPmat00},
		{"go", GoPmat00},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rbp = test.fn(2400000.5, 50123.9999)

		vvd(t, rbp[0][0], 0.9999995505175087260, 1e-12,
			tname, "11")
		vvd(t, rbp[0][1], 0.8695405883617884705e-3, 1e-14,
			tname, "12")
		vvd(t, rbp[0][2], 0.3779734722239007105e-3, 1e-14,
			tname, "13")

		vvd(t, rbp[1][0], -0.8695405990410863719e-3, 1e-14,
			tname, "21")
		vvd(t, rbp[1][1], 0.9999996219494925900, 1e-12,
			tname, "22")
		vvd(t, rbp[1][2], -0.1360775820404982209e-6, 1e-14,
			tname, "23")

		vvd(t, rbp[2][0], -0.3779734476558184991e-3, 1e-14,
			tname, "31")
		vvd(t, rbp[2][1], -0.1925857585832024058e-6, 1e-14,
			tname, "32")
		vvd(t, rbp[2][2], 0.9999999285680153377, 1e-12,
			tname, "33")
	}
}

func BenchmarkPmat00(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoPmat00},
		{"go", GoPmat00},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 50123.9999)
			}
		})
	}
}
