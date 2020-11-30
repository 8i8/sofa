package sofa

// #include "sofa.h"
import "C"
import "math"

//  Rz Rotate an r-matrix about the z-axis.
//
//  - - -
//   R z
//  - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     psi    double          angle (radians)
//
//  Given and returned:
//     r      double[3][3]    r-matrix, rotated
//
//  Notes:
//
//  1) Calling this function with positive psi incorporates in the
//     supplied r-matrix r an additional rotation, about the z-axis,
//     anticlockwise as seen looking towards the origin from positive z.
//
//  2) The additional rotation can be represented by this matrix:
//
//         (  + cos(psi)   + sin(psi)     0  )
//         (                                 )
//         (  - sin(psi)   + cos(psi)     0  )
//         (                                 )
//         (       0            0         1  )
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
func Rz(psi float64, r [3][3]float64) (rotated [3][3]float64) {
	cR := v3tGo2C(r)
	C.iauRz(C.double(psi), &cR[0])
	return v3tC2Go(cR)
}

// goRz Rotate an r-matrix about the z-axis.
func goRz(psi float64, r [3][3]float64) (rotated [3][3]float64) {
	var s, c, a00, a01, a02, a10, a11, a12 float64

	s = math.Sin(psi)
	c = math.Cos(psi)

	a00 = c*r[0][0] + s*r[1][0]
	a01 = c*r[0][1] + s*r[1][1]
	a02 = c*r[0][2] + s*r[1][2]
	a10 = -s*r[0][0] + c*r[1][0]
	a11 = -s*r[0][1] + c*r[1][1]
	a12 = -s*r[0][2] + c*r[1][2]

	r[0][0] = a00
	r[0][1] = a01
	r[0][2] = a02
	r[1][0] = a10
	r[1][1] = a11
	r[1][2] = a12
	return r
}
