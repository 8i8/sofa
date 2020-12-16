package sofa

// #include "sofa.h"
import "C"

//  CgoPvmpv Subtract one pv-vector from another.
//
//  - - - - - -
//   P v m p v
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     a       double[2][3]      first pv-vector
//     b       double[2][3]      second pv-vector
//
//  Returned:
//     amb     double[2][3]      a - b
//
//  Note:
//     It is permissible to re-use the same array for any of the
//     arguments.
//
//  Called:
//     iauPmp       p-vector minus p-vector
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPvmpv Subtract one pv-vector from another.
func CgoPvmpv(a, b [2][3]float64) (amb [2][3]float64) {
	cA, cB := v3dGo2C(a), v3dGo2C(b)
	var cAmb [2][3]C.double
	C.iauPvmpv(&cA[0], &cB[0], &cAmb[0])
	return v3dC2Go(cAmb)
}

//  GoPvmpv Subtract one pv-vector from another.
func GoPvmpv(a, b [2][3]float64) (amb [2][3]float64) {

	amb[0] = GoPmp(a[0], b[0])
	amb[1] = GoPmp(a[1], b[1])
	return
}
