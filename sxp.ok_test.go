package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t S x p
//  - - - - - - - -
//
//  Test Sxp function.
//
//  Called:  Sxp, vvd
//
//  This revision:  2013 August 7
//
func TestSxp(t *testing.T) {
	const fname = "Sxp"
	var s float64
	var p, sp [3]float64

	s = 2.0
	p[0] = 0.3
	p[1] = 1.2
	p[2] = -2.5

	sp = Sxp(s, p)

	vvd(t, sp[0], 0.6, 0.0, fname, "1")
	vvd(t, sp[1], 2.4, 0.0, fname, "2")
	vvd(t, sp[2], -5.0, 0.0, fname, "3")
}
