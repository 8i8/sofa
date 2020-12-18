package sofa

// #include "sofa.h"
import "C"

//  CgoS2xpv Multiply a pv-vector by two scalars.
//
//  - - - - - -
//   S 2 x p v
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     s1     double         scalar to multiply position component by
//     s2     double         scalar to multiply velocity component by
//     pv     double[2][3]   pv-vector
//
//  Returned:
//     spv    double[2][3]   pv-vector: p scaled by s1, v scaled by s2
//
//  Note:
//     It is permissible for pv and spv to be the same array.
//
//  Called:
//     iauSxp       multiply p-vector by scalar
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoS2xpv Multiply a pv-vector by two scalars.
func CgoS2xpv(s1, s2 float64, pv [2][3]float64) (spv [2][3]float64) {
	var cSpv [2][3]C.double
	cPv := v3dGo2C(pv)
	C.iauS2xpv(C.double(s1), C.double(s2), &cPv[0], &cSpv[0])
	return v3dC2Go(cSpv)
}

//  GoS2xpv Multiply a pv-vector by two scalars.
func GoS2xpv(s1, s2 float64, pv [2][3]float64) (spv [2][3]float64) {

	spv[0] = GoSxp(s1, pv[0])
	spv[1] = GoSxp(s2, pv[1])
	return
}
