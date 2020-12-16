package sofa

// #include "sofa.h"
import "C"

//  CgoEo06a Equation of the origins, IAU 2006 precession and IAU 2000A
//  nutation.
//
//  - - - - - -
//   E o 0 6 a
//  - - - - - -
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
//                  double    equation of the origins in radians
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
//  2) The equation of the origins is the distance between the true
//     equinox and the celestial intermediate origin and, equivalently,
//     the difference between Earth rotation angle and Greenwich
//     apparent sidereal time (ERA-GST).  It comprises the precession
//     (since J2000.0) in right ascension plus the equation of the
//     equinoxes (including the small correction terms).
//
//  Called:
//     iauPnm06a    classical NPB matrix, IAU 2006/2000A
//     iauBpn2xy    extract CIP X,Y coordinates from NPB matrix
//     iauS06       the CIO locator s, given X,Y, IAU 2006
//     iauEors      equation of the origins, given NPB matrix and s
//
//  References:
//
//     Capitaine, N. & Wallace, P.T., 2006, Astron.Astrophys. 450, 855
//
//     Wallace, P.T. & Capitaine, N., 2006, Astron.Astrophys. 459, 981
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoEo06a Equation of the origins, IAU 2006 precession and IAU 2000A
//  nutation.
func CgoEo06a(date1, date2 float64) (eo float64) {
	var cF C.double
	cF = C.iauEo06a(C.double(date1), C.double(date2))
	return float64(cF)
}

//  GoEo06a Equation of the origins, IAU 2006 precession and IAU 2000A
//  nutation.
func GoEo06a(date1, date2 float64) (eo float64) {
	var r [3][3]float64
	var x, y, s float64

	// Classical nutation x precession x bias matrix.
	r = GoPnm06a(date1, date2)

	// Extract CIP coordinates.
	x, y = GoBpn2xy(r)

	// The CIO locator, s.
	s = GoS06(date1, date2, x, y)

	// Solve for the EO.
	eo = GoEors(r, s)
	return
}
