package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t P p s p
//  - - - - - - - - -
//
//  Test Ppsp function.
//
//  Called:  Ppsp, vvd
//
//  This revision:  2013 August 7
//
func TestPpsp(t *testing.T) {
	const fname = "Ppsp"
	var a, d [3]float64
	var s float64

	a[0] = 2.0
	a[1] = 2.0
	a[2] = 3.0

	s = 5.0

	d[0] = 1.0
	d[1] = 3.0
	d[2] = 4.0

	tests := []struct {
		ref string
		fn  func([3]float64, float64, [3]float64) [3]float64
	}{
		{"cgo", CgoPpsp},
		{"go", GoPpsp},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		apsb := test.fn(a, s, d)

		vvd(t, apsb[0], 7.0, 1e-12, tname, "0")
		vvd(t, apsb[1], 17.0, 1e-12, tname, "1")
		vvd(t, apsb[2], 23.0, 1e-12, tname, "2")
	}
}

func BenchmarkPpsp(b *testing.B) {
	var a, d [3]float64
	var s float64

	a[0] = 2.0
	a[1] = 2.0
	a[2] = 3.0

	s = 5.0

	d[0] = 1.0
	d[1] = 3.0
	d[2] = 4.0

	tests := []struct {
		ref string
		fn  func([3]float64, float64, [3]float64) [3]float64
	}{
		{"cgo", CgoPpsp},
		{"go", GoPpsp},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(a, s, d)
			}
		})
	}
}
