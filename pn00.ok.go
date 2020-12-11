package sofa

// #include "sofa.h"
import "C"

//  CgoPn00 Precession-nutation, IAU 2000 model:  a multi-purpose
//  function, supporting classical (equinox-based) use directly and
//  CIO-based use indirectly.
//
//  - - - - -
//   P n 0 0
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     date1,date2  double          TT as a 2-part Julian Date (Note 1)
//     dpsi,deps    double          nutation (Note 2)
//
//  Returned:
//     epsa         double          mean obliquity (Note 3)
//     rb           double[3][3]    frame bias matrix (Note 4)
//     rp           double[3][3]    precession matrix (Note 5)
//     rbp          double[3][3]    bias-precession matrix (Note 6)
//     rn           double[3][3]    nutation matrix (Note 7)
//     rbpn         double[3][3]    GCRS-to-true matrix (Note 8)
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
//  2) The caller is responsible for providing the nutation components;
//     they are in longitude and obliquity, in radians and are with
//     respect to the equinox and ecliptic of date.  For high-accuracy
//     applications, free core nutation should be included as well as
//     any other relevant corrections to the position of the CIP.
//
//  3) The returned mean obliquity is consistent with the IAU 2000
//     precession-nutation models.
//
//  4) The matrix rb transforms vectors from GCRS to J2000.0 mean
//     equator and equinox by applying frame bias.
//
//  5) The matrix rp transforms vectors from J2000.0 mean equator and
//     equinox to mean equator and equinox of date by applying
//     precession.
//
//  6) The matrix rbp transforms vectors from GCRS to mean equator and
//     equinox of date by applying frame bias then precession.  It is
//     the product rp x rb.
//
//  7) The matrix rn transforms vectors from mean equator and equinox of
//     date to true equator and equinox of date by applying the nutation
//     (luni-solar + planetary).
//
//  8) The matrix rbpn transforms vectors from GCRS to true equator and
//     equinox of date.  It is the product rn x rbp, applying frame
//     bias, precession and nutation in that order.
//
//  9) It is permissible to re-use the same array in the returned
//     arguments.  The arrays are filled in the order given.
//
//  Called:
//     iauPr00      IAU 2000 precession adjustments
//     iauObl80     mean obliquity, IAU 1980
//     iauBp00      frame bias and precession matrices, IAU 2000
//     iauCr        copy r-matrix
//     iauNumat     form nutation matrix
//     iauRxr       product of two r-matrices
//
//  Reference:
//
//     Capitaine, N., Chapront, J., Lambert, S. and Wallace, P.,
//     "Expressions for the Celestial Intermediate Pole and Celestial
//     Ephemeris Origin consistent with the IAU 2000A precession-
//     nutation model", Astron.Astrophys. 400, 1145-1154 (2003)
//
//     n.b. The celestial ephemeris origin (CEO) was renamed "celestial
//          intermediate origin" (CIO) by IAU 2006 Resolution 2.
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPn00 Precession-nutation, IAU 2000 model:  a multi-purpose
//  function, supporting classical (equinox-based) use directly and
//  CIO-based use indirectly.
func CgoPn00(date1, date2, dpsi, deps float64) (epsa float64, rb, rp,
	rbp, rn, rbpn [3][3]float64) {
	var cEpsa C.double
	var cRb, cRp, cRbp, cRn, cRbpn [3][3]C.double
	C.iauPn00(C.double(date1), C.double(date2), C.double(dpsi),
		C.double(deps), &cEpsa, &cRb[0], &cRp[0], &cRbp[0],
		&cRn[0], &cRbpn[0])
	return float64(cEpsa), v3tC2Go(cRb), v3tC2Go(cRp),
		v3tC2Go(cRbp), v3tC2Go(cRn), v3tC2Go(cRbpn)
}

//  GoPn00 Precession-nutation, IAU 2000 model:  a multi-purpose
//  function, supporting classical (equinox-based) use directly and
//  CIO-based use indirectly.
func GoPn00(date1, date2, dpsi, deps float64) (epsa float64, rb, rp,
	rbp, rn, rbpn [3][3]float64) {

	var depspr float64

	// IAU 2000 precession-rate adjustments. 
	_, depspr = GoPr00(date1, date2)

	// Mean obliquity, consistent with IAU 2000 precession-nutation. 
	epsa = GoObl80(date1, date2) + depspr

	// Frame bias and precession matrices and their product. 
	rb, rp, rbp = GoBp00(date1, date2)

	// Nutation matrix. 
	rn = GoNumat(epsa, dpsi, deps)

	// Bias-precession-nutation matrix (classical). 
	rbpn = GoRxr(rn, rbp)
	return
}
