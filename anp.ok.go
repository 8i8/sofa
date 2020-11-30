package sofa

// #include "sofa.h"
import "C"
import "math"

// Anp Normalize angle into the range 0 <= a < 2pi.
//
//  - - - -
//   A n p
//  - - - -
//
//  Normalize angle into the range 0 <= a < 2pi.
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     a        double     angle (radians)
//
//  Returned (function value):
//              double     angle in range 0-2pi
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
func Anp(a float64) float64 {

	cA := C.iauAnp(C.double(a))
	return float64(cA)
}

// Anp Normalize angle into the range 0 <= a < 2pi.
func goAnp(a float64) float64 {

	a = math.Mod(a, D2PI)
	if a < 0 {
		a += D2PI
	}
	return a
}
