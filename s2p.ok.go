package sofa

// #include "sofa.h"
import "C"

//  CgoS2p Convert spherical polar coordinates to p-vector.
//
//  - - - -
//   S 2 p
//  - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     theta   double       longitude angle (radians)
//     phi     double       latitude angle (radians)
//     r       double       radial distance
//
//  Returned:
//     p       double[3]    Cartesian coordinates
//
//  Called:
//     iauS2c       spherical coordinates to unit vector
//     iauSxp       multiply p-vector by scalar
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoS2p Convert spherical polar coordinates to p-vector.
func CgoS2p(theta, phi, r float64) (p [3]float64) {
	var cP [3]C.double
	C.iauS2p(C.double(theta), C.double(phi), C.double(r), &cP[0])
	return v3sC2Go(cP)
}

//  GoS2p Convert spherical polar coordinates to p-vector.
func GoS2p(theta, phi, r float64) (p [3]float64) {

	var u [3]float64

	u = GoS2c(theta, phi)
	p = GoSxp(r, u)
	return
}
