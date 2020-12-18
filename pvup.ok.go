package sofa

// #include "sofa.h"
import "C"

//  CgoPvup Update a pv-vector, discarding the velocity component.
//
//  - - - - -
//   P v u p
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     dt       double            time interval
//     pv       double[2][3]      pv-vector
//
//  Returned:
//     p        double[3]         p-vector
//
//  Notes:
//
//  1) "Update" means "refer the position component of the vector to a
//     new date dt time units from the existing date".
//
//  2) The time units of dt must match those of the velocity.
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPvup Update a pv-vector, discarding the velocity component.
func CgoPvup(dt float64, pv [2][3]float64) (p [3]float64) {
	var cP [3]C.double
	cPv := v3dGo2C(pv)
	C.iauPvup(C.double(dt), &cPv[0], &cP[0])
	return v3sC2Go(cP)
}

//  GoPvup Update a pv-vector, discarding the velocity component.
func GoPvup(dt float64, pv [2][3]float64) (p [3]float64) {
	p[0] = pv[0][0] + dt*pv[1][0]
	p[1] = pv[0][1] + dt*pv[1][1]
	p[2] = pv[0][2] + dt*pv[1][2]
	return
}
