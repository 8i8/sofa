package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoGmst82 Universal Time to Greenwich mean sidereal time (IAU 1982
//  model).
//
//  - - - - - - -
//   G m s t 8 2
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  canonical model.
//
//  Given:
//     dj1,dj2    double    UT1 Julian Date (see note)
//
//  Returned (function value):
//                double    Greenwich mean sidereal time (radians)
//
//  Notes:
//
//  1) The UT1 date dj1+dj2 is a Julian Date, apportioned in any
//     convenient way between the arguments dj1 and dj2.  For example,
//     JD(UT1)=2450123.7 could be expressed in any of these ways,
//     among others:
//
//             dj1            dj2
//
//         2450123.7          0          (JD method)
//          2451545        -1421.3       (J2000 method)
//         2400000.5       50123.2       (MJD method)
//         2450123.5         0.2         (date & time method)
//
//     The JD method is the most natural and convenient to use in
//     cases where the loss of several decimal digits of resolution
//     is acceptable.  The J2000 and MJD methods are good compromises
//     between resolution and convenience.  The date & time method is
//     best matched to the algorithm used:  maximum accuracy (or, at
//     least, minimum noise) is delivered when the dj1 argument is for
//     0hrs UT1 on the day in question and the dj2 argument lies in the
//     range 0 to 1, or vice versa.
//
//  2) The algorithm is based on the IAU 1982 expression.  This is
//     always described as giving the GMST at 0 hours UT1.  In fact, it
//     gives the difference between the GMST and the UT, the steady
//     4-minutes-per-day drawing-ahead of ST with respect to UT.  When
//     whole days are ignored, the expression happens to equal the GMST
//     at 0 hours UT1 each day.
//
//  3) In this function, the entire UT1 (the sum of the two arguments
//     dj1 and dj2) is used directly as the argument for the standard
//     formula, the constant term of which is adjusted by 12 hours to
//     take account of the noon phasing of Julian Date.  The UT1 is then
//     added, but omitting whole days to conserve accuracy.
//
//  Called:
//     iauAnp       normalize angle into range 0 to 2pi
//
//  References:
//
//     Transactions of the International Astronomical Union,
//     XVIII B, 67 (1983).
//
//     Aoki et al., Astron.Astrophys., 105, 359-361 (1982).
//
//  This revision:  2020 January 12
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoGmst82 Universal Time to Greenwich mean sidereal time (IAU 1982
//  model).
func CgoGmst82(dj1, dj2 float64) (gmst float64) {
	var cF C.double
	cF = C.iauGmst82(C.double(dj1), C.double(dj2))
	return float64(cF)
}

//  GoGmst82 Universal Time to Greenwich mean sidereal time (IAU 1982
//  model).
func GoGmst82(dj1, dj2 float64) (gmst float64) {

	// Coefficients of IAU 1982 GMST-UT1 model
	var A = 24110.54841 - DAYSEC/2.0
	var B = 8640184.812866
	var C = 0.093104
	var D = -6.2e-6

	// The first constant, A, has to be adjusted by 12 hours because
	// the UT1 is supplied as a Julian date, which begins at noon.

	var d1, d2, t, f float64

	// Julian centuries since fundamental epoch.
	if dj1 < dj2 {
		d1 = dj1
		d2 = dj2
	} else {
		d1 = dj2
		d2 = dj1
	}
	t = (d1 + (d2 - DJ00)) / DJC

	// Fractional part of JD(UT1), in seconds.
	f = DAYSEC * (math.Mod(d1, 1.0) + math.Mod(d2, 1.0))

	// GMST at this UT1.
	gmst = GoAnp(DS2R * ((A + (B+(C+D*t)*t)*t) + f))
	return
}
