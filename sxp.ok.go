package sofa

// #include "sofa.h"
import "C"

//  CgoSxp Multiply a p-vector by a scalar.
//
//  - - - -
//   S x p
//  - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     s      double        scalar
//     p      double[3]     p-vector
//
//  Returned:
//     sp     double[3]     s * p
//
//  Note:
//     It is permissible for p and sp to be the same array.
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoSxp Multiply a p-vector by a scalar.
func CgoSxp(s float64, p [3]float64) (sp [3]float64) {
	var cP = v3sGo2C(p)
	var cSp [3]C.double
	C.iauSxp(C.double(s), &cP[0], &cSp[0])
	return v3sC2Go(cSp)
}

// GoSxp Multiply a p-vector by a scalar.
func GoSxp(s float64, p [3]float64) (sp [3]float64) {
	sp[0] = s * p[0]
	sp[1] = s * p[1]
	sp[2] = s * p[2]
	return
}
