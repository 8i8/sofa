package sofa

// #include "sofa.h"
import "C"
import (
	"math"

	"github.com/8i8/sofa/en"
)

var errTttai = en.New(0, "Tttai", []string{
	"",
})

//  CgoTttai Time scale transformation:  Terrestrial Time, TT, to
//  International Atomic Time, TAI.
//
//  - - - - - -
//   T t t a i
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  canonical.
//
//  Given:
//     tt1,tt2    double    TT as a 2-part Julian Date
//
//  Returned:
//     tai1,tai2  double    TAI as a 2-part Julian Date
//
//  Returned (function value):
//                int       status:  0 = OK
//
//  Note:
//
//     tt1+tt2 is Julian Date, apportioned in any convenient way between
//     the two arguments, for example where tt1 is the Julian Day Number
//     and tt2 is the fraction of a day.  The returned tai1,tai2 follow
//     suit.
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
//  CgoTttai Time scale transformation:  Terrestrial Time, TT, to
//  International Atomic Time, TAI.
func CgoTttai(tt1, tt2 float64) (tai1, tai2 float64, err en.ErrNum) {
	var cTai1, cTai2 C.double
	cI := C.iauTttai(C.double(tt1), C.double(tt2), &cTai1, &cTai2)
	if n := int(cI); n != 0 {
		err = errTttai.Set(n)
	}
	return float64(cTai1), float64(cTai2), err
}

//  GoTttai Time scale transformation:  Terrestrial Time, TT, to
//  International Atomic Time, TAI.
func GoTttai(tt1, tt2 float64) (tai1, tai2 float64, err en.ErrNum) {
	// TT minus TAI (days).
	const dtat = TTMTAI / DAYSEC

	// Result, safeguarding precision.
	if math.Abs(tt1) > math.Abs(tt2) {
		tai1 = tt1
		tai2 = tt2 - dtat
	} else {
		tai1 = tt1 - dtat
		tai2 = tt2
	}

	// Status (always OK).
	return
}
