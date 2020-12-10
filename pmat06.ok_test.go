package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t P m a t 0 6
//  - - - - - - - - - - -
//
//  Test Pmat06 function.
//
//  Called:  Pmat06, vvd
//
//  This revision:  2013 August 7
//
func TestPmat06(t *testing.T) {
	const fname = "Pmat06"
	var rbp [3][3]float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoPmat06},
		{"go", GoPmat06},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		rbp = test.fn(2400000.5, 50123.9999)

		vvd(t, rbp[0][0], 0.9999995505176007047, 1e-12,
			tname, "11")
		vvd(t, rbp[0][1], 0.8695404617348208406e-3, 1e-14,
			tname, "12")
		vvd(t, rbp[0][2], 0.3779735201865589104e-3, 1e-14,
			tname, "13")

		vvd(t, rbp[1][0], -0.8695404723772031414e-3, 1e-14,
			tname, "21")
		vvd(t, rbp[1][1], 0.9999996219496027161, 1e-12,
			tname, "22")
		vvd(t, rbp[1][2], -0.1361752497080270143e-6, 1e-14,
			tname, "23")

		vvd(t, rbp[2][0], -0.3779734957034089490e-3, 1e-14,
			tname, "31")
		vvd(t, rbp[2][1], -0.1924880847894457113e-6, 1e-14,
			tname, "32")
		vvd(t, rbp[2][2], 0.9999999285679971958, 1e-12,
			tname, "33")
	}
}

func BenchmarkPmat06(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoPmat06},
		{"go", GoPmat06},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 50123.9999)
			}
		})
	}
}
