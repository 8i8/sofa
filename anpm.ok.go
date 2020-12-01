package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoAnpm Normalize angle into the range -pi <= a < +pi.
//
//  - - - - -
//   A n p m
//  - - - - -
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
//              double     angle in range +/-pi
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoAnpm Normalize angle into the range -pi <= a < +pi.
func CgoAnpm(a float64) float64 {
	cA := C.iauAnpm(C.double(a))
	return float64(cA)
}

// GoAnpm Normalize angle into the range -pi <= a < +pi.
func GoAnpm(a float64) float64 {
	var w float64

	w = math.Mod(a, D2PI)
	if math.Abs(w) >= DPI {
		w -= dsign(D2PI, a)
	}

	return w
}
