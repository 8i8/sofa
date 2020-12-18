package sofa

// #include "sofa.h"
import "C"
import (
	"math"

	"github.com/8i8/sofa/en"
)

var errTcbtdb = en.New(0, "Tcbtdb", []string{
	"",
})

//  CgoTcbtdb Time scale transformation:  Barycentric Coordinate Time,
//  TCB, to Barycentric Dynamical Time, TDB.
//
//  - - - - - - -
//   T c b t d b
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  canonical.
//
//  Given:
//     tcb1,tcb2  double    TCB as a 2-part Julian Date
//
//  Returned:
//     tdb1,tdb2  double    TDB as a 2-part Julian Date
//
//  Returned (function value):
//                int       status:  0 = OK
//
//  Notes:
//
//  1) tcb1+tcb2 is Julian Date, apportioned in any convenient way
//     between the two arguments, for example where tcb1 is the Julian
//     Day Number and tcb2 is the fraction of a day.  The returned
//     tdb1,tdb2 follow suit.
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
//  CgoTcbtdb Time scale transformation:  Barycentric Coordinate Time,
//  TCB, to Barycentric Dynamical Time, TDB.
func CgoTcbtdb(tcb1, tcb2 float64) (tdb1, tdb2 float64, err en.ErrNum) {
	var cTdb1, cTdb2 C.double
	cI := C.iauTcbtdb(C.double(tcb1), C.double(tcb2), &cTdb1, &cTdb2)
	if n := int(cI); n != 0 {
		err = errTcbtdb.Set(n)
	}
	return float64(cTdb1), float64(cTdb2), err
}

//  GoTcbtdb Time scale transformation:  Barycentric Coordinate Time,
//  TCB, to Barycentric Dynamical Time, TDB.
func GoTcbtdb(tcb1, tcb2 float64) (tdb1, tdb2 float64, err en.ErrNum) {
	// 1977 Jan 1 00:00:32.184 TT, as two-part JD
	const t77td = DJM0 + DJM77
	const t77tf = TTMTAI / DAYSEC

	// TDB (days) at TAI 1977 Jan 1.0
	const tdb0 = TDB0 / DAYSEC

	var d float64

	// Result, safeguarding precision.
	if math.Abs(tcb1) > math.Abs(tcb2) {
		d = tcb1 - t77td
		tdb1 = tcb1
		tdb2 = tcb2 + tdb0 - (d+(tcb2-t77tf))*ELB
	} else {
		d = tcb2 - t77td
		tdb1 = tcb1 + tdb0 - (d+(tcb1-t77tf))*ELB
		tdb2 = tcb2
	}

	// Status (always OK).
	return
}
