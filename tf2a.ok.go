package sofa

// #include "sofa.h"
import "C"
import (
	"math"

	"github.com/8i8/sofa/en"
)

var errTf2a = en.New(0, "Tf2a", []string{
	"",
	"ihour outside range 0-23",
	"imin outside range 0-59",
	"sec outside range 0-59.999...",
})

//  CgoTf2a Convert hours, minutes, seconds to radians.
//
//  - - - - -
//   T f 2 a
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     s         char    sign:  '-' = negative, otherwise positive
//     ihour     int     hours
//     imin      int     minutes
//     sec       double  seconds
//
//  Returned:
//     rad       double  angle in radians
//
//  Returned (function value):
//               int     status:  0 = OK
//                                1 = ihour outside range 0-23
//                                2 = imin outside range 0-59
//                                3 = sec outside range 0-59.999...
//
//  Notes:
//
//  1)  The result is computed even if any of the range checks fail.
//
//  2)  Negative ihour, imin and/or sec produce a warning status, but
//      the absolute value is used in the conversion.
//
//  3)  If there are multiple errors, the status value reflects only the
//      first, the smallest taking precedence.
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoTf2a Convert hours, minutes, seconds to radians.
func CgoTf2a(s byte, ihour, imin int, sec float64) (
	rad float64, err en.ErrNum) {
	var cRad C.double
	cI := C.iauTf2a(C.char(s), C.int(ihour), C.int(imin),
		C.double(sec), &cRad)
	if n := int(cI); n != 0 {
		err = errTf2a.Set(n)
	}
	return float64(cRad), err
}

//  GoTf2a Convert hours, minutes, seconds to radians.
func GoTf2a(s byte, ihour, imin int, sec float64) (
	rad float64, err en.ErrNum) {
	// Compute the interval.
	rad = (60.0*(60.0*(math.Abs(float64(ihour)))+
		(math.Abs(float64(imin)))) +
		math.Abs(sec)) * DS2R
	if s == '-' {
		rad = -rad
	}

	// Validate arguments and return status.
	if ihour < 0 || ihour > 23 {
		err = errTf2a.Set(1)
		return
	}
	if imin < 0 || imin > 59 {
		err = errTf2a.Set(2)
		return
	}
	if sec < 0.0 || sec >= 60.0 {
		err = errTf2a.Set(3)
		return
	}
	return
}
