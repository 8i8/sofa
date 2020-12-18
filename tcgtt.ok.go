package sofa

// #include "sofa.h"
import "C"
import (
	"math"

	"github.com/8i8/sofa/en"
)

var errTcgtt = en.New(0, "Tcgtt", []string{
	"",
})

//  CgoTcgtt Time scale transformation:  Geocentric Coordinate Time,
//  TCG, to Terrestrial Time, TT.
//
//  - - - - - -
//   T c g t t
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  canonical.
//
//  Given:
//     tcg1,tcg2  double    TCG as a 2-part Julian Date
//
//  Returned:
//     tt1,tt2    double    TT as a 2-part Julian Date
//
//  Returned (function value):
//                int       status:  0 = OK
//
//  Note:
//
//     tcg1+tcg2 is Julian Date, apportioned in any convenient way
//     between the two arguments, for example where tcg1 is the Julian
//     Day Number and tcg22 is the fraction of a day.  The returned
//     tt1,tt2 follow suit.
//
//  References:
//
//     McCarthy, D. D., Petit, G. (eds.), IERS Conventions (2003),.
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
//  CgoTcgtt Time scale transformation:  Geocentric Coordinate Time,
//  TCG, to Terrestrial Time, TT.
func CgoTcgtt(tcg1, tcg2 float64) (tt1, tt2 float64, err en.ErrNum) {
	var cTt1, cTt2 C.double
	cI := C.iauTcgtt(C.double(tcg1), C.double(tcg2), &cTt1, &cTt2)
	if int(cI) != 0 {
		err = errTcgtt.Set(int(cI))
	}
	return float64(cTt1), float64(cTt2), err
}

//  GoTcgtt Time scale transformation:  Geocentric Coordinate Time,
//  TCG, to Terrestrial Time, TT.
func GoTcgtt(tcg1, tcg2 float64) (tt1, tt2 float64, err en.ErrNum) {

	// 1977 Jan 1 00:00:32.184 TT, as MJD
	const t77t = DJM77 + TTMTAI/DAYSEC

	// Result, safeguarding precision.
	if math.Abs(tcg1) > math.Abs(tcg2) {
		tt1 = tcg1
		tt2 = tcg2 - ((tcg1-DJM0)+(tcg2-t77t))*ELG
	} else {
		tt1 = tcg1 - ((tcg2-DJM0)+(tcg1-t77t))*ELG
		tt2 = tcg2
	}

	// OK status.
	return
}
