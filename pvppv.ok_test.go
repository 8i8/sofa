package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t P v p p v
//  - - - - - - - -
//
//  Test Pvppv function.
//
//  Called:  Pvppv, vvd
//
//  This revision:  2013 August 7
//
func TestPvppv(t *testing.T) {
	const fname = "Pvppv"
	var a, b, apb [2][3]float64

	a[0][0] = 2.0
	a[0][1] = 2.0
	a[0][2] = 3.0

	a[1][0] = 5.0
	a[1][1] = 6.0
	a[1][2] = 3.0

	b[0][0] = 1.0
	b[0][1] = 3.0
	b[0][2] = 4.0

	b[1][0] = 3.0
	b[1][1] = 2.0
	b[1][2] = 1.0

	tests := []struct {
		ref string
		fn  func(a1, a2 [2][3]float64) [2][3]float64
	}{
		{"cgo", CgoPvppv},
		{"go", GoPvppv},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		apb = test.fn(a, b)

		vvd(t, apb[0][0], 3.0, 1e-12, tname, "p1")
		vvd(t, apb[0][1], 5.0, 1e-12, tname, "p2")
		vvd(t, apb[0][2], 7.0, 1e-12, tname, "p3")

		vvd(t, apb[1][0], 8.0, 1e-12, tname, "v1")
		vvd(t, apb[1][1], 8.0, 1e-12, tname, "v2")
		vvd(t, apb[1][2], 4.0, 1e-12, tname, "v3")
	}
}

func BenchmarkPvppv(bm *testing.B) {
	var a, b [2][3]float64

	a[0][0] = 2.0
	a[0][1] = 2.0
	a[0][2] = 3.0

	a[1][0] = 5.0
	a[1][1] = 6.0
	a[1][2] = 3.0

	b[0][0] = 1.0
	b[0][1] = 3.0
	b[0][2] = 4.0

	b[1][0] = 3.0
	b[1][1] = 2.0
	b[1][2] = 1.0

	tests := []struct {
		ref string
		fn  func(a1, a2 [2][3]float64) [2][3]float64
	}{
		{"cgo", CgoPvppv},
		{"go", GoPvppv},
	}

	for _, test := range tests {
		bm.Run(test.ref, func(bm *testing.B) {
			for i := 0; i < bm.N; i++ {
				_ = test.fn(a, b)
			}
		})
	}
}
