package sofa

// #include "stdlib.h"
// #include "sofa.h"
import "C"
import (
	"strings"
	"unsafe"

	"github.com/8i8/sofa/en"
)

// var errDtf2dW3 = errors.New("time is after end of day (Note 5) and dubious year (Note 6)")
// var errDtf2dW2 = errors.New("time is after end of day (Note 5)")
// var errDtf2dW1 = errors.New("dubious year (Note 6)")
// var errDtf2dE1 = errors.New("bad year")
// var errDtf2dE2 = errors.New("bad month")
// var errDtf2dE3 = errors.New("bad day")
// var errDtf2dE4 = errors.New("bad hour")
// var errDtf2dE5 = errors.New("bad minute")
// var errDtf2dE6 = errors.New("bad second (<0)")

var errDtf2d = en.New(6, "dtf2d", []string{
	"bad second (<0)",
	"bad minute",
	"bad hour",
	"bad day",
	"bad month",
	"bad year",
	"",
	"dubious year (Note 6)",
	"time is after end of day (Note 5)",
	"time is after end of day (Note 5) and dubious year (Note 6)",
})

//  CgoDtf2d Encode date and time fields into 2-part Julian Date (or in
//  the case of UTC a quasi-JD form that includes special provision for
//  leap seconds).
//
//  - - - - - -
//   D t f 2 d
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     scale     char[]  time scale ID (Note 1)
//     iy,im,id  int     year, month, day in Gregorian calendar (Note 2)
//     ihr,imn   int     hour, minute
//     sec       double  seconds
//
//  Returned:
//     d1,d2     double  2-part Julian Date (Notes 3,4)
//
//  Returned (function value):
//               int     status: +3 = both of next two
//                               +2 = time is after end of day (Note 5)
//                               +1 = dubious year (Note 6)
//                                0 = OK
//                               -1 = bad year
//                               -2 = bad month
//                               -3 = bad day
//                               -4 = bad hour
//                               -5 = bad minute
//                               -6 = bad second (<0)
//
//  Notes:
//
//  1) scale identifies the time scale.  Only the value "UTC" (in upper
//     case) is significant, and enables handling of leap seconds (see
//     Note 4).
//
//  2) For calendar conventions and limitations, see iauCal2jd.
//
//  3) The sum of the results, d1+d2, is Julian Date, where normally d1
//     is the Julian Day Number and d2 is the fraction of a day.  In the
//     case of UTC, where the use of JD is problematical, special
//     conventions apply:  see the next note.
//
//  4) JD cannot unambiguously represent UTC during a leap second unless
//     special measures are taken.  The SOFA internal convention is that
//     the quasi-JD day represents UTC days whether the length is 86399,
//     86400 or 86401 SI seconds.  In the 1960-1972 era there were
//     smaller jumps (in either direction) each time the linear UTC(TAI)
//     expression was changed, and these "mini-leaps" are also included
//     in the SOFA convention.
//
//  5) The warning status "time is after end of day" usually means that
//     the sec argument is greater than 60.0.  However, in a day ending
//     in a leap second the limit changes to 61.0 (or 59.0 in the case
//     of a negative leap second).
//
//  6) The warning status "dubious year" flags UTCs that predate the
//     introduction of the time scale or that are too far in the future
//     to be trusted.  See iauDat for further details.
//
//  7) Only in the case of continuous and regular time scales (TAI, TT,
//     TCG, TCB and TDB) is the result d1+d2 a Julian Date, strictly
//     speaking.  In the other cases (UT1 and UTC) the result must be
//     used with circumspection;  in particular the difference between
//     two such results cannot be interpreted as a precise time
//     interval.
//
//  Called:
//     iauCal2jd    Gregorian calendar to JD
//     iauDat       delta(AT) = TAI-UTC
//     iauJd2cal    JD to Gregorian calendar
//
//  This revision:  2013 July 26
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoDtf2d Encode date and time fields into 2-part Julian Date (or in
//  the case of UTC a quasi-JD form that includes special provision for
//  leap seconds).
func CgoDtf2d(scale string, iy, im, id, ihr, imn int, sec float64) (
	d1, d2 float64, err en.ErrNum) {
	var cD1, cD2 C.double
	cS := C.CString(scale)
	cI := C.iauDtf2d(cS, C.int(iy), C.int(im), C.int(id), C.int(ihr),
		C.int(imn), C.double(sec), &cD1, &cD2)
	C.free(unsafe.Pointer(cS))
	switch int(cI) {
	case 0:
	default:
		err = errDtf2d.Set(int(cI))
	}
	return float64(cD1), float64(cD2), err
}

//  GoDtf2d Encode date and time fields into 2-part Julian Date (or in
//  the case of UTC a quasi-JD form that includes special provision for
//  leap seconds).
func GoDtf2d(scale string, iy, im, id, ihr, imn int, sec float64) (
	d1, d2 float64, err en.ErrNum) {

	var iy2, im2, id2 int
	var dj, w, day, seclim, dat0, dat12, dat24, dleap, time float64

	// Today's Julian Day Number.
	dj, w, err = GoCal2jd(iy, im, id)
	if err != nil {
		err = errDtf2d.Set(1)
		return
	}
	dj += w

	// Day length and final minute length in seconds (provisional).
	day = DAYSEC
	seclim = 60.0

	// Deal with the UTC leap second case.
	if strings.Compare(scale, "UTC") == 0 {

		// TAI-UTC at 0h today.
		dat0, err = GoDat(iy, im, id, 0.0)
		if err != nil {
			if err.Is() < 0 {
				err = errDtf2d.Wrap(err)
				return
			}
			err = errDtf2d.Wrap(err)
		}

		// TAI-UTC at 12h today (to detect drift).
		dat12, err = GoDat(iy, im, id, 0.5)
		if err != nil {
			if err.Is() < 0 {
				err = errDtf2d.Wrap(err)
				return
			}
			err = errDtf2d.Wrap(err)
		}

		// TAI-UTC at 0h tomorrow (to detect jumps).
		iy2, im2, id2, w, err = GoJd2cal(dj, 1.5)
		if err != nil {
			err = errDtf2d.Wrap(err)
			return
		}

		dat24, err = GoDat(iy2, im2, id2, 0.0)
		if err != nil {
			//if !errors.Is(err, errDatWarn) {
			if err.Is() < 0 {
				err = errDtf2d.Wrap(err)
				return
			}
			err = errDtf2d.Wrap(err)
		}

		// Any sudden change in TAI-UTC between today and tomorrow.
		dleap = dat24 - (2.0*dat12 - dat0)

		// If leap second day, correct the day and final minute lengths.
		day += dleap
		if ihr == 23 && imn == 59 {
			seclim += dleap
		}

		// End of UTC-specific actions.
	}

	// Validate the time.
	if ihr >= 0 && ihr <= 23 {
		if imn >= 0 && imn <= 59 {
			if sec >= 0 {
				if sec >= seclim {
					err = errDtf2d.Add(err, 2)
				}
			} else {
				err = err.Set(-6)
			}
		} else {
			err = err.Set(-5)
		}
	} else {
		err = err.Set(-4)
	}
	if err != nil && err.Is() < 0 {
		return
	}

	// The time in days.
	time = (60.0*(float64(60*ihr+imn)) + sec) / day

	// Return the date and time.
	d1 = dj
	d2 = time

	// Status.
	return
}
