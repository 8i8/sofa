package sofa

// #include "sofa.h"
import "C"

//  CgoRxp Multiply a p-vector by an r-matrix.
//
//  - - - -
//   R x p
//  - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     r        double[3][3]    r-matrix
//     p        double[3]       p-vector
//
//  Returned:
//     rp       double[3]       r * p
//
//  Note:
//     It is permissible for p and rp to be the same array.
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
//  CgoRxp Multiply a p-vector by an r-matrix.
func CgoRxp(r [3][3]float64, p [3]float64) (rp [3]float64) {
	var cRp [3]C.double
	cR := v3tGo2C(r)
	cP := v3sGo2C(p)
	C.iauRxp(&cR[0], &cP[0], &cRp[0])
	return v3sC2Go(cRp)
}

// GoRxp Multiply a p-vector by an r-matrix.
func GoRxp(r [3][3]float64, p [3]float64) (rp [3]float64) {
	var w float64
	var wrp [3]float64
	var i, j int

	// Matrix r * vector p.
	for j = 0; j < 3; j++ {
		w = 0.0
		for i = 0; i < 3; i++ {
			w += r[j][i] * p[i]
		}
		wrp[j] = w
	}

	// Return the result.
	rp = wrp
	return
}
