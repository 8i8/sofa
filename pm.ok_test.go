package sofa

import "testing"

//
//  - - - - - - -
//   T e s t P m
//  - - - - - - -
//
//  Test Pm function.
//
//  Called:  iauPm, vvd
//
//  This revision:  2013 August 7
//
func TestPm(t *testing.T) {
	const fname = "Pm"
	var p [3]float64
	var r float64
	p[0] = 0.3
	p[1] = 1.2
	p[2] = -2.5

	r = Pm(p)
	vvd(t, r, 2.789265136196270604, 1e-12, fname, "")
}
