package sofa

// #include "sofa.h"
import "C"

//  CgoNumat Form the matrix of nutation.
//
//  - - - - - -
//   N u m a t
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     epsa        double         mean obliquity of date (Note 1)
//     dpsi,deps   double         nutation (Note 2)
//
//  Returned:
//     rmatn       double[3][3]   nutation matrix (Note 3)
//
//  Notes:
//
//
//  1) The supplied mean obliquity epsa, must be consistent with the
//     precession-nutation models from which dpsi and deps were obtained.
//
//  2) The caller is responsible for providing the nutation components;
//     they are in longitude and obliquity, in radians and are with
//     respect to the equinox and ecliptic of date.
//
//  3) The matrix operates in the sense V(true) = rmatn * V(mean),
//     where the p-vector V(true) is with respect to the true
//     equatorial triad of date and the p-vector V(mean) is with
//     respect to the mean equatorial triad of date.
//
//  Called:
//     iauIr        initialize r-matrix to identity
//     iauRx        rotate around X-axis
//     iauRz        rotate around Z-axis
//
//  Reference:
//
//     Explanatory Supplement to the Astronomical Almanac,
//     P. Kenneth Seidelmann (ed), University Science Books (1992),
//     Section 3.222-3 (p114).
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoNumat Form the matrix of nutation.
func CgoNumat(epsa, dpsi, deps float64) (rmatn [3][3]float64) {
	var cRmatn [3][3]C.double
	C.iauNumat(C.double(epsa), C.double(dpsi), C.double(deps), &cRmatn[0])
	return v3tC2Go(cRmatn)
}

//  GoNumat Form the matrix of nutation.
func GoNumat(epsa, dpsi, deps float64) (rmatn [3][3]float64) {

	// Build the rotation matrix.
	rmatn = GoIr()
	rmatn = GoRx(epsa, rmatn)
	rmatn = GoRz(-dpsi, rmatn)
	rmatn = GoRx(-(epsa + deps), rmatn)
	return
}
