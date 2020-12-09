package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoC2s P-vector to spherical coordinates.
//
//  - - - -
//   C 2 s
//  - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     p      double[3]    p-vector
//
//  Returned:
//     theta  double       longitude angle (radians)
//     phi    double       latitude angle (radians)
//
//  Notes:
//
//  1) The vector p can have any magnitude; only its direction is used.
//
//  2) If p is null, zero theta and phi are returned.
//
//  3) At either pole, zero theta is returned.
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoC2s P-vector to spherical coordinates.
func CgoC2s(p [3]float64) (theta, phi float64) {
	var cTheta, cPhi C.double
	cP := v3sGo2C(p)
	C.iauC2s(&cP[0], &cTheta, &cPhi)
	return float64(cTheta), float64(cPhi)
}

//  GoC2s P-vector to spherical coordinates.
func GoC2s(p [3]float64) (theta, phi float64) {

	var x, y, z, d2 float64

	x = p[0]
	y = p[1]
	z = p[2]
	d2 = x*x + y*y

	if d2 == 0.0 {
		theta = 0.0
	} else {
		theta = math.Atan2(y, x)
	}
	if z == 0.0 {
		phi = 0.0
	} else {
		phi = math.Atan2(z, math.Sqrt(d2))
	}

	return

}
