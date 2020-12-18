package sofa

// #include "sofa.h"
import "C"
import (
	"math"

	"github.com/8i8/sofa/en"
)

var errTtut1 = en.New(0, "Ttut1", []string{
	"",
})

//  CgoTtut1 Time scale transformation:  Terrestrial Time, TT, to
//  Universal Time, UT1.
//
//  - - - - - -
//   T t u t 1
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  canonical.
//
//  Given:
//     tt1,tt2    double    TT as a 2-part Julian Date
//     dt         double    TT-UT1 in seconds
//
//  Returned:
//     ut11,ut12  double    UT1 as a 2-part Julian Date
//
//  Returned (function value):
//                int       status:  0 = OK
//
//  Notes:
//
//  1) tt1+tt2 is Julian Date, apportioned in any convenient way between
//     the two arguments, for example where tt1 is the Julian Day Number
//     and tt2 is the fraction of a day.  The returned ut11,ut12 follow
//     suit.
//
//  2) The argument dt is classical Delta T.
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
//  CgoTtut1 Time scale transformation:  Terrestrial Time, TT, to
//  Universal Time, UT1.
func CgoTtut1(tt1, tt2, dt float64) (ut11, ut12 float64, err en.ErrNum) {
	var cUt11, cUt12 C.double
	cI := C.iauTtut1(C.double(tt1), C.double(tt2), C.double(dt), &cUt11,
		&cUt12)
	if int(cI) != 0 {
		err = errTtut1.Set(int(cI))
	}
	return float64(cUt11), float64(cUt12), err
}

//  GoTtut1 Time scale transformation:  Terrestrial Time, TT, to
//  Universal Time, UT1.
func GoTtut1(tt1, tt2, dt float64) (ut11, ut12 float64, err en.ErrNum) {

	var dtd float64

	// Result, safeguarding precision.
	dtd = dt / DAYSEC
	if math.Abs(tt1) > math.Abs(tt2) {
		ut11 = tt1
		ut12 = tt2 - dtd
	} else {
		ut11 = tt1 - dtd
		ut12 = tt2
	}

	// Status (always OK).
	return
}
