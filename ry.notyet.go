package sofa

import "math"

//  Ry Rotate an r-matrix about the y-axis.
//
//  - - -
//   R y
//  - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     theta  double          angle (radians)
//
//  Given and returned:
//     r      double[3][3]    r-matrix, rotated
//
//  Notes:
//
//  1) Calling this function with positive theta incorporates in the
//     supplied r-matrix r an additional rotation, about the y-axis,
//     anticlockwise as seen looking towards the origin from positive y.
//
//  2) The additional rotation can be represented by this matrix:
//
//         (  + cos(theta)     0      - sin(theta)  )
//         (                                        )
//         (       0           1           0        )
//         (                                        )
//         (  + sin(theta)     0      + cos(theta)  )
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//void iauRy(double theta, double r[3][3])
func Ry(theta float64, r [3][3]float64) (res [3][3]float64) {
	var s, c, a00, a01, a02, a20, a21, a22 float64

	s = math.Sin(theta)
	c = math.Cos(theta)

	a00 = c*r[0][0] - s*r[2][0]
	a01 = c*r[0][1] - s*r[2][1]
	a02 = c*r[0][2] - s*r[2][2]
	a20 = s*r[0][0] + c*r[2][0]
	a21 = s*r[0][1] + c*r[2][1]
	a22 = s*r[0][2] + c*r[2][2]

	r[0][0] = a00
	r[0][1] = a01
	r[0][2] = a02
	r[2][0] = a20
	r[2][1] = a21
	r[2][2] = a22
	return r
}
