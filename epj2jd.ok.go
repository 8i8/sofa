package sofa

// #include "sofa.h"
import "C"

//  CgoEpj2jd Julian Epoch to Julian Date.
//
//  - - - - - - -
//   E p j 2 j d
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     epj      double    Julian Epoch (e.g. 1996.8)
//
//  Returned:
//     djm0     double    MJD zero-point: always 2400000.5
//     djm      double    Modified Julian Date
//
//  Note:
//
//     The Julian Date is returned in two pieces, in the usual SOFA
//     manner, which is designed to preserve time resolution.  The
//     Julian Date is available as a single number by adding djm0 and
//     djm.
//
//  Reference:
//
//     Lieske, J.H., 1979, Astron.Astrophys. 73, 282.
//
//  This revision:  2013 August 7
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoEpj2jd Julian Epoch to Julian Date.
func CgoEpj2jd(epj float64) (djm0, djm float64) {
	var cDjm0, cDjm C.double
	C.iauEpj2jd(C.double(epj), &cDjm0, &cDjm)
	return float64(cDjm0), float64(cDjm)
}

//  GoEpj2jd Julian Epoch to Julian Date.
func GoEpj2jd(epj float64) (djm0, djm float64) {

	djm0 = DJM0
	djm = DJM00 + (epj-2000.0)*365.25
	return
}
