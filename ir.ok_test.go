package sofa

import "testing"

//
//  - - - - - - -
//   T e s t I r
//  - - - - - - -
//
//  Test Ir function.
//
//  Called:  Ir, vvd
//
//  This revision:  2013 August 7
//
func TestIr(t *testing.T) {
	const fname = "Ir"
	var r [3][3]float64

	r = Ir()

	vvd(t, r[0][0], 1.0, 0.0, fname, "11")
	vvd(t, r[0][1], 0.0, 0.0, fname, "12")
	vvd(t, r[0][2], 0.0, 0.0, fname, "13")

	vvd(t, r[1][0], 0.0, 0.0, fname, "21")
	vvd(t, r[1][1], 1.0, 0.0, fname, "22")
	vvd(t, r[1][2], 0.0, 0.0, fname, "23")

	vvd(t, r[2][0], 0.0, 0.0, fname, "31")
	vvd(t, r[2][1], 0.0, 0.0, fname, "32")
	vvd(t, r[2][2], 1.0, 0.0, fname, "33")
}
