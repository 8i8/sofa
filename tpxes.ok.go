package sofa

// #include "sofa.h"
import "C"
import (
	"math"

	"github.com/8i8/sofa/en"
)

var errTpxex = en.New(0, "Tpxex", []string{
	"",
	"star too far from axis",
	"antistar on tangent plane",
	"antistar too far from axis",
})

//  CgoTpxes In the tangent plane projection, given celestial spherical
//  coordinates for a star and the tangent point, solve for the star's
//  rectangular coordinates in the tangent plane.
//
//  - - - - - -
//   T p x e s
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     a,b       double  star's spherical coordinates
//     a0,b0     double  tangent point's spherical coordinates
//
//  Returned:
//     *xi,*eta  double  rectangular coordinates of star image (Note 2)
//
//  Returned (function value):
//               int     status:  0 = OK
//                                1 = star too far from axis
//                                2 = antistar on tangent plane
//                                3 = antistar too far from axis
//
//  Notes:
//
//  1) The tangent plane projection is also called the "gnomonic
//     projection" and the "central projection".
//
//  2) The eta axis points due north in the adopted coordinate system.
//     If the spherical coordinates are observed (RA,Dec), the tangent
//     plane coordinates (xi,eta) are conventionally called the
//     "standard coordinates".  For right-handed spherical coordinates,
//     (xi,eta) are also right-handed.  The units of (xi,eta) are,
//     effectively, radians at the tangent point.
//
//  3) All angular arguments are in radians.
//
//  4) This function is a member of the following set:
//
//         spherical      vector         solve for
//
//       > iauTpxes <    iauTpxev         xi,eta
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
//  CgoTpxes In the tangent plane projection, given celestial spherical
//  coordinates for a star and the tangent point, solve for the star's
//  rectangular coordinates in the tangent plane.
func CgoTpxes(a, b, a0, b0 float64) (xi, eta float64, err en.ErrNum) {
	var cXi, cEta C.double
	C.iauTpxes(C.double(a), C.double(b), C.double(a0), C.double(b0),
		&cXi, &cEta)
	return float64(cXi), float64(cEta), err
}

//  GoTpxes In the tangent plane projection, given celestial spherical
//  coordinates for a star and the tangent point, solve for the star's
//  rectangular coordinates in the tangent plane.
func GoTpxes(a, b, a0, b0 float64) (xi, eta float64, err en.ErrNum) {
	const TINY = 1e-6
	var sb0, sb, cb0, cb, da, sda, cda, d float64

	// Functions of the spherical coordinates.
	sb0 = math.Sin(b0)
	sb = math.Sin(b)
	cb0 = math.Cos(b0)
	cb = math.Cos(b)
	da = a - a0
	sda = math.Sin(da)
	cda = math.Cos(da)

	// Reciprocal of star vector length to tangent plane.
	d = sb*sb0 + cb*cb0*cda

	// Check for error cases.
	if d > TINY {
		err = nil
	} else if d >= 0.0 {
		err = errTpxex.Set(1)
		d = TINY
	} else if d > -TINY {
		err = errTpxex.Set(2)
		d = -TINY
	} else {
		err = errTpxex.Set(3)
	}

	// Return the tangent plane coordinates (even in dubious cases).
	xi = cb * sda / d
	eta = (sb*cb0 - cb*sb0*cda) / d

	// Return the status.
	return
}
