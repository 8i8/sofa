package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t A e 2 h d
//  - - - - - - - - - -
//
//  Test Ae2hd function.
//
//  Called:  iauAe2hd and vvd
//
//  This revision:  2017 October 21
//
func TestAe2hd(t *testing.T) {
	const fname = "Ae2hd"
	var a, e, p, h, d float64

	a = 5.5
	e = 1.1
	p = 0.7

	h, d = Ae2hd(a, e, p)

	vvd(t, h, 0.5933291115507309663, 1e-14, fname, "h")
	vvd(t, d, 0.9613934761647817620, 1e-14, fname, "d")
}
