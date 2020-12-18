package sofa

// #include "sofa.h"
import "C"

//  CgoPvdpv Inner (=scalar=dot) product of two pv-vectors.
//
//  - - - - - -
//   P v d p v
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
//     adb      double[2]         a . b (see note)
//
//  Note:
//
//     If the position and velocity components of the two pv-vectors are
//     ( ap, av ) and ( bp, bv ), the result, a . b, is the pair of
//     numbers ( ap . bp , ap . bv + av . bp ).  The two numbers are the
//     dot-product of the two p-vectors and its derivative.
//
//  Called:
//     iauPdp       scalar product of two p-vectors
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPvdpv Inner (=scalar=dot) product of two pv-vectors.
func CgoPvdpv(a, b [2][3]float64) (adb [2]float64) {
	var cAdb [2]C.double
	cA, cB := v3dGo2C(a), v3dGo2C(b)
	C.iauPvdpv(&cA[0], &cB[0], &cAdb[0])
	return v2sC2Go(cAdb)
}

//  GoPvdpv Inners(=scalar=dot) product of two pv-vectors.
func GoPvdpv(a, b [2][3]float64) (adb [2]float64) {
	var adbd, addb float64

	// a . b = constant part of result.
	adb[0] = GoPdp(a[0], b[0])

	// a . bdot
	adbd = GoPdp(a[0], b[1])

	// adot . b
	addb = GoPdp(a[1], b[0])

	// Velocity part of result.
	adb[1] = adbd + addb
	return
}
