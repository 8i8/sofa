package sofa

// #include "sofa.h"
import "C"

//  CgoEpb2jd Besselian Epoch to Julian Date.
//
//  - - - - - - -
//   E p b 2 j d
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     epb      double    Besselian Epoch (e.g. 1957.3)
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
//  This revision:  2013 August 13
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoEpb2jd Besselian Epoch to Julian Date.
func CgoEpb2jd(epb float64) (djm0, djm float64) {
	var cDjm0, cDjm C.double
	C.iauEpb2jd(C.double(epb), &cDjm0, &cDjm)
	return float64(cDjm0), float64(cDjm)
}

//  GoEpb2jd Besselian Epoch to Julian Date.
func GoEpb2jd(epb float64) (djm0, djm float64) {

	djm0 = DJM0
	djm = 15019.81352 + (epb-1900.0)*DTY

	return
}
