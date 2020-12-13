package sofa

// #include "stdlib.h"
// #include "sofa.h"
import "C"
import (
	"errors"
	"strings"
	"unsafe"
)

var errD2dtfWarn = errors.New("dubious year (Note 5)")
var errD2dtfE1 = errors.New("unacceptable date (Note 4)")

//  CgoD2dtf Format for output a 2-part Julian Date (or in the case of
//  UTC a quasi-JD form that includes special provision for leap
//  seconds).
//
//  - - - - - -
//   D 2 d t f
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     scale     char[]  time scale ID (Note 1)
//     ndp       int     resolution (Note 2)
//     d1,d2     double  time as a 2-part Julian Date (Notes 3,4)
//
//  Returned:
//     iy,im,id  int     year, month, day in Gregorian calendar (Note 5)
//     ihmsf     int[4]  hours, minutes, seconds, fraction (Note 1)
//
//  Returned (function value):
//               int     status: +1 = dubious year (Note 5)
//                                0 = OK
//                               -1 = unacceptable date (Note 6)
//
//  Notes:
//
//  1) scale identifies the time scale.  Only the value "UTC" (in upper
//     case) is significant, and enables handling of leap seconds (see
//     Note 4).
//
//  2) ndp is the number of decimal places in the seconds field, and can
//     have negative as well as positive values, such as:
//
//     ndp         resolution
//     -4            1 00 00
//     -3            0 10 00
//     -2            0 01 00
//     -1            0 00 10
//      0            0 00 01
//      1            0 00 00.1
//      2            0 00 00.01
//      3            0 00 00.001
//
//     The limits are platform dependent, but a safe range is -5 to +9.
//
//  3) d1+d2 is Julian Date, apportioned in any convenient way between
//     the two arguments, for example where d1 is the Julian Day Number
//     and d2 is the fraction of a day.  In the case of UTC, where the
//     use of JD is problematical, special conventions apply:  see the
//     next note.
//
//  4) JD cannot unambiguously represent UTC during a leap second unless
//     special measures are taken.  The SOFA internal convention is that
//     the quasi-JD day represents UTC days whether the length is 86399,
//     86400 or 86401 SI seconds.  In the 1960-1972 era there were
//     smaller jumps (in either direction) each time the linear UTC(TAI)
//     expression was changed, and these "mini-leaps" are also included
//     in the SOFA convention.
//
//  5) The warning status "dubious year" flags UTCs that predate the
//     introduction of the time scale or that are too far in the future
//     to be trusted.  See iauDat for further details.
//
//  6) For calendar conventions and limitations, see iauCal2jd.
//
//  Called:
//     iauJd2cal    JD to Gregorian calendar
//     iauD2tf      decompose days to hms
//     iauDat       delta(AT) = TAI-UTC
//
//  This revision:  2014 February 15
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoD2dtf Format for output a 2-part Julian Date (or in the case of
//  UTC a quasi-JD form that includes special provision for leap
//  seconds).
func CgoD2dtf(scale string, ndp int, d1, d2 float64) (
	iy, im, id int, ihmsf [4]int, err error) {
	var cIhmsf [4]C.int
	cS := C.CString(scale)
	cIy, cIm, cId := C.int(iy), C.int(im), C.int(id)
	cI := C.iauD2dtf(cS, C.int(ndp), C.double(d1),
		C.double(d2), &cIy, &cIm, &cId, &cIhmsf[0])
	C.free(unsafe.Pointer(cS))
	switch int(cI) {
	case 0:
	case 1:
		err = errD2dtfWarn
	case -1:
		err = errD2dtfWarn
	default:
		err = errAdmin
	}
	return int(cIy), int(cIm), int(cId), v4sIntC2Go(cIhmsf), err
}

//  GoD2dtf Format for output a 2-part Julian Date (or in the case of
//  UTC a quasi-JD form that includes special provision for leap
//  seconds).
func GoD2dtf(scale string, ndp int, d1, d2 float64) (
	iy, im, id int, ihmsf [4]int, err error) {
	var leap bool
	var iy1, im1, id1, iy2, im2, id2, i int
	var ihmsf1 [4]int
	var a1, b1, fd, dat0, dat12, dat24, dleap float64

	// The two-part JD.
	a1 = d1
	b1 = d2

	// Provisional calendar date.
	iy1, im1, id1, fd, err = GoJd2cal(a1, b1)
	if err != nil {
		err = errD2dtfE1
		return
	}

	// Is this a leap second day?
	// TODO check this
	if strings.Compare(scale, "UTC") == 0 {

		// TAI-UTC at 0h today.
		dat0, err = GoDat(iy1, im1, id1, 0.0)
		if err != nil {
			if !errors.Is(err, errD2dtfWarn) {
				err = errD2dtfE1
				return
			}
			err = errD2dtfWarn
		}

		// TAI-UTC at 12h today (to detect drift).
		dat12, err = GoDat(iy1, im1, id1, 0.5)
		if err != nil {
			if !errors.Is(err, errD2dtfWarn) {
				err = errD2dtfE1
				return
			}
			err = errD2dtfWarn
		}

		// TAI-UTC at 0h tomorrow (to detect jumps).
		iy2, im2, id2, _, err = GoJd2cal(a1+1.5, b1-fd)
		if err != nil {
			err = errD2dtfE1
			return
		}
		dat24, err = GoDat(iy2, im2, id2, 0.0)
		if err != nil {
			if !errors.Is(err, errD2dtfWarn) {
				err = errD2dtfE1
				return
			}
			err = errD2dtfWarn
		}

		// Any sudden change in TAI-UTC (seconds).
		dleap = dat24 - (2.0*dat12 - dat0)

		// If leap second day, scale the fraction of a day into
		// SI.
		leap = (dleap != 0.0)
		if leap {
			fd += fd * dleap / DAYSEC
		}
	}

	// Provisional time of day.
	_, ihmsf1 = GoD2tf(ndp, fd)

	// Has the (rounded) time gone past 24h?
	if ihmsf1[0] > 23 {

		// Yes.  We probably need tomorrow's calendar date.
		iy2, im2, id2, _, err = GoJd2cal(a1+1.5, b1-fd)
		if err != nil {
			if !errors.Is(err, errD2dtfWarn) {
				err = errD2dtfE1
				return
			}
			err = errD2dtfWarn
		}

		// Is today a leap second day?
		if !leap {

			// No.  Use 0h tomorrow.
			iy1 = iy2
			im1 = im2
			id1 = id2
			ihmsf1[0] = 0
			ihmsf1[1] = 0
			ihmsf1[2] = 0

		} else {

			// Yes.  Are we past the leap second itself?
			if ihmsf1[2] > 0 {

				// Yes.  Use tomorrow but allow for the
				// leap second.
				iy1 = iy2
				im1 = im2
				id1 = id2
				ihmsf1[0] = 0
				ihmsf1[1] = 0
				ihmsf1[2] = 0

			} else {

				// No.  Use 23 59 60... today.
				ihmsf1[0] = 23
				ihmsf1[1] = 59
				ihmsf1[2] = 60
			}

			// If rounding to 10s or coarser always go up to
			// new day.
			if ndp < 0 && ihmsf1[2] == 60 {
				iy1 = iy2
				im1 = im2
				id1 = id2
				ihmsf1[0] = 0
				ihmsf1[1] = 0
				ihmsf1[2] = 0
			}
		}
	}

	// Results.
	iy = iy1
	im = im1
	id = id1
	for i = 0; i < 4; i++ {
		ihmsf[i] = ihmsf1[i]
	}

	// Status.
	return
}
