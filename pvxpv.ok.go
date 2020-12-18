package sofa

// #include "sofa.h"
import "C"

//  CgoPvxpv Outer (=vector=cross) product of two pv-vectors.
//
//  - - - - - -
//   P v x p v
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     a        double[2][3]      first pv-vector
//     b        double[2][3]      second pv-vector
//
//  Returned:
//     axb      double[2][3]      a x b
//
//  Notes:
//
//  1) If the position and velocity components of the two pv-vectors are
//     ( ap, av ) and ( bp, bv ), the result, a x b, is the pair of
//     vectors ( ap x bp, ap x bv + av x bp ).  The two vectors are the
//     cross-product of the two p-vectors and its derivative.
//
//  2) It is permissible to re-use the same array for any of the
//     arguments.
//
//  Called:
//     iauCpv       copy pv-vector
//     iauPxp       vector product of two p-vectors
//     iauPpp       p-vector plus p-vector
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPvxpv Outer (=vector=cross) product of two pv-vectors.
func CgoPvxpv(a, b [2][3]float64) (axb [2][3]float64) {
	var cAxb [2][3]C.double
	cA, cB := v3dGo2C(a), v3dGo2C(b)
	C.iauPvxpv(&cA[0], &cB[0], &cAxb[0])
	return v3dC2Go(cAxb)
}

//  GoPvxpv Outer (=vector=cross) product of two pv-vectors.
func GoPvxpv(a, b [2][3]float64) (axb [2][3]float64) {

	var axbd, adxb [3]float64

	// a x b = position part of result.
	axb[0] = GoPxp(a[0], b[0])

	// a x bdot + adot x b = velocity part of result.
	axbd = GoPxp(a[0], b[1])
	adxb = GoPxp(a[1], b[0])
	axb[1] = GoPpp(axbd, adxb)
	return
}
