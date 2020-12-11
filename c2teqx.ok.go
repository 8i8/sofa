package sofa

// #include "sofa.h"
import "C"

//  CgoC2teqx Assemble the celestial to terrestrial matrix from
//  equinox-based components (the celestial-to-true matrix, the
//  Greenwich Apparent Sidereal Time and the polar motion matrix).
//
//  - - - - - - -
//   C 2 t e q x
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     rbpn   double[3][3]  celestial-to-true matrix
//     gst    double        Greenwich (apparent) Sidereal Time (radians)
//     rpom   double[3][3]  polar-motion matrix
//
//  Returned:
//     rc2t   double[3][3]  celestial-to-terrestrial matrix (Note 2)
//
//  Notes:
//
//  1) This function constructs the rotation matrix that transforms
//     vectors in the celestial system into vectors in the terrestrial
//     system.  It does so starting from precomputed components, namely
//     the matrix which rotates from celestial coordinates to the
//     true equator and equinox of date, the Greenwich Apparent Sidereal
//     Time and the polar motion matrix.  One use of the present function
//     is when generating a series of celestial-to-terrestrial matrices
//     where only the Sidereal Time changes, avoiding the considerable
//     overhead of recomputing the precession-nutation more often than
//     necessary to achieve given accuracy objectives.
//
//  2) The relationship between the arguments is as follows:
//
//        [TRS] = rpom * R_3(gst) * rbpn * [CRS]
//
//              = rc2t * [CRS]
//
//     where [CRS] is a vector in the Geocentric Celestial Reference
//     System and [TRS] is a vector in the International Terrestrial
//     Reference System (see IERS Conventions 2003).
//
//  Called:
//     iauCr        copy r-matrix
//     iauRz        rotate around Z-axis
//     iauRxr       product of two r-matrices
//
//  Reference:
//
//     McCarthy, D. D., Petit, G. (eds.), IERS Conventions (2003),
//     IERS Technical Note No. 32, BKG (2004)
//
//  This revision:  2013 August 24
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoC2teqx Assemble the celestial to terrestrial matrix from
//  equinox-based components (the celestial-to-true matrix, the
//  Greenwich Apparent Sidereal Time and the polar motion matrix).
func CgoC2teqx(rbpn [3][3]float64, gst float64, rpom [3][3]float64) (
	rc2t [3][3]float64) {
	var cRc2t [3][3]C.double
	cRbpn, cRpom := v3tGo2C(rbpn), v3tGo2C(rpom)
	C.iauC2teqx(&cRbpn[0], C.double(gst), &cRpom[0], &cRc2t[0])
	return v3tC2Go(cRc2t)
}

//  GoC2teqx Assemble the celestial to terrestrial matrix from
//  equinox-based components (the celestial-to-true matrix, the
//  Greenwich Apparent Sidereal Time and the polar motion matrix).
func GoC2teqx(rbpn [3][3]float64, gst float64, rpom [3][3]float64) (
	rc2t [3][3]float64) {

	// Construct the matrix.
	rc2t = GoRxr(rpom, GoRz(gst, rbpn))
	return
}
