package sofa

// #include "sofa.h"
import "C"

//  CgoSxpv Multiply a pv-vector by a scalar.
//
//  - - - - -
//   S x p v
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     s       double          scalar
//     pv      double[2][3]    pv-vector
//
//  Returned:
//     spv     double[2][3]    s * pv
//
//  Note:
//     It is permissible for pv and spv to be the same array
//
//  Called:
//     iauS2xpv     multiply pv-vector by two scalars
//
//  This revision:  2013 August 7
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoSxpv Multiply a pv-vector by a scalar.
func CgoSxpv(s float64, pv [2][3]float64) (spv [2][3]float64) {
	var cSpv [2][3]C.double
	cPv := v3dGo2C(pv)
	C.iauSxpv(C.double(s), &cPv[0], &cSpv[0])
	return v3dC2Go(cSpv)
}

//  GoSxpv Multiply a pv-vector by a scalar.
func GoSxpv(s float64, pv [2][3]float64) (spv [2][3]float64) {

	spv = GoS2xpv(s, s, pv)
	return
}
