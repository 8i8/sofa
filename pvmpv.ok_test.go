package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t P v m p v
//  - - - - - - - - - -
//
//  Test Pvmpv function.
//
//  Called:  Pvmpv, vvd
//
//  This revision:  2013 August 7
//
func TestPvmpv(t *testing.T) {
	const fname = "Pvmpv"
	var a, b, amb [2][3]float64

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
		{"cgo", CgoPvmpv},
		{"go", GoPvmpv},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		amb = test.fn(a, b)

		vvd(t, amb[0][0], 1.0, 1e-12, tname, "11")
		vvd(t, amb[0][1], -1.0, 1e-12, tname, "21")
		vvd(t, amb[0][2], -1.0, 1e-12, tname, "31")

		vvd(t, amb[1][0], 2.0, 1e-12, tname, "12")
		vvd(t, amb[1][1], 4.0, 1e-12, tname, "22")
		vvd(t, amb[1][2], 2.0, 1e-12, tname, "32")
	}
}

func BenchmarkPvmpv(b *testing.B) {
	var a, c [2][3]float64

	a[0][0] = 2.0
	a[0][1] = 2.0
	a[0][2] = 3.0

	a[1][0] = 5.0
	a[1][1] = 6.0
	a[1][2] = 3.0

	c[0][0] = 1.0
	c[0][1] = 3.0
	c[0][2] = 4.0

	c[1][0] = 3.0
	c[1][1] = 2.0
	c[1][2] = 1.0

	tests := []struct {
		ref string
		fn  func(a1, a2 [2][3]float64) [2][3]float64
	}{
		{"cgo", CgoPvmpv},
		{"go", GoPvmpv},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(a, c)
			}
		})
	}
}
