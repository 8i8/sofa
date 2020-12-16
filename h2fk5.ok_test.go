package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t H 2 f k 5
//  - - - - - - - - - -
//
//  Test H2fk5 function.
//
//  Called:  H2fk5, vvd
//
//  This revision:  2017 January 3
//
func TestH2fk5(t *testing.T) {
	const fname = "H2fk5"
	var rh, dh, drh, ddh, pxh, rvh,
		r5, d5, dr5, dd5, px5, rv5 float64

	rh = 1.767794352
	dh = -0.2917512594
	drh = -2.76413026e-6
	ddh = -5.92994449e-6
	pxh = 0.379210
	rvh = -7.6

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6 float64) (
			c1, c2, c3, c4, c5, c6 float64)
	}{
		{"cgo", CgoH2fk5},
		{"go", GoH2fk5},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		r5, d5, dr5, dd5, px5, rv5 = test.fn(
			rh, dh, drh, ddh, pxh, rvh)

		vvd(t, r5, 1.767794455700065506, 1e-13,
			tname, "ra")
		vvd(t, d5, -0.2917513626469638890, 1e-13,
			tname, "dec")
		vvd(t, dr5, -0.27597945024511204e-5, 1e-18,
			tname, "dr5")
		vvd(t, dd5, -0.59308014093262838e-5, 1e-18,
			tname, "dd5")
		vvd(t, px5, 0.37921, 1e-13,
			tname, "px")
		vvd(t, rv5, -7.6000001309071126, 1e-11,
			tname, "rv")
	}
}

func BenchmarkH2fk5(b *testing.B) {
	var rh, dh, drh, ddh, pxh, rvh float64

	rh = 1.767794352
	dh = -0.2917512594
	drh = -2.76413026e-6
	ddh = -5.92994449e-6
	pxh = 0.379210
	rvh = -7.6

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6 float64) (
			c1, c2, c3, c4, c5, c6 float64)
	}{
		{"cgo", CgoH2fk5},
		{"go", GoH2fk5},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(rh, dh, drh, ddh, pxh, rvh)
			}
		})
	}
}
