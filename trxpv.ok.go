package sofa

// #include "sofa.h"
import "C"

//  CgoTrxpv Multiply a pv-vector by the transpose of an r-matrix.
//
//  - - - - - -
//   T r x p v
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     r        double[3][3]    r-matrix
//     pv       double[2][3]    pv-vector
//
//  Returned:
//     trpv     double[2][3]    r * pv
//
//  Note:
//     It is permissible for pv and trpv to be the same array.
//
//  Called:
//     iauTr        transpose r-matrix
//     iauRxpv      product of r-matrix and pv-vector
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoTrxpv Multiply a pv-vector by the transpose of an r-matrix.
func CgoTrxpv(r [3][3]float64, pv [2][3]float64) (trpv [2][3]float64) {
	var cTrpv [2][3]C.double
	cR := v3tGo2C(r)
	cPv := v3dGo2C(pv)
	C.iauTrxpv(&cR[0], &cPv[0], &cTrpv[0])
	return v3dC2Go(cTrpv)
}

// GoTrxpv Multiply a pv-vector by the transpose of an r-matrix.
func GoTrxpv(r [3][3]float64, pv [2][3]float64) (trpv [2][3]float64) {
	var tr [3][3]float64

	// Transpose of matrix r.
	tr = GoTr(r)

	// Matrix tr * vector pv -> vector trpv.
	trpv = GoRxpv(tr, pv)

	return
}
