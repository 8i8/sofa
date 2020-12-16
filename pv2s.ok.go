package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoPv2s Convert position/velocity from Cartesian to spherical
//  coordinates.
//
//  - - - - -
//   P v 2 s
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     pv       double[2][3]  pv-vector
//
//  Returned:
//     theta    double        longitude angle (radians)
//     phi      double        latitude angle (radians)
//     r        double        radial distance
//     td       double        rate of change of theta
//     pd       double        rate of change of phi
//     rd       double        rate of change of r
//
//  Notes:
//
//  1) If the position part of pv is null, theta, phi, td and pd
//     are indeterminate.  This is handled by extrapolating the
//     position through unit time by using the velocity part of
//     pv.  This moves the origin without changing the direction
//     of the velocity component.  If the position and velocity
//     components of pv are both null, zeroes are returned for all
//     six results.
//
//  2) If the position is a pole, theta, td and pd are indeterminate.
//     In such cases zeroes are returned for all three.
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPv2s Convert position/velocity from Cartesian to spherical
//  coordinates.
func CgoPv2s(pv [2][3]float64) (theta, phi, r, td, pd, rd float64) {
	var cTheta, cPhi, cR, cTd, cPd, cRd C.double
	cPv := v3dGo2C(pv)
	C.iauPv2s(&cPv[0], &cTheta, &cPhi, &cR, &cTd, &cPd, &cRd)
	return float64(cTheta), float64(cPhi), float64(cR),
		float64(cTd), float64(cPd), float64(cRd)
}

//  GoPv2s Convert position/velocity from Cartesian to spherical
//  coordinates.
func GoPv2s(pv [2][3]float64) (theta, phi, r, td, pd, rd float64) {

	var x, y, z, xd, yd, zd, rxy2, rxy, r2, rtrue, rw, xyp float64

	// Components of position/velocity vector. 
	x = pv[0][0]
	y = pv[0][1]
	z = pv[0][2]
	xd = pv[1][0]
	yd = pv[1][1]
	zd = pv[1][2]

	// Component of r in XY plane squared. 
	rxy2 = x*x + y*y

	// Modulus squared. 
	r2 = rxy2 + z*z

	// Modulus. 
	rtrue = math.Sqrt(r2)

	// If null vector, move the origin along the direction of
	// movement. 
	rw = rtrue
	if rtrue == 0.0 {
		x = xd
		y = yd
		z = zd
		rxy2 = x*x + y*y
		r2 = rxy2 + z*z
		rw = math.Sqrt(r2)
	}

	// Position and velocity in spherical coordinates. 
	rxy = math.Sqrt(rxy2)
	xyp = x*xd + y*yd
	if rxy2 != 0.0 {
		theta = math.Atan2(y, x)
		phi = math.Atan2(z, rxy)
		td = (x*yd - y*xd) / rxy2
		pd = (zd*rxy2 - z*xyp) / (r2 * rxy)
	} else {
		theta = 0.0
		if z != 0.0 {
			phi = math.Atan2(z, rxy)
		}
		td = 0.0
		pd = 0.0
	}
	r = rtrue
	if rw != 0.0 {
		rd = (xyp + z*zd) / rw
	}
	return
}
