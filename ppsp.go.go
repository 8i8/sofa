package sofa

// #include "sofa.h"
import "C"

//  CgoPpsp P-vector plus scaled p-vector.
//
//  - - - - -
//   P p s p
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     a      double[3]     first p-vector
//     s      double        scalar (multiplier for b)
//     b      double[3]     second p-vector
//
//  Returned:
//     apsb   double[3]     a + s*b
//
//  Note:
//     It is permissible for any of a, b and apsb to be the same array.
//
//  Called:
//     iauSxp       multiply p-vector by scalar
//     iauPpp       p-vector plus p-vector
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPpsp P-vector plus scaled p-vector.
func CgoPpsp(a [3]float64, s float64, b [3]float64) (apsb [3]float64) {
	var cApsb [3]C.double
	cA, cB := v3sGo2C(a), v3sGo2C(b)
	C.iauPpsp(&cA[0], C.double(s), &cB[0], &cApsb[0])
	return v3sC2Go(cApsb)
}

//  GoPpsp P-vector plus scaled p-vector.
func GoPpsp(a [3]float64, s float64, b [3]float64) (apsb [3]float64) {

	var sb [3]float64

	// s*b.
	sb = GoSxp(s, b)

	// a + s*b.
	apsb = GoPpp(a, sb)

	return
}
