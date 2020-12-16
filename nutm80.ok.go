package sofa

// #include "sofa.h"
import "C"

//  CgoNutm80 Form the matrix of nutation for a given date, IAU 1980
//  model.
//
//  - - - - - - -
//   N u t m 8 0
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     date1,date2    double          TDB date (Note 1)
//
//  Returned:
//     rmatn          double[3][3]    nutation matrix
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
//  2) The matrix operates in the sense V(true) = rmatn * V(mean),
//     where the p-vector V(true) is with respect to the true
//     equatorial triad of date and the p-vector V(mean) is with
//     respect to the mean equatorial triad of date.
//
//  Called:
//     iauNut80     nutation, IAU 1980
//     iauObl80     mean obliquity, IAU 1980
//     iauNumat     form nutation matrix
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoNutm80 Form the matrix of nutation for a given date, IAU 1980
//  model.
func CgoNutm80(date1, date2 float64) (rmatn [3][3]float64) {
	var cRmatn [3][3]C.double
	C.iauNutm80(C.double(date1), C.double(date2), &cRmatn[0])
	return v3tC2Go(cRmatn)
}

//  GoNutm80 Form the matrix of nutation for a given date, IAU 1980
//  model.
func GoNutm80(date1, date2 float64) (rmatn [3][3]float64) {
	var dpsi, deps, epsa float64

	// Nutation components and mean obliquity.
	dpsi, deps = GoNut80(date1, date2)
	epsa = GoObl80(date1, date2)

	// Build the rotation matrix.
	rmatn = GoNumat(epsa, dpsi, deps)
	return
}
