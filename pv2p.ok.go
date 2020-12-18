package sofa

// #include "sofa.h"
import "C"

//  CgoPv2p Discard velocity component of a pv-vector.
//
//  - - - - -
//   P v 2 p
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     pv      double[2][3]     pv-vector
//
//  Returned:
//     p       double[3]        p-vector
//
//  Called:
//     iauCp        copy p-vector
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPv2p Discard velocity component of a pv-vector.
func CgoPv2p(pv [2][3]float64) (p [3]float64) {
	var cP [3]C.double
	cPv := v3dGo2C(pv)
	C.iauPv2p(&cPv[0], &cP[0])
	return v3sC2Go(cP)
}

//  GoPv2p Discard velocity component of a pv-vector.
func GoPv2p(pv [2][3]float64) (p [3]float64) {
	p = pv[0]
	return
}
