package sofa

// #include "sofa.h"
import "C"

//  CgoPvu Update a pv-vector.
//
//  - - - -
//   P v u
//  - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     dt       double           time interval
//     pv       double[2][3]     pv-vector
//
//  Returned:
//     upv      double[2][3]     p updated, v unchanged
//
//  Notes:
//
//  1) "Update" means "refer the position component of the vector
//     to a new date dt time units from the existing date".
//
//  2) The time units of dt must match those of the velocity.
//
//  3) It is permissible for pv and upv to be the same array.
//
//  Called:
//     iauPpsp      p-vector plus scaled p-vector
//     iauCp        copy p-vector
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPvu Update a pv-vector.
func CgoPvu(dt float64, pv [2][3]float64) (upv [2][3]float64) {
	var cUpv [2][3]C.double
	cPv := v3dGo2C(pv)
	C.iauPvu(C.double(dt), &cPv[0], &cUpv[0])
	return v3dC2Go(cUpv)
}

//  GoPvu Update a pv-vector.
func GoPvu(dt float64, pv [2][3]float64) (upv [2][3]float64) {

	upv[0] = GoPpsp(pv[0], dt, pv[1])
	upv[1] = pv[1]
	return
}
