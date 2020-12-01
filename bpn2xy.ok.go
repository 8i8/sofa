package sofa

// #include "sofa.h"
import "C"

//  CgoBpn2xy Extract from the bias-precession-nutation matrix the X,Y
//  coordinates of the Celestial Intermediate Pole.
//
//  - - - - - - -
//   B p n 2 x y
//  - - - - - - -
//
//  Extract from the bias-precession-nutation matrix the X,Y coordinates
//  of the Celestial Intermediate Pole.
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     rbpn      double[3][3]  celestial-to-true matrix (Note 1)
//
//  Returned:
//     x,y       double        Celestial Intermediate Pole (Note 2)
//
//  Notes:
//
//  1) The matrix rbpn transforms vectors from GCRS to true equator (and
//     CIO or equinox) of date, and therefore the Celestial Intermediate
//     Pole unit vector is the bottom row of the matrix.
//
//  2) The arguments x,y are components of the Celestial Intermediate
//     Pole unit vector in the Geocentric Celestial Reference System.
//
//  Reference:
//
//     "Expressions for the Celestial Intermediate Pole and Celestial
//     Ephemeris Origin consistent with the IAU 2000A precession-
//     nutation model", Astron.Astrophys. 400, 1145-1154
//     (2003)
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
//  CgoBpn2xy Extract from the bias-precession-nutation matrix the X,Y
//  coordinates of the Celestial Intermediate Pole.
func CgoBpn2xy(rbpn [3][3]float64) (x, y float64) {
	cX := C.double(x)
	cY := C.double(y)
	cRbpn := v3tGo2C(rbpn)
	C.iauBpn2xy(&cRbpn[0], &cX, &cY)
	return float64(cX), float64(cY)
}

//  GoBpn2xy Extract from the bias-precession-nutation matrix the X,Y
//  coordinates of the Celestial Intermediate Pole.
func GoBpn2xy(rbpn [3][3]float64) (x, y float64) {
	// Extract the X,Y coordinates.
	x = rbpn[2][0]
	y = rbpn[2][1]
	return
}
