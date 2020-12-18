package sofa

// #include "sofa.h"
import "C"
import (
	"math"

	"github.com/8i8/sofa/en"
)

var errTttcg = en.New(0, "Tttcg", []string{
	"",
})

//  CgoTttcg Time scale transformation:  Terrestrial Time, TT, to
//  Geocentric Coordinate Time, TCG.
//
//  - - - - - -
//   T t t c g
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
//     tcg1,tcg2  double    TCG as a 2-part Julian Date
//
//  Returned (function value):
//                int       status:  0 = OK
//
//  Note:
//
//     tt1+tt2 is Julian Date, apportioned in any convenient way between
//     the two arguments, for example where tt1 is the Julian Day Number
//     and tt2 is the fraction of a day.  The returned tcg1,tcg2 follow
//     suit.
//
//  References:
//
//     McCarthy, D. D., Petit, G. (eds.), IERS Conventions (2003),
//     IERS Technical Note No. 32, BKG (2004)
//
//     IAU 2000 Resolution B1.9
//
//  This revision:  2019 June 20
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoTttcg Time scale transformation:  Terrestrial Time, TT, to
//  Geocentric Coordinate Time, TCG.
func CgoTttcg(tt1, tt2 float64) (tcg1, tcg2 float64, err en.ErrNum) {
	var cTcg1, cTcg2 C.double
	cI := C.iauTttcg(C.double(tt1), C.double(tt2), &cTcg1, &cTcg2)
	if cI != 0 {
		err = errTttcg.Set(int(cI))
	}
	return float64(cTcg1), float64(cTcg2), err
}

//  GoTttcg Time scale transformation:  Terrestrial Time, TT, to
//  Geocentric Coordinate Time, TCG.
func GoTttcg(tt1, tt2 float64) (tcg1, tcg2 float64, err en.ErrNum) {

	// 1977 Jan 1 00:00:32.184 TT, as MJD
	const t77t = DJM77 + TTMTAI/DAYSEC

	// TT to TCG rate
	const elgg = ELG / (1.0 - ELG)

	// Result, safeguarding precision.
	if math.Abs(tt1) > math.Abs(tt2) {
		tcg1 = tt1
		tcg2 = tt2 + ((tt1-DJM0)+(tt2-t77t))*elgg
	} else {
		tcg1 = tt1 + ((tt2-DJM0)+(tt1-t77t))*elgg
		tcg2 = tt2
	}

	// Status (always OK).
	return
}
