package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoRv2m Form the r-matrix corresponding to a given r-vector.
//
//  - - - - -
//   R v 2 m
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     w        double[3]      rotation vector (Note 1)
//
//  Returned:
//     r        double[3][3]    rotation matrix
//
//  Notes:
//
//  1) A rotation matrix describes a rotation through some angle about
//     some arbitrary axis called the Euler axis.  The "rotation vector"
//     supplied to This function has the same direction as the Euler
//     axis, and its magnitude is the angle in radians.
//
//  2) If w is null, the unit matrix is returned.
//
//  3) The reference frame rotates clockwise as seen looking along the
//     rotation vector from the origin.
//
//  This revision:  2015 January 30
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoRv2m Form the r-matrix corresponding to a given r-vector.
func CgoRv2m(w [3]float64) (r [3][3]float64) {
	var cR [3][3]C.double
	cW := v3sGo2C(w)
	C.iauRv2m(&cW[0], &cR[0])
	return v3tC2Go(cR)
}

//  GoRv2m Form the r-matrix corresponding to a given r-vector.
func GoRv2m(w [3]float64) (r [3][3]float64) {
	var x, y, z, phi, s, c, f float64

	// Euler angle (magnitude of rotation vector) and functions.
	x = w[0]
	y = w[1]
	z = w[2]
	phi = math.Sqrt(x*x + y*y + z*z)
	s = math.Sin(phi)
	c = math.Cos(phi)
	f = 1.0 - c

	// Euler axis (direction of rotation vector), perhaps null.
	if phi > 0.0 {
		x /= phi
		y /= phi
		z /= phi
	}

	// Form the rotation matrix.
	r[0][0] = x*x*f + c
	r[0][1] = x*y*f + z*s
	r[0][2] = x*z*f - y*s
	r[1][0] = y*x*f - z*s
	r[1][1] = y*y*f + c
	r[1][2] = y*z*f + x*s
	r[2][0] = z*x*f + y*s
	r[2][1] = z*y*f - x*s
	r[2][2] = z*z*f + c
	return
}
