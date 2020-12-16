package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - -
//   T e s t G c 2 g d
//  - - - - - - - - - -
//
//  Test Gc2gd function.
//
//  Called:  Gc2gd, viv, vvd
//
//  This revision:  2016 March 12
//
func TestGc2gd(t *testing.T) {
	const fname = "Gc2gd"
	var err en.ErrNum
	var e, p, h float64
	var xyz = [...]float64{2e6, 3e6, 5.244e6}

	tests := []struct {
		ref string
		fn  func(int, [3]float64) (
			c1, c2, c3 float64, c4 en.ErrNum)
	}{
		{"cgo", CgoGc2gd},
		{"go", GoGc2gd},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		e, p, h, err = test.fn(0, xyz)

		errEN(t, -1, err, tname, "j0")

		e, p, h, err = test.fn(WGS84, xyz)

		errT(t, nil, err, tname, "j1")
		vvd(t, e, 0.9827937232473290680, 1e-14, tname, "e1")
		vvd(t, p, 0.97160184819075459, 1e-14, tname, "p1")
		vvd(t, h, 331.4172461426059892, 1e-8, tname, "h1")

		e, p, h, err = test.fn(GRS80, xyz)

		errT(t, nil, err, tname, "j2")
		vvd(t, e, 0.9827937232473290680, 1e-14, tname, "e2")
		vvd(t, p, 0.97160184820607853, 1e-14, tname, "p2")
		vvd(t, h, 331.41731754844348, 1e-8, tname, "h2")

		e, p, h, err = test.fn(WGS72, xyz)

		errT(t, nil, err, tname, "j3")
		vvd(t, e, 0.9827937232473290680, 1e-14, tname, "e3")
		vvd(t, p, 0.9716018181101511937, 1e-14, tname, "p3")
		vvd(t, h, 333.2770726130318123, 1e-8, tname, "h3")

		e, p, h, err = test.fn(4, xyz)

		errEN(t, -1, err, tname, "j4")
	}
}

func BenchmarkGc2gd(b *testing.B) {
	var xyz = [...]float64{2e6, 3e6, 5.244e6}

	tests := []struct {
		ref string
		fn  func(int, [3]float64) (
			c1, c2, c3 float64, c4 en.ErrNum)
	}{
		{"cgo", CgoGc2gd},
		{"go", GoGc2gd},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(0, xyz)
				test.fn(WGS84, xyz)
				test.fn(GRS80, xyz)
				test.fn(WGS72, xyz)
				test.fn(4, xyz)
			}
		})
	}
}
