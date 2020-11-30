package sofa

// #include "sofa.h"
import "C"

//  Pn Convert a p-vector into modulus and unit vector.
//
//  - - -
//   P n
//  - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     p        double[3]      p-vector
//
//  Returned:
//     r        double         modulus
//     u        double[3]      unit vector
//
//  Notes:
//
//  1) If p is null, the result is null.  Otherwise the result is a unit
//     vector.
//
//  2) It is permissible to re-use the same array for any of the
//     arguments.
//
//  Called:
//     iauPm        modulus of p-vector
//     iauZp        zero p-vector
//     iauSxp       multiply p-vector by scalar
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
func Pn(p [3]float64) (r float64, u [3]float64) {
	var cU [3]C.double
	var cR C.double
	cP := v3sGo2C(p)
	C.iauPn(&cP[0], &cR, &cU[0])
	return float64(cR), v3sC2Go(cU)
}

// goPn Convert a p-vector into modulus and unit vector.
func goPn(p [3]float64) (r float64, u [3]float64) {

	// Obtain the modulus and test for zero.
	r = Pm(p)
	if r == 0.0 {
		// Null vector.
		return
	}
	// Unit vector.
	u = Sxp(1.0/r, p)
	return
}
