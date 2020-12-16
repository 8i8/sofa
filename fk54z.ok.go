package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoFk54z Convert a J2000.0 FK5 star position to B1950.0 FK4, assuming zero
//  proper motion in FK5 and parallax.
//
//  - - - - - -
//   F k 5 4 z
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     r2000,d2000    double   J2000.0 FK5 RA,Dec (rad)
//     bepoch         double   Besselian epoch (e.g. 1950.0)
//
//  Returned:
//     r1950,d1950    double   B1950.0 FK4 RA,Dec (rad) at epoch BEPOCH
//     dr1950,dd1950  double   B1950.0 FK4 proper motions (rad/trop.yr)
//
//  Notes:
//
//  1) In contrast to the iauFk524  routine, here the FK5 proper
//     motions, the parallax and the radial velocity are presumed zero.
//
//  2) This function converts a star position from the IAU 1976 FK5
//    (Fricke) system to the former FK4 (Bessel-Newcomb) system, for
//     cases such as distant radio sources where it is presumed there is
//     zero parallax and no proper motion.  Because of the E-terms of
//     aberration, such objects have (in general) non-zero proper motion
//     in FK4, and the present routine returns those fictitious proper
//     motions.
//
//  3) Conversion from B1950.0 FK4 to J2000.0 FK5 only is provided for.
//     Conversions involving other equinoxes would require additional
//     treatment for precession.
//
//  4) The position returned by this routine is in the B1950.0 FK4
//     reference system but at Besselian epoch BEPOCH.  For comparison
//     with catalogs the BEPOCH argument will frequently be 1950.0. (In
//     this context the distinction between Besselian and Julian epoch
//     is insignificant.)
//
//  5) The RA component of the returned (fictitious) proper motion is
//     dRA/dt rather than cos(Dec)*dRA/dt.
//
//  Called:
//     iauAnp       normalize angle into range 0 to 2pi
//     iauC2s       p-vector to spherical
//     iauFk524     FK4 to FK5
//     iauS2c       spherical to p-vector
//
//  This revision:   2018 December 5
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoFk54z Convert a J2000.0 FK5 star position to B1950.0 FK4, assuming zero
//  proper motion in FK5 and parallax.
func CgoFk54z(r2000, d2000, bepoch float64) (
	r1950, d1950, dr1950, dd1950 float64) {
	var cR1950, cD1950, cDr1950, cDd1950 C.double
	C.iauFk54z(C.double(r2000), C.double(d2000), C.double(bepoch),
		&cR1950, &cD1950, &cDr1950, &cDd1950)
	return float64(cR1950), float64(cD1950), float64(cDr1950),
		float64(cDd1950)
}

//  GoFk54z Convert a J2000.0 FK5 star position to B1950.0 FK4, assuming zero
//  proper motion in FK5 and parallax.
func GoFk54z(r2000, d2000, bepoch float64) (
	r1950, d1950, dr1950, dd1950 float64) {
	var r, d, pr, pd, w float64
	var p, v [3]float64
	var i int

	// FK5 equinox J2000.0 to FK4 equinox B1950.0.
	r, d, pr, pd, _, _ = GoFk524(r2000, d2000, 0.0, 0.0, 0.0, 0.0)

	// Spherical to Cartesian.
	p = GoS2c(r, d)

	// Fictitious proper motion (radians per year).
	v[0] = -pr*p[1] - pd*math.Cos(r)*math.Sin(d)
	v[1] = pr*p[0] - pd*math.Sin(r)*math.Sin(d)
	v[2] = pd * math.Cos(d)

	// Apply the motion.
	w = bepoch - 1950.0
	for i = 0; i < 3; i++ {
		p[i] += w * v[i]
	}

	// Cartesian to spherical.
	w, d1950 = GoC2s(p)
	r1950 = GoAnp(w)

	// Fictitious proper motion.
	dr1950 = pr
	dd1950 = pd

	// Finished.
	return
}
