package sofa

import "testing"

//
//  - - - - - - -
//   T e s t P n
//  - - - - - - -
//
//  Test Pn function.
//
//  Called:  Pn, vvd
//
//  This revision:  2013 August 7
//
func TestPn(t *testing.T) {
	const fname = "Pn"
	var p, u [3]float64
	var r float64

	p[0] = 0.3
	p[1] = 1.2
	p[2] = -2.5

	u, r = Pn(p)

	vvd(t, r, 2.789265136196270604, 1e-12, fname, "r")

	vvd(t, u[0], 0.1075552109073112058, 1e-12, fname, "u1")
	vvd(t, u[1], 0.4302208436292448232, 1e-12, fname, "u2")
	vvd(t, u[2], -0.8962934242275933816, 1e-12, fname, "u3")
}
