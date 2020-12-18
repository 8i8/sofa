package sofa

// #include "sofa.h"
import "C"
import (
	"math"

	"github.com/8i8/sofa/en"
)

var errTaiut1 = en.New(0, "Taiut1", []string{
	"",
})

//  CgoTaiut1 Time scale transformation:  International Atomic Time,
//  TAI, to Universal Time, UT1.
//
//  - - - - - - -
//   T a i u t 1
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  canonical.
//
//  Given:
//     tai1,tai2  double    TAI as a 2-part Julian Date
//     dta        double    UT1-TAI in seconds
//
//  Returned:
//     ut11,ut12  double    UT1 as a 2-part Julian Date
//
//  Returned (function value):
//                int       status:  0 = OK
//
//  Notes:
//
//  1) tai1+tai2 is Julian Date, apportioned in any convenient way
//     between the two arguments, for example where tai1 is the Julian
//     Day Number and tai2 is the fraction of a day.  The returned
//     UT11,UT12 follow suit.
//
//  2) The argument dta, i.e. UT1-TAI, is an observed quantity, and is
//     available from IERS tabulations.
//
//  Reference:
//
//     Explanatory Supplement to the Astronomical Almanac,
//     P. Kenneth Seidelmann (ed), University Science Books (1992)
//
//  This revision:  2019 June 20
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//
//  CgoTaiut1 Time scale transformation:  International Atomic Time,
//  TAI, to Universal Time, UT1.
func CgoTaiut1(tai1, tai2, dta float64) (ut11, ut12 float64, err en.ErrNum) {
	var cUt11, cUt12 C.double
	cI := C.iauTaiut1(C.double(tai1), C.double(tai2), C.double(dta),
		&cUt11, &cUt12)
	switch int(cI) {
	case 0:
	default:
		err = errTaiut1.Set(0)
	}
	return float64(cUt11), float64(cUt12), err
}

//  GoTaiut1 Time scale transformation:  International Atomic Time,
//  TAI, to Universal Time, UT1.
func GoTaiut1(tai1, tai2, dta float64) (ut11, ut12 float64, err en.ErrNum) {
	var dtad float64

	// Result, safeguarding precision.
	dtad = dta / DAYSEC
	if math.Abs(tai1) > math.Abs(tai2) {
		ut11 = tai1
		ut12 = tai2 + dtad
	} else {
		ut11 = tai1 + dtad
		ut12 = tai2
	}

	// Status (always OK).
	return
}
