package sofa

import (
	"testing"
)

//
//  - - - - - - - - - -
//   T e s t G d 2 g c
//  - - - - - - - - - -
//
//  Test Gd2gc function.
//
//  Called:  Gd2gc, viv, vvd
//
//  This revision:  2016 March 12
//
func TestGd2gc(t *testing.T) {
	const fname = "Gd2gc"
	var e, p, h float64

	e = 3.1
	p = -0.5
	h = 2500.0

	tests := []struct {
		ref string
		fn  func(a1 int, a2, a3, a4 float64) ([3]float64, error)
	}{
		{"cgo", CgoGd2gc},
		{"go", GoGd2gc},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref

		xyz, err := test.fn(int(0), e, p, h)
		errT(t, errGd2gcE1, err, tname, "-1")

		xyz, err = test.fn(WGS84, e, p, h)

		errT(t, nil, err, tname, "0")
		vvd(t, xyz[0], -5599000.5577049947, 1e-7, tname, "1/1")
		vvd(t, xyz[1], 233011.67223479203, 1e-7, tname, "2/1")
		vvd(t, xyz[2], -3040909.4706983363, 1e-7, tname, "3/1")

		xyz, err = test.fn(GRS80, e, p, h)

		errT(t, nil, err, tname, "0")
		vvd(t, xyz[0], -5599000.5577260984, 1e-7, tname, "1/2")
		vvd(t, xyz[1], 233011.6722356702949, 1e-7, tname, "2/2")
		vvd(t, xyz[2], -3040909.4706095476, 1e-7, tname, "3/2")

		xyz, err = test.fn(WGS72, e, p, h)

		errT(t, nil, err, tname, "0")
		vvd(t, xyz[0], -5598998.7626301490, 1e-7, tname, "1/3")
		vvd(t, xyz[1], 233011.5975297822211, 1e-7, tname, "2/3")
		vvd(t, xyz[2], -3040908.6861467111, 1e-7, tname, "3/3")

		xyz, err = test.fn(4, e, p, h)
		errT(t, errGd2gcE1, err, tname, "-1")
	}
}

func BenchmarkGd2gc(b *testing.B) {
	var e, p, h float64
	e = 3.1
	p = -0.5
	h = 2500.0

	tests := []struct {
		ref string
		fn  func(a1 int, a2, a3, a4 float64) ([3]float64, error)
	}{
		{"cgo", CgoGd2gc},
		{"go", GoGd2gc},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(int(0), e, p, h)
			}
		})
	}

}
