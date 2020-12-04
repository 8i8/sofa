package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoTaitt Time scale transformation:  International Atomic Time, TAI,
//  to Terrestrial Time, TT.
//
//  - - - - - -
//   T a i t t
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  canonical.
//
//  Given:
//     tai1,tai2  double    TAI as a 2-part Julian Date
//
//  Returned:
//     tt1,tt2    double    TT as a 2-part Julian Date
//
//  Returned (function value):
//                int       status:  0 = OK
//
//  Note:
//
//     tai1+tai2 is Julian Date, apportioned in any convenient way
//     between the two arguments, for example where tai1 is the Julian
//     Day Number and tai2 is the fraction of a day.  The returned
//     tt1,tt2 follow suit.
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
//  CgoTaitt Time scale transformation:  International Atomic Time, TAI,
//  to Terrestrial Time, TT.
func CgoTaitt(tai1, tai2 float64) (tt1, tt2 float64, err error) {
	var cTt1, cTt2 C.double
	cI := C.iauTaitt(C.double(tai1), C.double(tai2), &cTt1, &cTt2)
	switch int(cI) {
	case 0:
	default:
		err = errAdmin
	}
	return float64(cTt1), float64(cTt2), err
}

// GoTaitt Time scale transformation:  International Atomic Time, TAI,
// to Terrestrial Time, TT.
func GoTaitt(tai1, tai2 float64) (tt1, tt2 float64, err error) {

	// TT minus TAI (days).
	const dtat = TTMTAI / DAYSEC

	// Result, safeguarding precision.
	if math.Abs(tai1) > math.Abs(tai2) {
		tt1 = tai1
		tt2 = tai2 + dtat
	} else {
		tt1 = tai1 + dtat
		tt2 = tai2
	}
	return
}
