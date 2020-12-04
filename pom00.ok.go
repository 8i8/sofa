package sofa

// #include "sofa.h"
import "C"

//  CgoPom00 Form the matrix of polar motion for a given date, IAU 2000.
//
//  - - - - - - -
//   P o m 0 0
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     xp,yp    double    coordinates of the pole (radians, Note 1)
//     sp       double    the TIO locator s' (radians, Note 2)
//
//  Returned:
//     rpom     double[3][3]   polar-motion matrix (Note 3)
//
//  Notes:
//
//  1) The arguments xp and yp are the coordinates (in radians) of the
//     Celestial Intermediate Pole with respect to the International
//     Terrestrial Reference System (see IERS Conventions 2003),
//     measured along the meridians to 0 and 90 deg west respectively.
//
//  2) The argument sp is the TIO locator s', in radians, which
//     positions the Terrestrial Intermediate Origin on the equator.  It
//     is obtained from polar motion observations by numerical
//     integration, and so is in essence unpredictable.  However, it is
//     dominated by a secular drift of about 47 microarcseconds per
//     century, and so can be taken into account by using s' = -47*t,
//     where t is centuries since J2000.0.  The function iauSp00
//     implements this approximation.
//
//  3) The matrix operates in the sense V(TRS) = rpom * V(CIP), meaning
//     that it is the final rotation when computing the pointing
//     direction to a celestial source.
//
//  Called:
//     iauIr        initialize r-matrix to identity
//     iauRz        rotate around Z-axis
//     iauRy        rotate around Y-axis
//     iauRx        rotate around X-axis
//
//  Reference:
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
//  CgoPom00 Form the matrix of polar motion for a given date, IAU 2000.
func CgoPom00(xp, yp, sp float64) (rpom [3][3]float64) {
	var cRpom [3][3]C.double
	C.iauPom00(C.double(xp), C.double(yp), C.double(sp), &cRpom[0])
	return v3tC2Go(cRpom)
}

// GoPom00 Form the matrix of polar motion for a given date, IAU 2000.
func GoPom00(xp, yp, sp float64) (rpom [3][3]float64) {

	// Construct the matrix.
	rpom = GoIr()
	rpom = GoRz(sp, rpom)
	rpom = GoRy(-xp, rpom)
	rpom = GoRx(-yp, rpom)
	return
}
