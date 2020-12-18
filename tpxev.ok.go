package sofa

// #include "sofa.h"
import "C"
import (
	"math"

	"github.com/8i8/sofa/en"
)

var errTpxev = en.New(0, "Tpxev", []string{
	"",
	"star too far from axis",
	"antistar on tangent plane",
	"antistar too far from axis",
})

//  CgoTpxev In the tangent plane projection, given celestial direction
//  cosines for a star and the tangent point, solve for the star's
//  rectangular coordinates in the tangent plane.
//
//  - - - - - -
//   T p x e v
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     v         double[3]  direction cosines of star (Note 4)
//     v0        double[3]  direction cosines of tangent point (Note 4)
//
//  Returned:
//     *xi,*eta  double     tangent plane coordinates of star
//
//  Returned (function value):
//               int        status: 0 = OK
//                                  1 = star too far from axis
//                                  2 = antistar on tangent plane
//                                  3 = antistar too far from axis
//
//  Notes:
//
//  1) The tangent plane projection is also called the "gnomonic
//     projection" and the "central projection".
//
//  2) The eta axis points due north in the adopted coordinate system.
//     If the direction cosines represent observed (RA,Dec), the tangent
//     plane coordinates (xi,eta) are conventionally called the
//     "standard coordinates".  If the direction cosines are with
//     respect to a right-handed triad, (xi,eta) are also right-handed.
//     The units of (xi,eta) are, effectively, radians at the tangent
//     point.
//
//  3) The method used is to extend the star vector to the tangent
//     plane and then rotate the triad so that (x,y) becomes (xi,eta).
//     Writing (a,b) for the celestial spherical coordinates of the
//     star, the sequence of rotations is (a+pi/2) around the z-axis
//     followed by (pi/2-b) around the x-axis.
//
//  4) If vector v0 is not of unit length, or if vector v is of zero
//     length, the results will be wrong.
//
//  5) If v0 points at a pole, the returned (xi,eta) will be based on
//     the arbitrary assumption that the longitude coordinate of the
//     tangent point is zero.
//
//  6) This function is a member of the following set:
//
//         spherical      vector         solve for
//
//         iauTpxes    > iauTpxev <       xi,eta
//         iauTpsts      iauTpstv          star
//         iauTpors      iauTporv         origin
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
//  CgoTpxev In the tangent plane projection, given celestial direction
//  cosines for a star and the tangent point, solve for the star's
//  rectangular coordinates in the tangent plane.
func CgoTpxev(v, v0 [3]float64) (xi, eta float64, err en.ErrNum) {
	var cXi, cEta C.double
	cV, cV0 := v3sGo2C(v), v3sGo2C(v0)
	cI := C.iauTpxev(&cV[0], &cV0[0], &cXi, &cEta)
	if n := int(cI); n != 0 {
		err = errTpxev.Set(n)
	}
	return float64(cXi), float64(cEta), err
}

//  GoTpxev In the tangent plane projection, given celestial direction
//  cosines for a star and the tangent point, solve for the star's
//  rectangular coordinates in the tangent plane.
func GoTpxev(v, v0 [3]float64) (xi, eta float64, err en.ErrNum) {
	const TINY = 1e-6
	var x, y, z, x0, y0, z0, r2, r, w, d float64

	// Star and tangent point.
	x = v[0]
	y = v[1]
	z = v[2]
	x0 = v0[0]
	y0 = v0[1]
	z0 = v0[2]

	// Deal with polar case.
	r2 = x0*x0 + y0*y0
	r = math.Sqrt(r2)
	if r == 0.0 {
		r = 1e-20
		x0 = r
	}

	// Reciprocal of star vector length to tangent plane.
	w = x*x0 + y*y0
	d = w + z*z0

	// Check for error cases.
	if d > TINY {
		err = nil
	} else if d >= 0.0 {
		err = errTpxev.Set(1)
		d = TINY
	} else if d > -TINY {
		err = errTpxev.Set(2)
		d = -TINY
	} else {
		err = errTpxev.Set(3)
	}

	// Return the tangent plane coordinates (even in dubious cases).
	d *= r
	xi = (y*x0 - x*y0) / d
	eta = (z*r2 - z0*w) / d

	// Return the status.
	return
}
