package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t P v d p v
//  - - - - - - - - - -
//
//  Test Pvdpv function.
//
//  Called:  Pvdpv, vvd
//
//  This revision:  2013 August 7
//
func TestPvdpv(t *testing.T) {
	const fname = "Pvdpv"
	var adb [2]float64
	var a, b [2][3]float64

	a[0][0] = 2.0
	a[0][1] = 2.0
	a[0][2] = 3.0

	a[1][0] = 6.0
	a[1][1] = 0.0
	a[1][2] = 4.0

	b[0][0] = 1.0
	b[0][1] = 3.0
	b[0][2] = 4.0

	b[1][0] = 0.0
	b[1][1] = 2.0
	b[1][2] = 8.0

	tests := []struct {
		ref string
		fn  func(a1, a2 [2][3]float64) [2]float64
	}{
		{"cgo", CgoPvdpv},
		{"go", GoPvdpv},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		adb = test.fn(a, b)

		vvd(t, adb[0], 20.0, 1e-12, tname, "1")
		vvd(t, adb[1], 50.0, 1e-12, tname, "2")
	}
}

func BenchmarkPvdpv(bm *testing.B) {
	var a, b [2][3]float64

	a[0][0] = 2.0
	a[0][1] = 2.0
	a[0][2] = 3.0

	a[1][0] = 6.0
	a[1][1] = 0.0
	a[1][2] = 4.0

	b[0][0] = 1.0
	b[0][1] = 3.0
	b[0][2] = 4.0

	b[1][0] = 0.0
	b[1][1] = 2.0
	b[1][2] = 8.0

	tests := []struct {
		ref string
		fn  func(a1, a2 [2][3]float64) [2]float64
	}{
		{"cgo", CgoPvdpv},
		{"go", GoPvdpv},
	}

	for _, test := range tests {
		bm.Run(test.ref, func(bm *testing.B) {
			for i := 0; i < bm.N; i++ {
				_ = test.fn(a, b)
			}
		})
	}
}
