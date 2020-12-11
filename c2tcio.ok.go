package sofa

// #include "sofa.h"
import "C"

//  CgoC2tcio Assemble the celestial to terrestrial matrix from
//  CIO-based components (the celestial-to-intermediate matrix, the
//  Earth Rotation Angle and the polar motion matrix).
//
//  - - - - - - -
//   C 2 t c i o
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     rc2i     double[3][3]    celestial-to-intermediate matrix
//     era      double          Earth rotation angle (radians)
//     rpom     double[3][3]    polar-motion matrix
//
//  Returned:
//     rc2t     double[3][3]    celestial-to-terrestrial matrix
//
//  Notes:
//
//  1) This function constructs the rotation matrix that transforms
//     vectors in the celestial system into vectors in the terrestrial
//     system.  It does so starting from precomputed components, namely
//     the matrix which rotates from celestial coordinates to the
//     intermediate frame, the Earth rotation angle and the polar motion
//     matrix.  One use of the present function is when generating a
//     series of celestial-to-terrestrial matrices where only the Earth
//     Rotation Angle changes, avoiding the considerable overhead of
//     recomputing the precession-nutation more often than necessary to
//     achieve given accuracy objectives.
//
//  2) The relationship between the arguments is as follows:
//
//        [TRS] = RPOM * R_3(ERA) * rc2i * [CRS]
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
//     McCarthy, D. D., Petit, G. (eds.), 2004, IERS Conventions (2003),
//     IERS Technical Note No. 32, BKG
//
//  This revision:  2013 August 24
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoC2tcio Assemble the celestial to terrestrial matrix from
//  CIO-based components (the celestial-to-intermediate matrix, the
//  Earth Rotation Angle and the polar motion matrix).
func CgoC2tcio(rc2i [3][3]float64, era float64, rpom [3][3]float64) (
	rc2t [3][3]float64) {
	var cRc2t [3][3]C.double
	cRc2i, cRpom := v3tGo2C(rc2i), v3tGo2C(rpom)
	C.iauC2tcio(&cRc2i[0], C.double(era), &cRpom[0], &cRc2t[0])
	return v3tC2Go(cRc2t)
}

//  GoC2tcio Assemble the celestial to terrestrial matrix from
//  CIO-based components (the celestial-to-intermediate matrix, the
//  Earth Rotation Angle and the polar motion matrix).
func GoC2tcio(rc2i [3][3]float64, era float64, rpom [3][3]float64) (
	rc2t [3][3]float64) {

	// Construct the matrix.
	rc2t = GoRxr(rpom, GoRz(era, rc2i))
	return
}
