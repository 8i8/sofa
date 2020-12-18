package sofa

// #include "sofa.h"
import "C"
import (
	"errors"
	"math"
	"github.com/8i8/sofa/en"
)

var errJd2calE1 = errors.New("unacceptable date (Note 1)")

var errJd2cal = en.New(1, "jd2cal", []string{"unacceptable date (Note 1)"})

//  CgoJd2cal Julian Date to Gregorian year, month, day, and fraction of
//  a day.
//
//  - - - - - - -
//   J d 2 c a l
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     dj1,dj2   double   Julian Date (Notes 1, 2)
//
//  Returned (arguments):
//     iy        int      year
//     im        int      month
//     id        int      day
//     fd        double   fraction of day
//
//  Returned (function value):
//               int      status:
//                           0 = OK
//                          -1 = unacceptable date (Note 1)
//
//  Notes:
//
//  1) The earliest valid date is -68569.5 (-4900 March 1).  The
//     largest value accepted is 1e9.
//
//  2) The Julian Date is apportioned in any convenient way between
//     the arguments dj1 and dj2.  For example, JD=2450123.7 could
//     be expressed in any of these ways, among others:
//
//            dj1             dj2
//
//         2450123.7           0.0       (JD method)
//         2451545.0       -1421.3       (J2000 method)
//         2400000.5       50123.2       (MJD method)
//         2450123.5           0.2       (date & time method)
//
//     Separating integer and fraction uses the "compensated summation"
//     algorithm of Kahan-Neumaier to preserve as much precision as
//     possible irrespective of the jd1+jd2 apportionment.
//
//  3) In early eras the conversion is from the "proleptic Gregorian
//     calendar";  no account is taken of the date(s) of adoption of
//     the Gregorian calendar, nor is the AD/BC numbering convention
//     observed.
//
//  References:
//
//     Explanatory Supplement to the Astronomical Almanac,
//     P. Kenneth Seidelmann (ed), University Science Books (1992),
//     Section 12.92 (p604).
//
//     Klein, A., A Generalized Kahan-Babuska-Summation-Algorithm.
//     Computing 76, 279-293 (2006), Section 3.
//
//  This revision:  2020 June 24
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoJd2cal Julian Date to Gregorian year, month, day, and fraction of
//  a day.
func CgoJd2cal(dj1, dj2 float64) (
	iy, im, id int, fd float64, err en.ErrNum) {

	var cIy, cIm, cId C.int
	var cFd C.double
	cI := C.iauJd2cal(C.double(dj1), C.double(dj2), 
	&cIy, &cIm, &cId, &cFd)
	if n := int(cI); n != 0 {
		err = errJd2cal.Set(n)
	}
	return int(cIy), int(cIm), int(cId), float64(cFd), err
}

// GoJd2cal Julian Date to Gregorian year, month, day, and fraction of a
// day.
func GoJd2cal(dj1, dj2 float64) (
	iy, im, id int, fd float64, err en.ErrNum) {

	// Minimum and maximum allowed JD
	const DJMIN = -68569.5
	const DJMAX = 1e9

	var jd, i, l, n, k int
	var dj, f1, f2, d, s, cs, x, t, f float64
	var v [2]float64

	// Verify date is acceptable.
	dj = dj1 + dj2
	if dj < DJMIN || dj > DJMAX {
		err = errJd2cal.Set(-1)
		return
	}

	// Separate day and fraction (where -0.5 <= fraction < 0.5).
	//d = dnint(dj1);
	d = math.Round(dj1)
	f1 = dj1 - d
	jd = int(d)
	//d = dnint(dj2);
	d = math.Round(dj2)
	f2 = dj2 - d
	jd += int(d)

	// Compute f1+f2+0.5 using compensated summation (Klein 2006).
	s = 0.5
	cs = 0.0
	v[0] = f1
	v[1] = f2
	for i = 0; i < 2; i++ {
		x = v[i]
		t = s + x
		if math.Abs(s) >= math.Abs(x) {
			cs += (s - t) + x
		} else {
			cs += (x - t) + s
		}

		s = t
		if s >= 1.0 {
			jd++
			s -= 1.0
		}
	}
	f = s + cs
	cs = f - s

	// Deal with negative f.
	if f < 0.0 {

		// Compensated summation: assume that |s| <= 1.0.
		f = s + 1.0
		cs += (1.0 - f) + s
		s = f
		f = s + cs
		cs = f - s
		jd--
	}

	// Deal with f that is 1.0 or more (when rounded to double).
	if (f - 1.0) >= -DBL_EPSILON/4.0 {

		// Compensated summation: assume that |s| <= 1.0.
		t = s - 1.0
		cs += (s - t) - 1.0
		s = t
		f = s + cs
		if -DBL_EPSILON/2.0 < f {
			jd++
			f = fmax(f, 0.0)
		}
	}

	// Express day in Gregorian calendar.
	l = jd + 68569
	n = (4 * l) / 146097
	l -= (146097*n + 3) / 4
	i = (4000 * (l + 1)) / 1461001
	l -= (1461*i)/4 - 31
	k = (80 * l) / 2447
	id = (l - (2447*k)/80)
	l = k / 11
	im = (k + 2 - 12*l)
	iy = (100*(n-49) + i + l)
	fd = f

	return
}
