package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t P b 0 6
//  - - - - - - - - -
//
//  Test Pb06 function.
//
//  Called:  Pb06, vvd
//
//  This revision:  2013 August 7
//
func TestPb06(t *testing.T) {
	const fname = "Pb06"
	var bzeta, bz, btheta float64

	tests := []struct {
		ref string
		fn  func(a1,a2 float64) (c1,c2,c3 float64)
	}{
		{"cgo", CgoPb06},
		{"go", GoPb06},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		bzeta, bz, btheta = test.fn(2400000.5, 50123.9999)

		vvd(t, bzeta, -0.5092634016326478238e-3, 1e-12,
			tname, "bzeta")
		vvd(t, bz, -0.3602772060566044413e-3, 1e-12,
			tname, "bz")
		vvd(t, btheta, -0.3779735537167811177e-3, 1e-12,
			tname, "btheta")
	}
}

func BenchmarkPb06(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1,a2 float64) (c1,c2,c3 float64)
	}{
		{"cgo", CgoPb06},
		{"go", GoPb06},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(2400000.5, 50123.9999)
			}
		})
	}
}
