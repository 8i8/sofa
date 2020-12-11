package sofa

// #include "sofa.h"
import "C"

//  CgoC2i00b Form the celestial-to-intermediate matrix for a given date
//  using the IAU 2000B precession-nutation model.
//
//  - - - - - - -
//   C 2 i 0 0 b
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     date1,date2 double       TT as a 2-part Julian Date (Note 1)
//
//  Returned:
//     rc2i        double[3][3] celestial-to-intermediate matrix (Note 2)
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
//  2) The matrix rc2i is the first stage in the transformation from
//     celestial to terrestrial coordinates:
//
//        [TRS]  =  RPOM * R_3(ERA) * rc2i * [CRS]
//
//               =  rc2t * [CRS]
//
//     where [CRS] is a vector in the Geocentric Celestial Reference
//     System and [TRS] is a vector in the International Terrestrial
//     Reference System (see IERS Conventions 2003), ERA is the Earth
//     Rotation Angle and RPOM is the polar motion matrix.
//
//  3) The present function is faster, but slightly less accurate (about
//     1 mas), than the iauC2i00a function.
//
//  Called:
//     iauPnm00b    classical NPB matrix, IAU 2000B
//     iauC2ibpn    celestial-to-intermediate matrix, given NPB matrix
//
//  References:
//
//     "Expressions for the Celestial Intermediate Pole and Celestial
//     Ephemeris Origin consistent with the IAU 2000A precession-
//     nutation model", Astron.Astrophys. 400, 1145-1154
//     (2003)
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
//  CgoC2i00b Form the celestial-to-intermediate matrix for a given date
//  using the IAU 2000B precession-nutation model.
func CgoC2i00b(date1, date2 float64) (rc2i [3][3]float64) {
	var cRc2i [3][3]C.double
	C.iauC2i00b(C.double(date1), C.double(date2), &cRc2i[0])
	return v3tC2Go(cRc2i)
}

//  GoC2i00b Form the celestial-to-intermediate matrix for a given date
//  using the IAU 2000B precession-nutation model.
func GoC2i00b(date1, date2 float64) (rc2i [3][3]float64) {
	var rbpn [3][3]float64

	// Obtain the celestial-to-true matrix (IAU 2000B).
	rbpn = GoPnm00b(date1, date2)

	// Form the celestial-to-intermediate matrix.
	rc2i = GoC2ibpn(date1, date2, rbpn)
	return
}
