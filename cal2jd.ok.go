package sofa

// #include <sofa.h>
import "C"
import (
	"errors"

	"github.com/8i8/sofa/en"
)

var errCal2jdE1 = errors.New("bad year (cal2jd documentation note 3: JD not computed)")
var errCal2jdE2 = errors.New("bad month (JD not computed)")
var errCal2jdE3 = errors.New("bad day (JD computed)")

var errCal2jd = en.New(3, "cal2jd", []string{
	"bad year (cal2jd documentation note 3: JD not computed)",
	"bad month (JD not computed)",
	"bad day (JD computed)",
})

//  CgoCal2jd returns the MJD, Modified Julian Date of the given date.
//
//  - - - - - - -
//   C a l 2 j d
//  - - - - - - -
//
//  Gregorian Calendar to Julian Date.
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     iy,im,id  int     year, month, day in Gregorian calendar (Note 1)
//
//  Returned:
//     djm0      double  MJD zero-point: always 2400000.5
//     djm       double  Modified Julian Date for 0 hrs
//
//  Returned (function value):
//     err       error   nil = OK
//                       errCal2jdE1 = bad year   (Note 3: JD not computed)
//                       errCal2jdE2 = bad month  (JD not computed)
//                       errCal2jdE3 = bad day    (JD computed)
//
//  Notes:
//
//  1) The algorithm used is valid from -4800 March 1, but this
//     implementation rejects dates before -4799 January 1.
//
//  2) The Julian Date is returned in two pieces, in the usual SOFA
//     manner, which is designed to preserve time resolution.  The
//     Julian Date is available as a single number by adding djm0 and
//     djm.
//
//  3) In early eras the conversion is from the "Proleptic Gregorian
//     Calendar";  no account is taken of the date(s) of adoption of
//     the Gregorian Calendar, nor is the AD/BC numbering convention
//     observed.
//
//  Reference:
//
//     Explanatory Supplement to the Astronomical Almanac,
//     P. Kenneth Seidelmann (ed), University Science Books (1992),
//     Section 12.92 (p604).
//
//  This revision:  2013 August 7
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoCal2jd returns the MJD, Modified Julian Date of the given date.
func CgoCal2jd(iy, im, id int) (djm0, djm float64, err en.ErrNum) {
	var cDjm0, cDjm C.double
	cI := C.iauCal2jd(C.int(iy), C.int(im), C.int(id), &cDjm0, &cDjm)
	if n := int(cI); n != 0 {
		err = errCal2jd.Set(n)
	}
	return float64(cDjm0), float64(cDjm), err
}

// GoCal2jd returns the MJD, Modified Julian Date of the given date.
func GoCal2jd(iy, im, id int) (djm0, djm float64, err en.ErrNum) {
	var ly, my int
	var iypmy int // was a long in c code

	// Earliest year allowed (4800BC)
	const IYMIN = -4799

	// Month lengths in days
	var mtab = [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// Validate year and month.
	if iy < IYMIN {
		err = errCal2jd.Set(-1)
		return
	}
	if im < 1 || im > 12 {
		err = errCal2jd.Set(-2)
		return
	}

	// If February in a leap year, 1, otherwise 0.
	if (im == 2) && ((iy % 4) == 0) && ((iy%100 != 0) || ((iy % 400) != 0)) {
		ly = 1
	}

	// Validate day, taking into account leap years.
	if (id < 1) || (id > (mtab[im-1] + ly)) {
		err = errCal2jd.Set(-3)
		return
	}

	// Return result.
	my = (im - 14) / 12
	iypmy = iy + my
	djm0 = DJM0
	djm = float64(((1461*(iypmy+4800))/4 +
		(367*(im-2-12*my))/12 -
		(3*((iypmy+4900)/100))/4 +
		id - 2432076))
	return
}
