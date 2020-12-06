package sofa

// #include "sofa.h"
import "C"
import (
	"errors"
	"math"
)

var errUtctaiPlus1 = errors.New("dubious year (Note 2)")
var errUtctaiMin1 = errors.New("unacceptable date")

//  CgoUtctai Time scale transformation:  Coordinated Universal Time,
//  UTC, to International Atomic Time, TAI.
//
//  - - - - - - -
//   U t c t a i
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  canonical.
//
//  Given:
//     utc1,utc2  double   UTC as a 2-part quasi Julian Date (Notes 1-4)
//
//  Returned:
//     tai1,tai2  double   TAI as a 2-part Julian Date (Note 5)
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
//     length is 86399, 86400 or 86401 SI seconds.  In the 1960-1972 era
//     there were smaller jumps (in either direction) each time the
//     linear UTC(TAI) expression was changed, and these "mini-leaps"
//     are also included in the SOFA convention.
//
//  3) The warning status "dubious year" flags UTCs that predate the
//     introduction of the time scale or that are too far in the future
//     to be trusted.  See iauDat for further details.
//
//  4) The function iauDtf2d converts from calendar date and time of day
//     into 2-part Julian Date, and in the case of UTC implements the
//     leap-second-ambiguity convention described above.
//
//  5) The returned TAI1,TAI2 are such that their sum is the TAI Julian
//     Date.
//
//  Called:
//     iauJd2cal    JD to Gregorian calendar
//     iauDat       delta(AT) = TAI-UTC
//     iauCal2jd    Gregorian calendar to JD
//
//  References:
//
//     McCarthy, D. D., Petit, G. (eds.), IERS Conventions (2003),
//     IERS Technical Note No. 32, BKG (2004)
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
//  CgoUtctai Time scale transformation:  Coordinated Universal Time,
//  UTC, to International Atomic Time, TAI.
func CgoUtctai(utc1, utc2 float64) (tai1, tai2 float64, err error) {
	var cTai1, cTai2 C.double
	cI := C.iauUtctai(C.double(utc1), C.double(utc2), &cTai1, &cTai2)
	switch int(cI) {
	case 0:
	case -1:
		err = errUtctaiMin1
	case 1:
		err = errUtctaiPlus1
	default:
		err = errAdmin
	}
	return float64(cTai1), float64(cTai2), err
}

// GoUtctai Time scale transformation:  Coordinated Universal Time, UTC,
// to International Atomic Time, TAI.
func GoUtctai(utc1, utc2 float64) (tai1, tai2 float64, err error) {
	var big1 bool
	var iy, im, id, iyt, imt, idt int
	var u1, u2, fd, dat0, dat12, dat24,
		dlod, dleap, z1, z2, a2 float64

	// Put the two parts of the UTC into big-first order.
	if math.Abs(utc1) >= math.Abs(utc2) {
		big1 = true
	}
	if big1 {
		u1 = utc1
		u2 = utc2
	} else {
		u1 = utc2
		u2 = utc1
	}

	// Get TAI-UTC at 0h today.
	iy, im, id, fd, err = GoJd2cal(u1, u2)
	if err != nil {
		return
	}
	dat0, err = GoDat(iy, im, id, 0.0)
	if err != nil && !errors.Is(err, errDat1) {
		return
	}

	// Get TAI-UTC at 12h today (to detect drift).
	dat12, err = GoDat(iy, im, id, 0.5)
	if err != nil && !errors.Is(err, errDat1) {
		return
	}

	// Get TAI-UTC at 0h tomorrow (to detect jumps).
	iyt, imt, idt, _, err = GoJd2cal(u1+1.5, u2-fd)
	if err != nil {
		return
	}
	dat24, err = GoDat(iyt, imt, idt, 0.0)
	if err != nil && !errors.Is(err, errDat1) {
		return
	}

	// Separate TAI-UTC change into per-day (DLOD) and any jump (DLEAP).
	dlod = 2.0 * (dat12 - dat0)
	dleap = dat24 - (dat0 + dlod)

	// Remove any scaling applied to spread leap into preceding day.
	fd *= (DAYSEC + dleap) / DAYSEC

	// Scale from (pre-1972) UTC seconds to SI seconds.
	fd *= (DAYSEC + dlod) / DAYSEC

	// Today's calendar date to 2-part JD.
	z1, z2, err = GoCal2jd(iy, im, id)
	if err != nil {
		return
	}

	// Assemble the TAI result, preserving the UTC split and order.
	a2 = z1 - u1
	a2 += z2
	a2 += fd + dat0/DAYSEC
	if big1 {
		tai1 = u1
		tai2 = a2
	} else {
		tai1 = a2
		tai2 = u1
	}

	return
}
