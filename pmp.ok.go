package sofa

// #include "sofa.h"
import "C"

//  CgoPmp P-vector subtraction.
//
//  - - - -
//   P m p
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
//     amb      double[3]      a - b
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
//  CgoPmp P-vector subtraction.
func CgoPmp(a, b [3]float64) (amb [3]float64) {
	var cAmb [3]C.double
	cA, cB := v3sGo2C(a), v3sGo2C(b)
	C.iauPmp(&cA[0], &cB[0], &cAmb[0])
	return v3sC2Go(cAmb)
}

//  GoPmp P-vector subtraction.
func GoPmp(a, b [3]float64) (amb [3]float64) {

	amb[0] = a[0] - b[0]
	amb[1] = a[1] - b[1]
	amb[2] = a[2] - b[2]

	return
}
