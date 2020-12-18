package sofa

// #include "sofa.h"
import "C"
import (
	"math"

	"github.com/8i8/sofa/en"
)

var errJdcalf = en.New(1, "Jdcalf", []string{
	"date out of range",
	"OK",
	"NDP not 0-9 (interpreted as 0)",
})

//  CgoJdcalf Julian Date to Gregorian Calendar, expressed in a form
//  convenient for formatting messages:  rounded to a specified
//  precision.
//
//  - - - - - - -
//   J d c a l f
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     ndp       int      number of decimal places of days in fraction
//     dj1,dj2   double   dj1+dj2 = Julian Date (Note 1)
//
//  Returned:
//     iymdf     int[4]   year, month, day, fraction in Gregorian
//                        calendar
//
//  Returned (function value):
//               int      status:
//                          -1 = date out of range
//                           0 = OK
//                          +1 = NDP not 0-9 (interpreted as 0)
//
//  Notes:
//
//  1) The Julian Date is apportioned in any convenient way between
//     the arguments dj1 and dj2.  For example, JD=2450123.7 could
//     be expressed in any of these ways, among others:
//
//             dj1            dj2
//
//         2450123.7           0.0       (JD method)
//         2451545.0       -1421.3       (J2000 method)
//         2400000.5       50123.2       (MJD method)
//         2450123.5           0.2       (date & time method)
//
//  2) In early eras the conversion is from the "Proleptic Gregorian
//     Calendar";  no account is taken of the date(s) of adoption of
//     the Gregorian Calendar, nor is the AD/BC numbering convention
//     observed.
//
//  3) Refer to the function iauJd2cal.
//
//  4) NDP should be 4 or less if internal overflows are to be
//     avoided on machines which use 16-bit integers.
//
//  Called:
//     iauJd2cal    JD to Gregorian calendar
//
//  Reference:
//
//     Explanatory Supplement to the Astronomical Almanac,
//     P. Kenneth Seidelmann (ed), University Science Books (1992),
//     Section 12.92 (p604).
//
//  This revision:  2020 April 13
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoJdcalf Julian Date to Gregorian Calendar, expressed in a form
//  convenient for formatting messages:  rounded to a specified
//  precision.
func CgoJdcalf(ndp int, dj1, dj2 float64) (
	iymdf [4]int, err en.ErrNum) {
	var cIymdf [4]C.int
	cI := C.iauJdcalf(C.int(ndp), C.double(dj1), C.double(dj2),
		&cIymdf[0])
	if n := int(cI); n != 0 {
		err = errJdcalf.Set(n)
	}
	return v4sIntC2Go(cIymdf), err
}

//  GoJdcalf Julian Date to Gregorian Calendar, expressed in a form
//  convenient for formatting messages:  rounded to a specified
//  precision.
func GoJdcalf(ndp int, dj1, dj2 float64) (
	iymdf [4]int, err en.ErrNum) {
	var err1 en.ErrNum
	var denom, d1, d2, f1, f2, d, djd, f, rf float64

	// Denominator of fraction (e.g. 100 for 2 decimal places).
	if (ndp >= 0) && (ndp <= 9) {
		denom = float64(pow(10.0, ndp))
	} else {
		err = errJdcalf.Set(1)
		denom = 1.0
	}

	// Copy the date, big then small.
	if math.Abs(dj1) >= math.Abs(dj2) {
		d1 = dj1
		d2 = dj2
	} else {
		d1 = dj2
		d2 = dj1
	}

	// Realign to midnight (without rounding error).
	d1 -= 0.5

	// Separate day and fraction (as precisely as possible).
	d = dnint(d1)
	f1 = d1 - d
	djd = d
	d = dnint(d2)
	f2 = d2 - d
	djd += d
	d = dnint(f1 + f2)
	f = (f1 - d) + f2
	if f < 0.0 {
		f += 1.0
		d -= 1.0
	}
	djd += d

	// Round the total fraction to the specified number of places.
	rf = dnint(f*denom) / denom

	// Re-align to noon.
	djd += 0.5

	// Convert to Gregorian calendar.
	iymdf[0], iymdf[1], iymdf[2], f, err1 = GoJd2cal(djd, rf)
	if err1 == nil {
		iymdf[3] = int(dnint(f * denom))
	} else {
		err = err1
	}

	return
}
