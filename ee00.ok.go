package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoEe00 The equation of the equinoxes, compatible with IAU 2000
//  resolutions, given the nutation in longitude and the mean obliquity.
//
//  - - - - -
//   E e 0 0
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  canonical model.
//
//  Given:
//     date1,date2  double    TT as a 2-part Julian Date (Note 1)
//     epsa         double    mean obliquity (Note 2)
//     dpsi         double    nutation in longitude (Note 3)
//
//  Returned (function value):
//                  double    equation of the equinoxes (Note 4)
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
//  2) The obliquity, in radians, is mean of date.
//
//  3) The result, which is in radians, operates in the following sense:
//
//        Greenwich apparent ST = GMST + equation of the equinoxes
//
//  4) The result is compatible with the IAU 2000 resolutions.  For
//     further details, see IERS Conventions 2003 and Capitaine et al.
//     (2002).
//
//  Called:
//     iauEect00    equation of the equinoxes complementary terms
//
//  References:
//
//     Capitaine, N., Wallace, P.T. and McCarthy, D.D., "Expressions to
//     implement the IAU 2000 definition of UT1", Astronomy &
//     Astrophysics, 406, 1135-1149 (2003)
//
//     McCarthy, D. D., Petit, G. (eds.), IERS Conventions (2003),
//     IERS Technical Note No. 32, BKG (2004)
//
//  This revision:  2008 May 16
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoEe00 The equation of the equinoxes, compatible with IAU 2000
//  resolutions, given the nutation in longitude and the mean obliquity.
func CgoEe00(date1, date2, epsa, dpsi float64) float64 {
	var cF C.double
	cF = C.iauEe00(C.double(date1), C.double(date2), C.double(epsa),
		C.double(dpsi))
	return float64(cF)
}

//  GoEe00 The equation of the equinoxes, compatible with IAU 2000
//  resolutions, given the nutation in longitude and the mean obliquity.
func GoEe00(date1, date2, epsa, dpsi float64) (ee float64) {

	// Equation of the equinoxes.
	ee = dpsi*math.Cos(epsa) + GoEect00(date1, date2)
	return
}
