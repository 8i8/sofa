package sofa

// #include "sofa.h"
import "C"
import (
	"math"

	"github.com/8i8/sofa/en"
)

var errUt1utc = en.New(1, "Ut1utc", []string{
	"unacceptable date",
	"",
	"dubious year (Note 5)",
})

//  CgoUt1utc Time scale transformation:  Universal Time, UT1, to
//  Coordinated Universal Time, UTC.
//
//  - - - - - - -
//   U t 1 u t c
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  canonical.
//
//  Given:
//     ut11,ut12  double   UT1 as a 2-part Julian Date (Note 1)
//     dut1       double   Delta UT1: UT1-UTC in seconds (Note 2)
//
//  Returned:
//     utc1,utc2  double   UTC as a 2-part quasi Julian Date (Notes 3,4)
//
//  Returned (function value):
//                int      status: +1 = dubious year (Note 5)
//                                  0 = OK
//                                 -1 = unacceptable date
//
//  Notes:
//
//  1) ut11+ut12 is Julian Date, apportioned in any convenient way
//     between the two arguments, for example where ut11 is the Julian
//     Day Number and ut12 is the fraction of a day.  The returned utc1
//     and utc2 form an analogous pair, except that a special convention
//     is used, to deal with the problem of leap seconds - see Note 3.
//
//  2) Delta UT1 can be obtained from tabulations provided by the
//     International Earth Rotation and Reference Systems Service.  The
//     value changes abruptly by 1s at a leap second;  however, close to
//     a leap second the algorithm used here is tolerant of the "wrong"
//     choice of value being made.
//
//  3) JD cannot unambiguously represent UTC during a leap second unless
//     special measures are taken.  The convention in the present
//     function is that the returned quasi JD day UTC1+UTC2 represents
//     UTC days whether the length is 86399, 86400 or 86401 SI seconds.
//
//  4) The function iauD2dtf can be used to transform the UTC quasi-JD
//     into calendar date and clock time, including UTC leap second
//     handling.
//
//  5) The warning status "dubious year" flags UTCs that predate the
//     introduction of the time scale or that are too far in the future
//     to be trusted.  See iauDat for further details.
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
//  CgoUt1utc Time scale transformation:  Universal Time, UT1, to
//  Coordinated Universal Time, UTC.
func CgoUt1utc(ut11, ut12, dut1 float64) (
	utc1, utc2 float64, err en.ErrNum) {
	var cUtc1, cUtc2 C.double
	cI := C.iauUt1utc(C.double(ut11), C.double(ut12), C.double(dut1),
		&cUtc1, &cUtc2)
	if n := int(cI); n != 0 {
		err = errUt1utc.Set(n)
	}
	return float64(cUtc1), float64(cUtc2), err
}

//  GoUt1utc Time scale transformation:  Universal Time, UT1, to
//  Coordinated Universal Time, UTC.
func GoUt1utc(ut11, ut12, dut1 float64) (
	utc1, utc2 float64, err en.ErrNum) {
	var big1 bool
	var i, iy, im, id int
	var duts, u1, u2, d1, dats1, d2,
		fd, dats2, ddats, us1, us2, du float64

	// UT1-UTC in seconds.
	duts = dut1

	// Put the two parts of the UT1 into big-first order.
	big1 = (math.Abs(ut11) >= math.Abs(ut12))
	if big1 {
		u1 = ut11
		u2 = ut12
	} else {
		u1 = ut12
		u2 = ut11
	}

	// See if the UT1 can possibly be in a leap-second day.
	d1 = u1
	dats1 = 0
	for i = -1; i <= 3; i++ {
		d2 = u2 + float64(i)
		iy, im, id, fd, err = GoJd2cal(d1, d2)
		if err != nil {
			err = errUt1utc.Set(-1)
			return
		}
		dats2, err = GoDat(iy, im, id, 0.0)
		if err != nil && err.Is() < 0 {
			err = errUt1utc.Set(-1)
			return
		}
		if i == -1 {
			dats1 = dats2
		}
		ddats = dats2 - dats1
		if math.Abs(ddats) >= 0.5 {

			// Yes, leap second nearby: ensure UT1-UTC is
			// "before" value.
			if ddats*duts >= 0 {
				duts -= ddats
			}

			// UT1 for the start of the UTC day that ends in
			// a leap.
			d1, d2, err = GoCal2jd(iy, im, id)
			if err != nil {
				err = errUt1utc.Set(-1)
				return
			}
			us1 = d1
			us2 = d2 - 1.0 + duts/DAYSEC

			// Is the UT1 after this point?
			du = u1 - us1
			du += u2 - us2
			if du > 0 {

				// Yes:  fraction of the current UTC day
				// that has elapsed.
				fd = du * DAYSEC / (DAYSEC + ddats)

				// Ramp UT1-UTC to bring about SOFA's
				// JD(UTC) convention.
				if fd <= 1.0 {
					ddats *= fd
				}
				duts += ddats
			}

			// Done.
			break
		}
		dats1 = dats2
	}

	// Subtract the (possibly adjusted) UT1-UTC from UT1 to give
	// UTC.
	u2 -= duts / DAYSEC

	// Result, safeguarding precision.
	if big1 {
		utc1 = u1
		utc2 = u2
	} else {
		utc1 = u2
		utc2 = u1
	}
	return
}
