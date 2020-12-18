package sofa

// #include "sofa.h"
import "C"
import (
	"math"

	"github.com/8i8/sofa/en"
)

var errUt1tai = en.New(0, "Ut1tai", []string{
	"",
})

//  CgoUt1tai Time scale transformation:  Universal Time, UT1, to
//  International Atomic Time, TAI.
//
//  - - - - - - -
//   U t 1 t a i
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  canonical.
//
//  Given:
//     ut11,ut12  double    UT1 as a 2-part Julian Date
//     dta        double    UT1-TAI in seconds
//
//  Returned:
//     tai1,tai2  double    TAI as a 2-part Julian Date
//
//  Returned (function value):
//                int       status:  0 = OK
//
//  Notes:
//
//  1) ut11+ut12 is Julian Date, apportioned in any convenient way
//     between the two arguments, for example where ut11 is the Julian
//     Day Number and ut12 is the fraction of a day.  The returned
//     tai1,tai2 follow suit.
//
//  2) The argument dta, i.e. UT1-TAI, is an observed quantity, and is
//     available from IERS tabulations.
//
//  Reference:
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
//  CgoUt1tai Time scale transformation:  Universal Time, UT1, to
//  International Atomic Time, TAI.
func CgoUt1tai(ut11, ut12, dta float64) (
	tai1, tai2 float64, err en.ErrNum) {
	var cTai1, cTai2 C.double
	cI := C.iauUt1tai(C.double(ut11), C.double(ut12), C.double(dta),
		&cTai1, &cTai2)
	if n := int(cI); n != 0 {
		err = errUt1tai.Set(n)
	}
	return float64(cTai1), float64(cTai2), err
}

//  GoUt1tai Time scale transformation:  Universal Time, UT1, to
//  International Atomic Time, TAI.
func GoUt1tai(ut11, ut12, dta float64) (
	tai1, tai2 float64, err en.ErrNum) {

	var dtad float64

	// Result, safeguarding precision.
	dtad = dta / DAYSEC
	if math.Abs(ut11) > math.Abs(ut12) {
		tai1 = ut11
		tai2 = ut12 - dtad
	} else {
		tai1 = ut11 - dtad
		tai2 = ut12
	}

	// Status (always OK).
	return
}
