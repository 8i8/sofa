package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoS2pv Convert position/velocity from spherical to Cartesian
//  coordinates.
//
//  - - - - -
//   S 2 p v
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     theta    double          longitude angle (radians)
//     phi      double          latitude angle (radians)
//     r        double          radial distance
//     td       double          rate of change of theta
//     pd       double          rate of change of phi
//     rd       double          rate of change of r
//
//  Returned:
//     pv       double[2][3]    pv-vector
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoS2pv Convert position/velocity from spherical to Cartesian
//  coordinates.
func CgoS2pv(theta, phi, r, td, pd, rd float64) (pv [2][3]float64) {
	var cPv [2][3]C.double
	C.iauS2pv(C.double(theta), C.double(phi), C.double(r),
		C.double(td), C.double(pd), C.double(rd), &cPv[0])
	return v3dC2Go(cPv)
}

//  GoS2pv Convert position/velocity from spherical to Cartesian
//  coordinates.
func GoS2pv(theta, phi, r, td, pd, rd float64) (pv [2][3]float64) {

	var st, ct, sp, cp, rcp, x, y, rpd, w float64

	st = math.Sin(theta)
	ct = math.Cos(theta)
	sp = math.Sin(phi)
	cp = math.Cos(phi)
	rcp = r * cp
	x = rcp * ct
	y = rcp * st
	rpd = r * pd
	w = rpd*sp - cp*rd

	pv[0][0] = x
	pv[0][1] = y
	pv[0][2] = r * sp
	pv[1][0] = -y*td - w*ct
	pv[1][1] = x*td - w*st
	pv[1][2] = rpd*cp + sp*rd

	return
}
