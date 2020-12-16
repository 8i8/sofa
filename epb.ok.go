package sofa

// #include "sofa.h"
import "C"

//  CgoEpb Julian Date to Besselian Epoch.
//
//  - - - -
//   E p b
//  - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     dj1,dj2    double     Julian Date (see note)
//
//  Returned (function value):
//                double     Besselian Epoch.
//
//  Note:
//
//     The Julian Date is supplied in two pieces, in the usual SOFA
//     manner, which is designed to preserve time resolution.  The
//     Julian Date is available as a single number by adding dj1 and
//     dj2.  The maximum resolution is achieved if dj1 is 2451545.0
//     (J2000.0).
//
//  Reference:
//
//     Lieske, J.H., 1979. Astron.Astrophys., 73, 282.
//
//  This revision:  2013 August 21
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoEpb Julian Date to Besselian Epoch.
func CgoEpb(dj1, dj2 float64) (be float64) {
	var cF C.double
	cF = C.iauEpb(C.double(dj1), C.double(dj2))
	return float64(cF)
}

//  GoEpb Julian Date to Besselian Epoch.
func GoEpb(dj1, dj2 float64) (be float64) {

	// J2000.0-B1900.0 (2415019.81352) in days.
	const D1900 = 36524.68648

	return 1900.0 + ((dj1-DJ00)+(dj2+D1900))/DTY
}
