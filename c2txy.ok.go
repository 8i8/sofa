package sofa

// #include "sofa.h"
import "C"

//  CgoC2txy Form the celestial to terrestrial matrix given the date,
//  the UT1, the CIP coordinates and the polar motion.  IAU 2000.
//
//  - - - - - -
//   C 2 t x y
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     tta,ttb  double         TT as a 2-part Julian Date (Note 1)
//     uta,utb  double         UT1 as a 2-part Julian Date (Note 1)
//     x,y      double         Celestial Intermediate Pole (Note 2)
//     xp,yp    double         coordinates of the pole (radians, Note 3)
//
//  Returned:
//     rc2t     double[3][3]   celestial-to-terrestrial matrix (Note 4)
//
//  Notes:
//
//  1) The TT and UT1 dates tta+ttb and uta+utb are Julian Dates,
//     apportioned in any convenient way between the arguments uta and
//     utb.  For example, JD(UT1)=2450123.7 could be expressed in any o
//     these ways, among others:
//
//             uta            utb
//
//         2450123.7           0.0       (JD method)
//         2451545.0       -1421.3       (J2000 method)
//         2400000.5       50123.2       (MJD method)
//         2450123.5           0.2       (date & time method)
//
//     The JD method is the most natural and convenient to use in
//     cases where the loss of several decimal digits of resolution is
//     acceptable.  The J2000 and MJD methods are good compromises
//     between resolution and convenience.  In the case of uta,utb, the
//     date & time method is best matched to the Earth rotation angle
//     algorithm used:  maximum precision is delivered when the uta
//     argument is for 0hrs UT1 on the day in question and the utb
//     argument lies in the range 0 to 1, or vice versa.
//
//  2) The Celestial Intermediate Pole coordinates are the x,y
//     components of the unit vector in the Geocentric Celestial
//     Reference System.
//
//  3) The arguments xp and yp are the coordinates (in radians) of the
//     Celestial Intermediate Pole with respect to the International
//     Terrestrial Reference System (see IERS Conventions 2003),
//     measured along the meridians to 0 and 90 deg west respectively.
//
//  4) The matrix rc2t transforms from celestial to terrestrial
//     coordinates:
//
//        [TRS] = RPOM * R_3(ERA) * RC2I * [CRS]
//
//              = rc2t * [CRS]
//
//     where [CRS] is a vector in the Geocentric Celestial Reference
//     System and [TRS] is a vector in the International Terrestrial
//     Reference System (see IERS Conventions 2003), ERA is the Earth
//     Rotation Angle and RPOM is the polar motion matrix.
//
//  5) Although its name does not include "00", This function is in fact
//     specific to the IAU 2000 models.
//
//  Called:
//     iauC2ixy     celestial-to-intermediate matrix, given X,Y
//     iauEra00     Earth rotation angle, IAU 2000
//     iauSp00      the TIO locator s', IERS 2000
//     iauPom00     polar motion matrix
//     iauC2tcio    form CIO-based celestial-to-terrestrial matrix
//
// Reference:
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
//  CgoC2txy Form the celestial to terrestrial matrix given the date,
//  the UT1, the CIP coordinates and the polar motion.  IAU 2000.
func CgoC2txy(tta, ttb, uta, utb, x, y, xp, yp float64) (
	rc2t [3][3]float64) {
	var cRc2t [3][3]C.double
	C.iauC2txy(C.double(tta), C.double(ttb), C.double(uta),
		C.double(utb), C.double(x), C.double(y), C.double(xp),
		C.double(yp), &cRc2t[0])
	return v3tC2Go(cRc2t)
}

//  GoC2txy Form the celestial to terrestrial matrix given the date,
//  the UT1, the CIP coordinates and the polar motion.  IAU 2000.
func GoC2txy(tta, ttb, uta, utb, x, y, xp, yp float64) (
	rc2t [3][3]float64) {
	var rc2i, rpom [3][3]float64
	var era, sp float64

	// Form the celestial-to-intermediate matrix for this TT. 
	rc2i = GoC2ixy(tta, ttb, x, y)

	// Predict the Earth rotation angle for this UT1. 
	era = GoEra00(uta, utb)

	// Estimate s'. 
	sp = GoSp00(tta, ttb)

	// Form the polar motion matrix. 
	rpom = GoPom00(xp, yp, sp)

	// Combine to form the celestial-to-terrestrial matrix. 
	rc2t = GoC2tcio(rc2i, era, rpom)
	return
}
