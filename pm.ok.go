package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoPm Modulus of p-vector.
//
//  - - -
//   P m
//  - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     p      double[3]     p-vector
//
//  Returned (function value):
//            double        modulus
//
//  This revision:  2013 August 7
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPm Modulus of p-vector.
func CgoPm(p [3]float64) float64 {
	var cP [3]C.double
	var cF C.double
	cP = v3sGo2C(p)
	cF = C.iauPm(&cP[0])
	return float64(cF)
}

// GoPm Modulus of p-vector.
func GoPm(p [3]float64) float64 {
	return math.Sqrt(p[0]*p[0] + p[1]*p[1] + p[2]*p[2])
}
