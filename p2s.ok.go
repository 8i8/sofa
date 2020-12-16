package sofa

// #include "sofa.h"
import "C"

//  CgoP2s P-vector to spherical polar coordinates.
//
//  - - - -
//   P 2 s
//  - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     p        double[3]    p-vector
//
//  Returned:
//     theta    double       longitude angle (radians)
//     phi      double       latitude angle (radians)
//     r        double       radial distance
//
//  Notes:
//
//  1) If P is null, zero theta, phi and r are returned.
//
//  2) At either pole, zero theta is returned.
//
//  Called:
//     iauC2s       p-vector to spherical
//     iauPm        modulus of p-vector
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoP2s P-vector to spherical polar coordinates.
func CgoP2s(p [3]float64) (theta, phi, r float64) {
	var cTheta, cPhi, cR C.double
	cP := v3sGo2C(p)
	C.iauP2s(&cP[0], &cTheta, &cPhi, &cR)
	return float64(cTheta), float64(cPhi), float64(cR)
}

//  GoP2s P-vector to spherical polar coordinates.
func GoP2s(p [3]float64) (theta, phi, r float64) {

	theta, phi = GoC2s(p)
	r = GoPm(p)
	return
}
