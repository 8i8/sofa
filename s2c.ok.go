package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoS2c Convert spherical coordinates to Cartesian.
//
//  - - - -
//   S 2 c
//  - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     theta    double       longitude angle (radians)
//     phi      double       latitude angle (radians)
//
//  Returned:
//     c        double[3]    direction cosines
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoS2c Convert spherical coordinates to Cartesian.
func CgoS2c(theta, phi float64) (c [3]float64) {
	var cC [3]C.double
	C.iauS2c(C.double(theta), C.double(phi), &cC[0])
	return v3sC2Go(cC)
}

//  GoS2c Convert spherical coordinates to Cartesian.
func GoS2c(theta, phi float64) (c [3]float64) {
	var cp float64
	cp = math.Cos(phi)
	c[0] = math.Cos(theta) * cp
	c[1] = math.Sin(theta) * cp
	c[2] = math.Sin(phi)
	return
}
