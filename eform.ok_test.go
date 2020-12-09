package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t E f o r m
//  - - - - - - - - - -
//
//  Test Eform function.
//
//  Called:  Eform, viv, vvd
//
//  This revision:  2016 March 12
//
func TestEform(t *testing.T) {
	const fname = "Eform"
	tests := []struct {
		ref string
		fn  func(int) (float64, float64, error)
	}{
		{"cgo", CgoEform},
		{"go", GoEform},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref

		a, f, err := test.fn(0) // error raised
		errT(t, errEformE1, err, tname+" 0")

		a, f, err = test.fn(WGS84)
		errT(t, nil, err, tname+" 1")
		vvd(t, a, 6378137.0, 1e-10, tname, " 1")
		vvd(t, f, 0.3352810664747480720e-2, 1e-18, tname, "f1")

		a, f, err = test.fn(GRS80)
		errT(t, nil, err, tname+" 2")
		vvd(t, a, 6378137.0, 1e-10, tname, "a2")
		vvd(t, f, 0.3352810681182318935e-2, 1e-18, tname, "f2")

		a, f, err = test.fn(WGS72)
		errT(t, nil, err, tname+" 2")
		vvd(t, a, 6378135.0, 1e-10, tname, "a3")
		vvd(t, f, 0.3352779454167504862e-2, 1e-18, tname, "f3")

		a, f, err = test.fn(4) // error raised
		errT(t, errEformE1, err, tname+" 3")
	}
}

func BenchmarkEform(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(int) (float64, float64, error)
	}{
		{"cgo", CgoEform},
		{"go", GoEform},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = test.fn(0)
				_, _, _ = test.fn(WGS84)
				_, _, _ = test.fn(GRS80)
				_, _, _ = test.fn(WGS72)
				_, _, _ = test.fn(4)
			}
		})
	}
}
