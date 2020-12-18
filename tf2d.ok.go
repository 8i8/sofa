package sofa

// #include "sofa.h"
import "C"
import (
	"math"

	"github.com/8i8/sofa/en"
)

var errTf2d = en.New(0, "Tf2d", []string{
	"",
	"ihour outside range 0-23",
	"imin outside range 0-59",
	"sec outside range 0-59.999...",
})

//  CgoTf2d Convert hours, minutes, seconds to days.
//
//  - - - - -
//   T f 2 d
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
//     days      double  interval in days
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
//  CgoTf2d Convert hours, minutes, seconds to days.
func CgoTf2d(s byte, ihour, imin int, sec float64) (
	days float64, err en.ErrNum) {
	var cDays C.double
	cI := C.iauTf2d(C.char(s), C.int(ihour), C.int(imin),
		C.double(sec), &cDays)
	if int(cI) != 0 {
		err = errTf2d.Set(int(cI))
	}
	return float64(cDays), err
}

//  GoTf2d Convert hours, minutes, seconds to days.
func GoTf2d(s byte, ihour, imin int, sec float64) (
	days float64, err en.ErrNum) {

	// Compute the interval.
	days = (60.0*(60.0*(math.Abs(float64(ihour)))+
		(math.Abs(float64(imin)))) +
		math.Abs(sec)) / DAYSEC
	if s == '-' {
		days = -days
	}

	// Validate arguments and return status.
	if ihour < 0 || ihour > 23 {
		err = errTf2d.Set(1)
		return
	}
	if imin < 0 || imin > 59 {
		err = errTf2d.Set(2)
		return
	}
	if sec < 0.0 || sec >= 60.0 {
		err = errTf2d.Set(3)
		return
	}
	return
}
