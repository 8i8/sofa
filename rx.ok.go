package sofa

// #include "sofa.h"
import "C"
import "math"

//  Rx Rotate an r-matrix about the x-axis.
//
//  - - -
//   R x
//  - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     phi    double          angle (radians)
//
//  Given and returned:
//     r      double[3][3]    r-matrix, rotated
//
//  Notes:
//
//  1) Calling this function with positive phi incorporates in the
//     supplied r-matrix r an additional rotation, about the x-axis,
//     anticlockwise as seen looking towards the origin from positive x.
//
//  2) The additional rotation can be represented by this matrix:
//
//         (  1        0            0      )
//         (                               )
//         (  0   + cos(phi)   + sin(phi)  )
//         (                               )
//         (  0   - sin(phi)   + cos(phi)  )
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoRx Rotate an r-matrix about the x-axis.
func CgoRx(phi float64, r [3][3]float64) (rotated [3][3]float64) {
	cR := v3tGo2C(r)
	C.iauRx(C.double(phi), &cR[0])
	return v3tC2Go(cR)
}

// GoRx Rotate an r-matrix about the x-axis.
func GoRx(phi float64, r [3][3]float64) (rotated [3][3]float64) {
	var s, c, a10, a11, a12, a20, a21, a22 float64

	s = math.Sin(phi)
	c = math.Cos(phi)

	a10 = c*r[1][0] + s*r[2][0]
	a11 = c*r[1][1] + s*r[2][1]
	a12 = c*r[1][2] + s*r[2][2]
	a20 = -s*r[1][0] + c*r[2][0]
	a21 = -s*r[1][1] + c*r[2][1]
	a22 = -s*r[1][2] + c*r[2][2]

	r[1][0] = a10
	r[1][1] = a11
	r[1][2] = a12
	r[2][0] = a20
	r[2][1] = a21
	r[2][2] = a22
	return r
}
