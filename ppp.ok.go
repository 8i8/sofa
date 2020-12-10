package sofa

// #include "sofa.h"
import "C"

//  CgoPpp P-vector addition.
//
//  - - - -
//   P p p
//  - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     a        double[3]      first p-vector
//     b        double[3]      second p-vector
//
//  Returned:
//     apb      double[3]      a + b
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
//  CgoPpp P-vector addition.
func CgoPpp(a, b [3]float64) (apb [3]float64) {
	var cApb [3]C.double
	cA, cB := v3sGo2C(a), v3sGo2C(b)
	C.iauPpp(&cA[0], &cB[0], &cApb[0])
	return v3sC2Go(cApb)
}

//  GoPpp P-vector addition.
func GoPpp(a, b [3]float64) (apb [3]float64) {

	apb[0] = a[0] + b[0]
	apb[1] = a[1] + b[1]
	apb[2] = a[2] + b[2]
	return
}
