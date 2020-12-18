package sofa

// #include "sofa.h"
import "C"

//  CgoPnm80 Form the matrix of precession/nutation for a given date,
//  IAU 1976 precession model, IAU 1980 nutation model.
//
//  - - - - - -
//   P n m 8 0
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     date1,date2    double         TDB date (Note 1)
//
//  Returned:
//     rmatpn         double[3][3]   combined precession/nutation matrix
//
//  Notes:
//
//  1) The TDB date date1+date2 is a Julian Date, apportioned in any
//     convenient way between the two arguments.  For example,
//     JD(TDB)=2450123.7 could be expressed in any of these ways,
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
//  2) The matrix operates in the sense V(date) = rmatpn * V(J2000),
//     where the p-vector V(date) is with respect to the true equatorial
//     triad of date date1+date2 and the p-vector V(J2000) is with
//     respect to the mean equatorial triad of epoch J2000.0.
//
//  Called:
//     iauPmat76    precession matrix, IAU 1976
//     iauNutm80    nutation matrix, IAU 1980
//     iauRxr       product of two r-matrices
//
//  Reference:
//
//     Explanatory Supplement to the Astronomical Almanac,
//     P. Kenneth Seidelmann (ed), University Science Books (1992),
//     Section 3.3 (p145).
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPnm80 Form the matrix of precession/nutation for a given date,
//  IAU 1976 precession model, IAU 1980 nutation model.
func CgoPnm80(date1, date2 float64) (rmatpn [3][3]float64) {
	var cRmatpn [3][3]C.double
	C.iauPnm80(C.double(date1), C.double(date2), &cRmatpn[0])
	return v3tC2Go(cRmatpn)
}

//  GoPnm80 Form the matrix of precession/nutation for a given date,
//  IAU 1976 precession model, IAU 1980 nutation model.
func GoPnm80(date1, date2 float64) (rmatpn [3][3]float64) {
	var rmatp, rmatn [3][3]float64

	// Precession matrix, J2000.0 to date.
	rmatp = GoPmat76(date1, date2)

	// Nutation matrix.
	rmatn = GoNutm80(date1, date2)

	// Combine the matrices:  PN = N x P.
	rmatpn = GoRxr(rmatn, rmatp)
	return
}
