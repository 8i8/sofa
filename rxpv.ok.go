package sofa

// #include "sofa.h"
import "C"

//  CgoRxpv Multiply a pv-vector by an r-matrix.
//
//  - - - - -
//   R x p v
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     r        double[3][3]    r-matrix
//     pv       double[2][3]    pv-vector
//
//  Returned:
//     rpv      double[2][3]    r * pv
//
//  Note:
//     It is permissible for pv and rpv to be the same array.
//
//  Called:
//     iauRxp       product of r-matrix and p-vector
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoRxpv Multiply a pv-vector by an r-matrix.
func CgoRxpv(r [3][3]float64, pv [2][3]float64) (rpv [2][3]float64) {
	var cRpv [2][3]C.double
	cR := v3tGo2C(r)
	cPv := v3dGo2C(pv)
	C.iauRxpv(&cR[0], &cPv[0], &cRpv[0])
	return v3dC2Go(cRpv)
}

//  GoRxpv Multiply a pv-vector by an r-matrix.
func GoRxpv(r [3][3]float64, pv [2][3]float64) (rpv [2][3]float64) {
	rpv[0] = GoRxp(r, pv[0])
	rpv[1] = GoRxp(r, pv[1])
	return
}
