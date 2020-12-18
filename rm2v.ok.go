package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoRm2v Express an r-matrix as an r-vector.
//
//  - - - - -
//   R m 2 v
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     r        double[3][3]    rotation matrix
//
//  Returned:
//     w        double[3]       rotation vector (Note 1)
//
//  Notes:
//
//  1) A rotation matrix describes a rotation through some angle about
//     some arbitrary axis called the Euler axis.  The "rotation vector"
//     returned by this function has the same direction as the Euler axis,
//     and its magnitude is the angle in radians.  (The magnitude and
//     direction can be separated by means of the function iauPn.)
//
//  2) If r is null, so is the result.  If r is not a rotation matrix
//     the result is undefined;  r must be proper (i.e. have a positive
//     determinant) and real orthogonal (inverse = transpose).
//
//  3) The reference frame rotates clockwise as seen looking along
//     the rotation vector from the origin.
//
//  This revision:  2015 January 30
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoRm2v Express an r-matrix as an r-vector.
func CgoRm2v(r [3][3]float64) (w [3]float64) {
	var cW [3]C.double
	cR := v3tGo2C(r)
	C.iauRm2v(&cR[0], &cW[0])
	return v3sC2Go(cW)
}

//  GoRm2v Express an r-matrix as an r-vector.
func GoRm2v(r [3][3]float64) (w [3]float64) {
	var x, y, z, s2, c2, phi, f float64

	x = r[1][2] - r[2][1]
	y = r[2][0] - r[0][2]
	z = r[0][1] - r[1][0]
	s2 = math.Sqrt(x*x + y*y + z*z)
	if s2 > 0 {
		c2 = r[0][0] + r[1][1] + r[2][2] - 1.0
		phi = math.Atan2(s2, c2)
		f = phi / s2
		w[0] = x * f
		w[1] = y * f
		w[2] = z * f
	} else {
		w[0] = 0.0
		w[1] = 0.0
		w[2] = 0.0
	}
	return
}
