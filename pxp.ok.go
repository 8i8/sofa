package sofa

// #include "sofa.h"
import "C"

//  CgoPxp p-vector outer (=vector=cross) product.
//
//  - - - -
//   P x p
//  - - - -
//
//  p-vector outer (=vector=cross) product.
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
//     axb      double[3]      a x b
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
//  CgoPxp p-vector outer (=vector=cross) product.
func CgoPxp(a, b [3]float64) (axb [3]float64) {
	var cAxb [3]C.double
	cA, cB := v3sGo2C(a), v3sGo2C(b)
	C.iauPxp(&cA[0], &cB[0], &cAxb[0])
	return v3sC2Go(cAxb)
}

//  GoPxp p-vector outer (=vector=cross) product.
func GoPxp(a, b [3]float64) (axb [3]float64) {

	var xa, ya, za, xb, yb, zb float64

	xa = a[0]
	ya = a[1]
	za = a[2]
	xb = b[0]
	yb = b[1]
	zb = b[2]
	axb[0] = ya*zb - za*yb
	axb[1] = za*xb - xa*zb
	axb[2] = xa*yb - ya*xb

	return
}
