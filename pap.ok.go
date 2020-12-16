package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoPap Position-angle from two p-vectors.
//
//  - - - -
//   P a p
//  - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     a      double[3]  direction of reference point
//     b      double[3]  direction of point whose PA is required
//
//  Returned (function value):
//            double     position angle of b with respect to a (radians)
//
//  Notes:
//
//  1) The result is the position angle, in radians, of direction b with
//     respect to direction a.  It is in the range -pi to +pi.  The
//     sense is such that if b is a small distance "north" of a the
//     position angle is approximately zero, and if b is a small
//     distance "east" of a the position angle is approximately +pi/2.
//
//  2) The vectors a and b need not be of unit length.
//
//  3) Zero is returned if the two directions are the same or if either
//     vector is null.
//
//  4) If vector a is at a pole, the result is ill-defined.
//
//  Called:
//     iauPn        decompose p-vector into modulus and direction
//     iauPm        modulus of p-vector
//     iauPxp       vector product of two p-vectors
//     iauPmp       p-vector minus p-vector
//     iauPdp       scalar product of two p-vectors
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPap Position-angle from two p-vectors.
func CgoPap(a, b [3]float64) (pa float64) {
	var cPa C.double
	cA := v3sGo2C(a)
	cB := v3sGo2C(b)
	cPa = C.iauPap(&cA[0], &cB[0])
	return float64(cPa)
}

//  GoPap Position-angle from two p-vectors.
func GoPap(a, b [3]float64) (pa float64) {
	var am, bm, st, ct, xa, ya, za float64
	var au, eta, xi, a2b [3]float64

	// Modulus and direction of the a vector.
	am, au = GoPn(a)

	// Modulus of the b vector.
	bm = GoPm(b)

	// Deal with the case of a null vector.
	if (am == 0.0) || (bm == 0.0) {
		st = 0.0
		ct = 1.0
	} else {

		// The "north" axis tangential from a (arbitrary length).
		xa = a[0]
		ya = a[1]
		za = a[2]
		eta[0] = -xa * za
		eta[1] = -ya * za
		eta[2] = xa*xa + ya*ya

		// The "east" axis tangential from a (same length).
		xi = GoPxp(eta, au)

		// The vector from a to b.
		a2b = GoPmp(b, a)

		// Resolve into components along the north and east axes.
		st = GoPdp(a2b, xi)
		ct = GoPdp(a2b, eta)

		// Deal with degenerate cases.
		if (st == 0.0) && (ct == 0.0) {
			ct = 1.0
		}
	}

	// Position angle.
	pa = math.Atan2(st, ct)
	return 
}
