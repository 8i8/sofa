package sofa

// #include "sofa.h"
import "C"
import (
	"math"

	"github.com/8i8/sofa/en"
)

var errTdbtt = en.New(0, "Tdbtt", []string{
	"",
})

//  CgoTdbtt Time scale transformation:  Barycentric Dynamical Time,
//  TDB, to Terrestrial Time, TT.
//
//  - - - - - -
//   T d b t t
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  canonical.
//
//  Given:
//     tdb1,tdb2  double    TDB as a 2-part Julian Date
//     dtr        double    TDB-TT in seconds
//
//  Returned:
//     tt1,tt2    double    TT as a 2-part Julian Date
//
//  Returned (function value):
//                int       status:  0 = OK
//
//  Notes:
//
//  1) tdb1+tdb2 is Julian Date, apportioned in any convenient way
//     between the two arguments, for example where tdb1 is the Julian
//     Day Number and tdb2 is the fraction of a day.  The returned
//     tt1,tt2 follow suit.
//
//  2) The argument dtr represents the quasi-periodic component of the
//     GR transformation between TT and TCB.  It is dependent upon the
//     adopted solar-system ephemeris, and can be obtained by numerical
//     integration, by interrogating a precomputed time ephemeris or by
//     evaluating a model such as that implemented in the SOFA function
//     iauDtdb.   The quantity is dominated by an annual term of 1.7 ms
//     amplitude.
//
//  3) TDB is essentially the same as Teph, the time argument for the
//     JPL solar system ephemerides.
//
//  References:
//
//     McCarthy, D. D., Petit, G. (eds.), IERS Conventions (2003),
//     IERS Technical Note No. 32, BKG (2004)
//
//     IAU 2006 Resolution 3
//
//  This revision:  2019 June 20
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//
//  CgoTdbtt Time scale transformation:  Barycentric Dynamical Time,
//  TDB, to Terrestrial Time, TT.
func CgoTdbtt(tdb1, tdb2, dtr float64) (
	tt1, tt2 float64, err en.ErrNum) {

	var cTt1, cTt2 C.double
	cI := C.iauTdbtt(C.double(tdb1), C.double(tdb2), C.double(dtr),
		&cTt1, &cTt2)
	if int(cI) != 0 {
		err = errTdbtt.Set(int(cI))
	}
	return float64(cTt1), float64(cTt2), err
}

//  GoTdbtt Time scale transformation:  Barycentric Dynamical Time,
//  TDB, to Terrestrial Time, TT.
func GoTdbtt(tdb1, tdb2, dtr float64) (
	tt1, tt2 float64, err en.ErrNum) {

	var dtrd float64

	// Result, safeguarding precision.
	dtrd = dtr / DAYSEC
	if math.Abs(tdb1) > math.Abs(tdb2) {
		tt1 = tdb1
		tt2 = tdb2 - dtrd
	} else {
		tt1 = tdb1 - dtrd
		tt2 = tdb2
	}

	// Status (always OK).
	return
}
