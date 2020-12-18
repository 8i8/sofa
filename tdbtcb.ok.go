package sofa

// #include "sofa.h"
import "C"
import (
	"math"

	"github.com/8i8/sofa/en"
)

var errTdbtcb = en.New(0, "Tdbtcb", []string{
	"",
})

//  CgoTdbtcb Time scale transformation:  Barycentric Dynamical Time,
//  TDB, to Barycentric Coordinate Time, TCB.
//
//  - - - - - - -
//   T d b t c b
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  canonical.
//
//  Given:
//     tdb1,tdb2  double    TDB as a 2-part Julian Date
//
//  Returned:
//     tcb1,tcb2  double    TCB as a 2-part Julian Date
//
//  Returned (function value):
//                int       status:  0 = OK
//
//  Notes:
//
//  1) tdb1+tdb2 is Julian Date, apportioned in any convenient way
//     between the two arguments, for example where tdb1 is the Julian
//     Day Number and tdb2 is the fraction of a day.  The returned
//     tcb1,tcb2 follow suit.
//
//  2) The 2006 IAU General Assembly introduced a conventional linear
//     transformation between TDB and TCB.  This transformation
//     compensates for the drift between TCB and terrestrial time TT,
//     and keeps TDB approximately centered on TT.  Because the
//     relationship between TT and TCB depends on the adopted solar
//     system ephemeris, the degree of alignment between TDB and TT over
//     long intervals will vary according to which ephemeris is used.
//     Former definitions of TDB attempted to avoid this problem by
//     stipulating that TDB and TT should differ only by periodic
//     effects.  This is a good description of the nature of the
//     relationship but eluded precise mathematical formulation.  The
//     conventional linear relationship adopted in 2006 sidestepped
//     these difficulties whilst delivering a TDB that in practice was
//     consistent with values before that date.
//
//  3) TDB is essentially the same as Teph, the time argument for the
//     JPL solar system ephemerides.
//
//  Reference:
//
//     IAU 2006 Resolution B3
//
//  This revision:  2019 June 20
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoTdbtcb Time scale transformation:  Barycentric Dynamical Time,
//  TDB, to Barycentric Coordinate Time, TCB.
func CgoTdbtcb(tdb1, tdb2 float64) (tcb1, tcb2 float64, err en.ErrNum) {
	var cTcb1, cTcb2 C.double
	cI := C.iauTdbtcb(C.double(tdb1), C.double(tdb2), &cTcb1, &cTcb2)
	if n := int(cI); n != 0 {
		err = errTdbtcb.Set(n)
	}
	return float64(cTcb1), float64(cTcb2), err
}

//  GoTdbtcb Time scale transformation:  Barycentric Dynamical Time,
//  TDB, to Barycentric Coordinate Time, TCB.
func GoTdbtcb(tdb1, tdb2 float64) (tcb1, tcb2 float64, err en.ErrNum) {
	// 1977 Jan 1 00:00:32.184 TT, as two-part JD
	const t77td = DJM0 + DJM77
	const t77tf = TTMTAI / DAYSEC

	// TDB (days) at TAI 1977 Jan 1.0
	const tdb0 = TDB0 / DAYSEC

	// TDB to TCB rate
	const elbb = ELB / (1.0 - ELB)

	var d, f float64

	// Result, preserving date format but safeguarding precision.
	if math.Abs(tdb1) > math.Abs(tdb2) {
		d = t77td - tdb1
		f = tdb2 - tdb0
		tcb1 = tdb1
		tcb2 = f - (d-(f-t77tf))*elbb
	} else {
		d = t77td - tdb2
		f = tdb1 - tdb0
		tcb1 = f - (d-(f-t77tf))*elbb
		tcb2 = tdb2
	}

	// Status (always OK).
	return
}
