package sofa

// #include "sofa.h"
import "C"

//  CgoPdp p-vector inner (=scalar=dot) product.
//
//  - - - -
//   P d p
//  - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     a      double[3]     first p-vector
//     b      double[3]     second p-vector
//
//  Returned (function value):
//            double        a . b
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPdp p-vector inner (=scalar=dot) product.
func CgoPdp(a, b [3]float64) float64 {
	var cF C.double
	var cA, cB [3]C.double
	cA = v3sGo2C(a)
	cB = v3sGo2C(b)
	cF = C.iauPdp(&cA[0], &cB[0])
	return float64(cF)
}

// GoPdp p-vector inner (=scalar=dot) product.
func GoPdp(a, b [3]float64) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}
