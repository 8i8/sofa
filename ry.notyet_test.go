package sofa

import "testing"

//
//  - - - - - - -
//   T e s t R y
//  - - - - - - -
//
//  Test iauRy function.
//
//  Called:  Ry, vvd
//
//  This revision:  2013 August 7
//
func TestRy(t *testing.T) {
	const fname = "Ry"
	var theta float64
	var r [3][3]float64

	theta = 0.3456789

	r[0][0] = 2.0
	r[0][1] = 3.0
	r[0][2] = 2.0

	r[1][0] = 3.0
	r[1][1] = 2.0
	r[1][2] = 3.0

	r[2][0] = 3.0
	r[2][1] = 4.0
	r[2][2] = 5.0

	r = Ry(theta, r)

	vvd(t, r[0][0], 0.8651847818978159930, 1e-12, fname, "11")
	vvd(t, r[0][1], 1.467194920539316554, 1e-12, fname, "12")
	vvd(t, r[0][2], 0.1875137911274457342, 1e-12, fname, "13")

	vvd(t, r[1][0], 3, 1e-12, fname, "21")
	vvd(t, r[1][1], 2, 1e-12, fname, "22")
	vvd(t, r[1][2], 3, 1e-12, fname, "23")

	vvd(t, r[2][0], 3.500207892850427330, 1e-12, fname, "31")
	vvd(t, r[2][1], 4.779889022262298150, 1e-12, fname, "32")
	vvd(t, r[2][2], 5.381899160903798712, 1e-12, fname, "33")
}
