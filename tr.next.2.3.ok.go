package sofa

// #include "sofa.h"
import "C"

//  CgoTr Transpose an r-matrix.
//
//  - - -
//   T r
//  - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     r        double[3][3]    r-matrix
//
//  Returned:
//     rt       double[3][3]    transpose
//
//  Note:
//     It is permissible for r and rt to be the same array.
//
//  Called:
//     iauCr        copy r-matrix
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoTr Transpose an r-matrix.
func CgoTr(r [3][3]float64) (rt [3][3]float64) {
	var cRt [3][3]C.double
	cR := v3tGo2C(r)
	C.iauTr(&cR[0], &cRt[0])
	return v3tC2Go(cRt)
}

// GoTr Transpose an r-matrix.
func GoTr(r [3][3]float64) (rt [3][3]float64) {
	var wm [3][3]float64
	var i, j int

	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			wm[i][j] = r[j][i]
		}
	}
	rt = wm
	return
}
