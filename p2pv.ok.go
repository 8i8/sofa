package sofa

// #include "sofa.h"
import "C"

//  CgoP2pv Extend a p-vector to a pv-vector by appending a zero
//  velocity.
//
//  - - - - -
//   P 2 p v
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     p        double[3]       p-vector
//
//  Returned:
//     pv       double[2][3]    pv-vector
//
//  Called:
//     iauCp        copy p-vector
//     iauZp        zero p-vector
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoP2pv Extend a p-vector to a pv-vector by appending a zero
//  velocity.
func CgoP2pv(p [3]float64) (pv [2][3]float64) {
	var cPv [2][3]C.double
	cP := v3sGo2C(p)
	C.iauP2pv(&cP[0], &cPv[0])
	return v3dC2Go(cPv)
}

//  GoP2pv Extend a p-vector to a pv-vector by appending a zero
//  velocity.
func GoP2pv(p [3]float64) (pv [2][3]float64) {

	pv[0] = p
	return
}
