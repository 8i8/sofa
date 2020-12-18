package sofa

// #include "sofa.h"
import "C"

//  CgoPvm Modulus of pv-vector.
//
//  - - - -
//   P v m
//  - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     pv     double[2][3]   pv-vector
//
//  Returned:
//     r      double         modulus of position component
//     s      double         modulus of velocity component
//
//  Called:
//     iauPm        modulus of p-vector
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPvm Modulus of pv-vector.
func CgoPvm(pv [2][3]float64) (r, s float64) {
	var cR, cS C.double
	cPv := v3dGo2C(pv)
	C.iauPvm(&cPv[0], &cR, &cS)
	return float64(cR), float64(cS)
}

//  GoPvm Modulus of pv-vector.
func GoPvm(pv [2][3]float64) (r, s float64) {

	// Distance.
	r = GoPm(pv[0])

	// Speed.
	s = GoPm(pv[1])
	return
}
