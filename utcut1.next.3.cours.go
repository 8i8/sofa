package sofa

// #include "sofa.h"
import "C"
import "errors"

var errUtcut1Min1 = errors.New("unacceptable date")
var errUtcut1Plus1 = errors.New("dubious year (Note 3)")

//  CgoUtcut1 Time scale transformation:  Coordinated Universal Time,
//  UTC, to Universal Time, UT1.
//
//  - - - - - - -
//   U t c u t 1
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  canonical.
//
//  Given:
//     utc1,utc2  double   UTC as a 2-part quasi Julian Date (Notes 1-4)
//     dut1       double   Delta UT1 = UT1-UTC in seconds (Note 5)
//
//  Returned:
//     ut11,ut12  double   UT1 as a 2-part Julian Date (Note 6)
//
//  Returned (function value):
//                int      status: +1 = dubious year (Note 3)
//                                  0 = OK
//                                 -1 = unacceptable date
//
//  Notes:
//
//  1) utc1+utc2 is quasi Julian Date (see Note 2), apportioned in any
//     convenient way between the two arguments, for example where utc1
//     is the Julian Day Number and utc2 is the fraction of a day.
//
//  2) JD cannot unambiguously represent UTC during a leap second unless
//     special measures are taken.  The convention in the present
//     function is that the JD day represents UTC days whether the
//     length is 86399, 86400 or 86401 SI seconds.
//
//  3) The warning status "dubious year" flags UTCs that predate the
//     introduction of the time scale or that are too far in the future
//     to be trusted.  See iauDat for further details.
//
//  4) The function iauDtf2d converts from calendar date and time of
//     day into 2-part Julian Date, and in the case of UTC implements
//     the leap-second-ambiguity convention described above.
//
//  5) Delta UT1 can be obtained from tabulations provided by the
//     International Earth Rotation and Reference Systems Service.
//     It is the caller's responsibility to supply a dut1 argument
//     containing the UT1-UTC value that matches the given UTC.
//
//  6) The returned ut11,ut12 are such that their sum is the UT1 Julian
//     Date.
//
//  References:
//
//     McCarthy, D. D., Petit, G. (eds.), IERS Conventions (2003),
//     IERS Technical Note No. 32, BKG (2004)
//
//     Explanatory Supplement to the Astronomical Almanac,
//     P. Kenneth Seidelmann (ed), University Science Books (1992)
//
//  Called:
//     iauJd2cal    JD to Gregorian calendar
//     iauDat       delta(AT) = TAI-UTC
//     iauUtctai    UTC to TAI
//     iauTaiut1    TAI to UT1
//
//  This revision:  2013 August 12
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoUtcut1 Time" scale transformation:  Coordinated Universal Time,
//  UTC, to Universal Time, UT1.
func CgoUtcut1(utc1, utc2, dut1 float64) (ut11, ut12 float64, err error) {
	var cUt11, cUt12 C.double
	cI := C.iauUtcut1(C.double(utc1), C.double(utc2), C.double(dut1),
		&cUt11, &cUt12)
	switch int(cI) {
	case 0:
	case -1:
		err = errUtcut1Min1
	case 1:
		err = errUtcut1Plus1
	default:
		err = errAdmin
	}
	return float64(cUt11), float64(cUt12), err
}

// GoUtcut1 Time scale transformation:  Coordinated Universal Time,
// UTC, to Universal Time, UT1.
func GoUtcut1(utc1, utc2, dut1 float64) (ut11, ut12 float64, err error) {

	var iy, im, id int
	var dat, dta, tai1, tai2 float64
	var errs, errw error

	// Look up TAI-UTC.
	iy, im, id, _, err = GoJd2cal(utc1, utc2)
	if err != nil {
		err = errUtcut1Min1
		return
	}
	dat, errs = GoDat(iy, im, id, 0.0)
	if errs != nil {
		if !errors.Is(errs, errDat1) {
			err = errUtcut1Min1
			return
		}
		err = errUtcut1Plus1
	}

	// Form UT1-TAI.
	dta = dut1 - dat

	// UTC to TAI to UT1.
	tai1, tai2, errw = GoUtctai(utc1, utc2)
	if errw != nil {
		if !errors.Is(errw, errUtctaiPlus1) {
			err = errUtcut1Min1
			return
		}
		err = errUtcut1Plus1
	}
	ut11, ut12, err = GoTaiut1(tai1, tai2, dta)
	if err != nil {
		err = errUtcut1Min1
		return
	}

	return
}
