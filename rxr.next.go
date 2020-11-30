package sofa

// #include "sofa.h"
import "C"

//  Rxt Multiply two r-matrices.
//
//  - - - -
//   R x r
//  - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     a        double[3][3]    first r-matrix
//     b        double[3][3]    second r-matrix
//
//  Returned:
//     atb      double[3][3]    a * b
//
//  Note:
//     It is permissible to re-use the same array for any of the
//     arguments.
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
func Rxr(a, b [3][3]float64) (atb [3][3]float64) {
	var cAtb [3][3]C.double
	cA := v3tGo2C(a)
	cB := v3tGo2C(b)
	C.iauRxr(&cA[0], &cB[0], &cAtb[0])
	return v3tC2Go(cAtb)
}

// goRxt Multiply two r-matrices.
func goRxr(a, b [3][3]float64) (atb [3][3]float64) {
	var i, j, k int
	var w float64

	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			w = 0.0
			for k = 0; k < 3; k++ {
				w += a[i][k] * b[k][j]
			}
			atb[i][j] = w
		}
	}
	return
}
