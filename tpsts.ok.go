package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoTpsts In the tangent plane projection, given the star's
//  rectangular coordinates and the spherical coordinates of the tangent
//  point, solve for the spherical coordinates of the star.
//
//  - - - - - -
//   T p s t s
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     xi,eta    double  rectangular coordinates of star image (Note 2)
//     a0,b0     double  tangent point's spherical coordinates
//
//  Returned:
//     *a,*b     double  star's spherical coordinates
//
//  1) The tangent plane projection is also called the "gnomonic
//     projection" and the "central projection".
//
//  2) The eta axis points due north in the adopted coordinate system.
//     If the spherical coordinates are observed (RA,Dec), the tangent
//     plane coordinates (xi,eta) are conventionally called the
//     "standard coordinates".  If the spherical coordinates are with
//     respect to a right-handed triad, (xi,eta) are also right-handed.
//     The units of (xi,eta) are, effectively, radians at the tangent
//     point.
//
//  3) All angular arguments are in radians.
//
//  4) This function is a member of the following set:
//
//         spherical      vector         solve for
//
//         iauTpxes      iauTpxev         xi,eta
//       > iauTpsts <    iauTpstv          star
//         iauTpors      iauTporv         origin
//
//  Called:
//     iauAnp       normalize angle into range 0 to 2pi
//
//  References:
//
//     Calabretta M.R. & Greisen, E.W., 2002, "Representations of
//     celestial coordinates in FITS", Astron.Astrophys. 395, 1077
//
//     Green, R.M., "Spherical Astronomy", Cambridge University Press,
//     1987, Chapter 13.
//
//  This revision:   2018 January 2
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoTpsts In the tangent plane projection, given the star's
//  rectangular coordinates and the spherical coordinates of the tangent
//  point, solve for the spherical coordinates of the star.
func CgoTpsts(xi, eta, a0, b0 float64) (a, b float64) {
	var cA, cB C.double
	C.iauTpsts(C.double(xi), C.double(eta), C.double(a0), C.double(b0),
		&cA, &cB)
	return float64(cA), float64(cB)
}

//  GoTpsts In the tangent plane projection, given the star's
//  rectangular coordinates and the spherical coordinates of the tangent
//  point, solve for the spherical coordinates of the star.
func GoTpsts(xi, eta, a0, b0 float64) (a, b float64) {
	var sb0, cb0, d float64

	sb0 = math.Sin(b0)
	cb0 = math.Cos(b0)
	d = cb0 - eta*sb0
	a = GoAnp(math.Atan2(xi, d) + a0)
	b = math.Atan2(sb0+eta*cb0, math.Sqrt(xi*xi+d*d))
	return
}
