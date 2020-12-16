package sofa

// #include "sofa.h"
import "C"

//  CgoPvppv Add one pv-vector to another.
//
//  - - - - - -
//   P v p p v
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
//     apb      double[2][3]      a + b
//
//  Note:
//     It is permissible to re-use the same array for any of the
//     arguments.
//
//  Called:
//     iauPpp       p-vector plus p-vector
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPvppv Add one pv-vector to another.
func CgoPvppv(a, b [2][3]float64) (apb [2][3]float64) {
	var cApb [2][3]C.double
	cA, cB := v3dGo2C(a), v3dGo2C(b)
	C.iauPvppv(&cA[0], &cB[0], &cApb[0])
	return v3dC2Go(cApb)
}

//  GoPvppv Add one pv-vector to another.
func GoPvppv(a, b [2][3]float64) (apb [2][3]float64) {
	apb[0] = GoPpp(a[0], b[0])
	apb[1] = GoPpp(a[1], b[1])
	return
}
