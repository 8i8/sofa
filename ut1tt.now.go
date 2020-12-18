package sofa

// #include "sofa.h"
import "C"
import "github.com/8i8/sofa/en"

var errUt1tt = en.New(0, "Ut1tt", []string{
	"",
})

//  CgoUt1tt Time scale transformation:  Universal Time, UT1, to
//  Terrestrial Time, TT.
//
//  - - - - - -
//   U t 1 t t
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  canonical.
//
//  Given:
//     ut11,ut12  double    UT1 as a 2-part Julian Date
//     dt         double    TT-UT1 in seconds
//
//  Returned:
//     tt1,tt2    double    TT as a 2-part Julian Date
//
//  Returned (function value):
//                int       status:  0 = OK
//
//  Notes:
//
//  1) ut11+ut12 is Julian Date, apportioned in any convenient way
//     between the two arguments, for example where ut11 is the Julian
//     Day Number and ut12 is the fraction of a day.  The returned
//     tt1,tt2 follow suit.
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
//  CgoUt1tt Time scale transformation:  Universal Time, UT1, to
//  Terrestrial Time, TT.
// int iauUt1tt(double ut11, double ut12, double dt,
//              double *tt1, double *tt2)
func CgoUt1tt(ut11, ut12, dt float64) (tt1, tt2 float64, err en.ErrNum) {
	var cTt1, cTt2 C.double
	cI := C.iauUt1tt(C.double(ut11), C.double(ut12), C.double(dt),
		&cTt1, &cTt2)
	if n := int(cI); n != 0 {
		err = errUt1tt.Set(n)
	}
	return float64(cTt1), float64(cTt2), err
}

// double dtd;

// /* Result, safeguarding precision. */
// dtd = dt / DAYSEC;
// if ( fabs(ut11) > fabs(ut12) ) {
//    *tt1 = ut11;
//    *tt2 = ut12 + dtd;
// } else {
//    *tt1 = ut11 + dtd;
//    *tt2 = ut12;
// }

// /* Status (always OK). */
// return 0;
