package sofa

// #include "sofa.h"
import "C"

//  CgoTrxp Multiply a p-vector by the transpose of an r-matrix.
//
//  - - - - -
//   T r x p
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     r        double[3][3]   r-matrix
//     p        double[3]      p-vector
//
//  Returned:
//     trp      double[3]      r^T * p
//
//  Note:
//     It is permissible for p and trp to be the same array.
//
//  Called:
//     iauTr        transpose r-matrix
//     iauRxp       product of r-matrix and p-vector
//
//  This revision:  2020 May 24
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoTrxp Multiply a p-vector by the transpose of an r-matrix.
// void iauTrxp(double r[3][3], double p[3], double trp[3])
func CgoTrxp(r [3][3]float64, p [3]float64) (trp [3]float64) {
	var cTrp [3]C.double
	cR := v3tGo2C(r)
	cP := v3sGo2C(p)
	C.iauTrxp(&cR[0], &cP[0], &cTrp[0])
	return v3sC2Go(cTrp)
}

// GoTrxp Multiply a p-vector by the transpose of an r-matrix.
func GoTrxp(r [3][3]float64, p [3]float64) (trp [3]float64) {

	// Transpose of matrix r.
	tr := GoTr(r)

	// Matrix tr * vector p -> vector trp.
	trp = GoRxp(tr, p)

	return
}
