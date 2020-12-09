package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t A p e r 1 3
//  - - - - - - - - - - -
//
//  Test Aper13 function.
//
//  Called:  Aper13, vvd
//
//  This revision:  2013 October 3
//
func TestAper13(t *testing.T) {
	const fname = "Aper13"
	var ut11, ut12 float64
	var astrom ASTROM

	astrom.along = 1.234
	ut11 = 2456165.5
	ut12 = 0.401182685

	tests := []struct {
		ref string
		fn  func(a1, a2 float64, a3 ASTROM) ASTROM
	}{
		{"cgo", CgoAper13},
		{"go", GoAper13},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		astr := test.fn(ut11, ut12, astrom)

		vvd(t, astr.eral, 3.316236661789694933, 1e-12,
			tname, "pmt")
	}
}

func BenchmarkAper13(b *testing.B) {
	var ut11, ut12 float64
	var astrom ASTROM

	astrom.along = 1.234
	ut11 = 2456165.5
	ut12 = 0.401182685

	tests := []struct {
		ref string
		fn  func(a1, a2 float64, a3 ASTROM) ASTROM
	}{
		{"cgo", CgoAper13},
		{"go", GoAper13},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(ut11, ut12, astrom)
			}
		})
	}
}
