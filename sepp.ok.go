package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoSepp Angular separation between two p-vectors.
//
//  - - - - -
//   S e p p
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     a      double[3]    first p-vector (not necessarily unit length)
//     b      double[3]    second p-vector (not necessarily unit length)
//
//  Returned (function value):
//            double       angular separation (radians, always positive)
//
//  Notes:
//
//  1) If either vector is null, a zero result is returned.
//
//  2) The angular separation is most simply formulated in terms of
//     scalar product.  However, this gives poor accuracy for angles
//     near zero and pi.  The present algorithm uses both cross product
//     and dot product, to deliver full accuracy whatever the size of
//     the angle.
//
//  Called:
//     iauPxp       vector product of two p-vectors
//     iauPm        modulus of p-vector
//     iauPdp       scalar product of two p-vectors
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoSepp Angular separation between two p-vectors.
func CgoSepp(a, b [3]float64) (r float64) {
	var cR C.double
	cA, cB := v3sGo2C(a), v3sGo2C(b)
	cR = C.iauSepp(&cA[0], &cB[0])
	return float64(cR)
}

//  GoSepp Angular separation between two p-vectors.
func GoSepp(a, b [3]float64) (s float64) {

	var axb [3]float64
	var ss, cs float64

	// Sine of angle between the vectors, multiplied by the two
	// moduli.
	axb = GoPxp(a, b)
	ss = GoPm(axb)

	// Cosine of the angle, multiplied by the two moduli.
	cs = GoPdp(a, b)

	// The angle.
	if (ss != 0.0) || (cs != 0.0) {
		s = math.Atan2(ss, cs)
	}
	return
}
