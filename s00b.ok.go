package sofa

// #include "sofa.h"
import "C"

//  CgoS00b The CIO locator s, positioning the Celestial Intermediate
//  Origin on the equator of the Celestial Intermediate Pole, using the
//  IAU 2000B precession-nutation model.
//
//  - - - - -
//   S 0 0 b
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     date1,date2  double    TT as a 2-part Julian Date (Note 1)
//
//  Returned (function value):
//                  double    the CIO locator s in radians (Note 2)
//
//  Notes:
//
//  1) The TT date date1+date2 is a Julian Date, apportioned in any
//     convenient way between the two arguments.  For example,
//     JD(TT)=2450123.7 could be expressed in any of these ways,
//     among others:
//
//            date1          date2
//
//         2450123.7           0.0       (JD method)
//         2451545.0       -1421.3       (J2000 method)
//         2400000.5       50123.2       (MJD method)
//         2450123.5           0.2       (date & time method)
//
//     The JD method is the most natural and convenient to use in
//     cases where the loss of several decimal digits of resolution
//     is acceptable.  The J2000 method is best matched to the way
//     the argument is handled internally and will deliver the
//     optimum resolution.  The MJD method and the date & time methods
//     are both good compromises between resolution and convenience.
//
//  2) The CIO locator s is the difference between the right ascensions
//     of the same point in two systems.  The two systems are the GCRS
//     and the CIP,CIO, and the point is the ascending node of the
//     CIP equator.  The CIO locator s remains a small fraction of
//     1 arcsecond throughout 1900-2100.
//
//  3) The series used to compute s is in fact for s+XY/2, where X and Y
//     are the x and y components of the CIP unit vector;  this series
//     is more compact than a direct series for s would be.  The present
//     function uses the IAU 2000B truncated nutation model when
//     predicting the CIP position.  The function iauS00a uses instead
//     the full IAU 2000A model, but with no significant increase in
//     accuracy and at some cost in speed.
//
//  Called:
//     iauPnm00b    classical NPB matrix, IAU 2000B
//     iauBnp2xy    extract CIP X,Y from the BPN matrix
//     iauS00       the CIO locator s, given X,Y, IAU 2000A
//
//  References:
//
//     Capitaine, N., Chapront, J., Lambert, S. and Wallace, P.,
//     "Expressions for the Celestial Intermediate Pole and Celestial
//     Ephemeris Origin consistent with the IAU 2000A precession-
//     nutation model", Astron.Astrophys. 400, 1145-1154 (2003)
//
//     n.b. The celestial ephemeris origin (CEO) was renamed "celestial
//          intermediate origin" (CIO) by IAU 2006 Resolution 2.
//
//     McCarthy, D. D., Petit, G. (eds.), IERS Conventions (2003),
//     IERS Technical Note No. 32, BKG (2004)
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoS00b The CIO locator s, positioning the Celestial Intermediate
//  Origin on the equator of the Celestial Intermediate Pole, using the
//  IAU 2000B precession-nutation model.
func CgoS00b(date1, date2 float64) (s float64) {
	var cS C.double
	cS = C.iauS00b(C.double(date1), C.double(date2))
	return float64(cS)
}

//  GoS00b The CIO locator s, positioning the Celestial Intermediate
//  Origin on the equator of the Celestial Intermediate Pole, using the
//  IAU 2000B precession-nutation model.
func GoS00b(date1, date2 float64) (s float64) {
	var rbpn [3][3]float64
	var x, y float64

	// Bias-precession-nutation-matrix, IAU 2000B.
	rbpn = GoPnm00b(date1, date2)

	// Extract the CIP coordinates.
	x, y = GoBpn2xy(rbpn)

	// Compute the CIO locator s, given the CIP coordinates.
	s = GoS00(date1, date2, x, y)
	return
}
