package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t P d p
//  - - - - - - - -
//
//  Test Pdp function.
//
//  Called:  Pdp, vvd
//
//  This revision:  2013 August 7
//
func TestPdp(t *testing.T) {
	const fname = "Pdp"
	var a, b [3]float64
	var adb float64

	a[0] = 2.0
	a[1] = 2.0
	a[2] = 3.0

	b[0] = 1.0
	b[1] = 3.0
	b[2] = 4.0

	adb = Pdp(a, b)

	vvd(t, adb, 20, 1e-12, fname, "")
}
