package sofa

// #include "sofa.h"
import "C"

//  CgoPn06a Precession-nutation, IAU 2006/2000A models:  a
//  multi-purpose function, supporting classical (equinox-based) use
//  directly and CIO-based use indirectly.
//
//  - - - - - -
//   P n 0 6 a
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     date1,date2  double          TT as a 2-part Julian Date (Note 1)
//
//  Returned:
//     dpsi,deps    double          nutation (Note 2)
//     epsa         double          mean obliquity (Note 3)
//     rb           double[3][3]    frame bias matrix (Note 4)
//     rp           double[3][3]    precession matrix (Note 5)
//     rbp          double[3][3]    bias-precession matrix (Note 6)
//     rn           double[3][3]    nutation matrix (Note 7)
//     rbpn         double[3][3]    GCRS-to-true matrix (Notes 8,9)
//
//  Notes:
//
//  1)  The TT date date1+date2 is a Julian Date, apportioned in any
//      convenient way between the two arguments.  For example,
//      JD(TT)=2450123.7 could be expressed in any of these ways,
//      among others:
//
//             date1          date2
//
//          2450123.7           0.0       (JD method)
//          2451545.0       -1421.3       (J2000 method)
//          2400000.5       50123.2       (MJD method)
//          2450123.5           0.2       (date & time method)
//
//      The JD method is the most natural and convenient to use in
//      cases where the loss of several decimal digits of resolution
//      is acceptable.  The J2000 method is best matched to the way
//      the argument is handled internally and will deliver the
//      optimum resolution.  The MJD method and the date & time methods
//      are both good compromises between resolution and convenience.
//
//  2)  The nutation components (luni-solar + planetary, IAU 2000A) in
//      longitude and obliquity are in radians and with respect to the
//      equinox and ecliptic of date.  Free core nutation is omitted;
//      for the utmost accuracy, use the iauPn06 function, where the
//      nutation components are caller-specified.
//
//  3)  The mean obliquity is consistent with the IAU 2006 precession.
//
//  4)  The matrix rb transforms vectors from GCRS to mean J2000.0 by
//      applying frame bias.
//
//  5)  The matrix rp transforms vectors from mean J2000.0 to mean of
//      date by applying precession.
//
//  6)  The matrix rbp transforms vectors from GCRS to mean of date by
//      applying frame bias then precession.  It is the product rp x rb.
//
//  7)  The matrix rn transforms vectors from mean of date to true of
//      date by applying the nutation (luni-solar + planetary).
//
//  8)  The matrix rbpn transforms vectors from GCRS to true of date
//      (CIP/equinox).  It is the product rn x rbp, applying frame bias,
//      precession and nutation in that order.
//
//  9)  The X,Y,Z coordinates of the IAU 2006/2000A Celestial
//      Intermediate Pole are elements (3,1-3) of the GCRS-to-true
//      matrix, i.e. rbpn[2][0-2].
//
//  10) It is permissible to re-use the same array in the returned
//      arguments.  The arrays are filled in the stated order.
//
//  Called:
//     iauNut06a    nutation, IAU 2006/2000A
//     iauPn06      bias/precession/nutation results, IAU 2006
//
//  Reference:
//
//     Capitaine, N. & Wallace, P.T., 2006, Astron.Astrophys. 450, 855
//
//  This revision:  2013 November 13
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPn06a Precession-nutation, IAU 2006/2000A models:  a
//  multi-purpose function, supporting classical (equinox-based) use
//  directly and CIO-based use indirectly.
func CgoPn06a(date1, date2 float64) (dpsi, deps, epsa float64,
	rb, rp, rbp, rn, rbpn [3][3]float64) {
	var cDpsi, cDeps, cEps C.double
	var cRb, cRp, cRbp, cRn, cRbpn [3][3]C.double
	C.iauPn06a(C.double(date1), C.double(date2),
		&cDpsi, &cDeps, &cEps,
		&cRb[0], &cRp[0], &cRbp[0], &cRn[0], &cRbpn[0])
	return float64(cDpsi), float64(cDeps), float64(cEps),
		v3tC2Go(cRb), v3tC2Go(cRp), v3tC2Go(cRbp), v3tC2Go(cRn),
		v3tC2Go(cRbpn)
}

//  GoPn06a Precession-nutation, IAU 2006/2000A models:  a
//  multi-purpose function, supporting classical (equinox-based) use
//  directly and CIO-based use indirectly.
func GoPn06a(date1, date2 float64) (dpsi, deps, epsa float64,
	rb, rp, rbp, rn, rbpn [3][3]float64) {

	// Nutation.
	dpsi, deps = GoNut06a(date1, date2)

	// Remaining results.
	epsa, rb, rp, rbp, rn, rbpn = GoPn06(date1, date2, dpsi, deps)
	return
}
