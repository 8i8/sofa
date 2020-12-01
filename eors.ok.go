package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoEors Equation of the origins, given the classical NPB matrix and
//  the quantity s.
//
//  - - - - -
//   E o r s
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     rnpb  double[3][3]  classical nutation x precession x bias matrix
//     s     double        the quantity s (the CIO locator)
//
//  Returned (function value):
//           double        the equation of the origins in radians.
//
//  Notes:
//
//  1)  The equation of the origins is the distance between the true
//      equinox and the celestial intermediate origin and, equivalently,
//      the difference between Earth rotation angle and Greenwich
//      apparent sidereal time (ERA-GST).  It comprises the precession
//      (since J2000.0) in right ascension plus the equation of the
//      equinoxes (including the small correction terms).
//
//  2)  The algorithm is from Wallace & Capitaine (2006).
//
// References:
//
//     Capitaine, N. & Wallace, P.T., 2006, Astron.Astrophys. 450, 855
//
//     Wallace, P. & Capitaine, N., 2006, Astron.Astrophys. 459, 981
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoEors Equation of the origins, given the classical NPB matrix and
//  the quantity s.
func CgoEors(rnpb [3][3]float64, s float64) (eo float64) {
	cRnpb := v3tGo2C(rnpb)
	cF := C.iauEors(&cRnpb[0], C.double(s))
	return float64(cF)
}

// GoEors Equation of the origins, given the classical NPB matrix and the
// quantity s.
func GoEors(rnpb [3][3]float64, s float64) (eo float64) {
	var x, ax, xs, ys, zs, p, q float64

	// Evaluate Wallace & Capitaine (2006) expression (16).
	x = rnpb[2][0]
	ax = x / (1.0 + rnpb[2][2])
	xs = 1.0 - ax*x
	ys = -ax * rnpb[2][1]
	zs = -x
	p = rnpb[0][0]*xs + rnpb[0][1]*ys + rnpb[0][2]*zs
	q = rnpb[1][0]*xs + rnpb[1][1]*ys + rnpb[1][2]*zs
	eo = s
	if p != 0 || q != 0 {
		eo = s - math.Atan2(q, p)
	}
	return
}
