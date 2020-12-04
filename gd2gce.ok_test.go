package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t G d 2 g c e
//  - - - - - - - - - - -
//
//  Test Gd2gce function.
//
//  Called:  Gd2gce, viv, vvd
//
//  This revision:  2016 March 12
//
func TestGd2gce(t *testing.T) {
	const fname = "Gd2gce"
	var a, f, e, p, h float64
	a = 6378136.0
	f = 0.0033528
	e = 3.1
	p = -0.5
	h = 2500.0
	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5 float64) ([3]float64, error)
	}{
		{"cgo", CgoGd2gce},
		{"go", GoGd2gce},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref

		xyz, err := test.fn(a, f, e, p, h)

		errT(t, nil, err, tname+" 0")
		vvd(t, xyz[0], -5598999.6665116328, 1e-7, tname, "1")
		vvd(t, xyz[1], 233011.6351463057189, 1e-7, tname, "2")
		vvd(t, xyz[2], -3040909.0517314132, 1e-7, tname, "3")
	}
}

func BenchmarkGd2gce(b *testing.B) {
	var a, f, e, p, h float64
	a = 6378136.0
	f = 0.0033528
	e = 3.1
	p = -0.5
	h = 2500.0
	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5 float64) ([3]float64, error)
	}{
		{"cgo", CgoGd2gce},
		{"go", GoGd2gce},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(a, f, e, p, h)
			}
		})
	}
}
