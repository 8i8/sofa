package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t J d 2 c a l
//  - - - - - - - - - - -
//
//  Test Jd2cal function.
//
//  Called:  Jd2cal, viv, vvd
//
//  This revision:  2013 August 7
//
func TestJd2cal(t *testing.T) {
	const fname = "Jd2cal"
	var dj1, dj2 float64
	dj1 = 2400000.5
	dj2 = 50123.9999
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (b1, b2, b3 int, b4 float64, b5 error)
	}{
		{"cgo", CgoJd2cal},
		{"go", GoJd2cal},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		iy, im, id, fd, err := test.fn(dj1, dj2)

		viv(t, iy, 1996, tname, "y")
		viv(t, im, 2, tname, "m")
		viv(t, id, 10, tname, "d")
		vvd(t, fd, 0.9999, 1e-7, tname, "fd")
		errT(t, nil, err, tname)
	}
}

func BenchmarkJd2cal(b *testing.B) {
	var dj1, dj2 float64
	dj1 = 2400000.5
	dj2 = 50123.9999
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (b1, b2, b3 int, b4 float64, b5 error)
	}{
		{"cgo", CgoJd2cal},
		{"go", GoJd2cal},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(dj1, dj2)
			}
		})
	}
}
