package sofa

// #include "sofa.h"
import "C"

//  CgoEpj Julian Date to Julian Epoch.
//
//  - - - -
//   E p j
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
//                double     Julian Epoch
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
//     Lieske, J.H., 1979, Astron.Astrophys. 73, 282.
//
//  This revision:  2013 August 7
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoEpj Julian Date to Julian Epoch.
func CgoEpj(dj1, dj2 float64) (epj float64) {
	var cEpj C.double
	cEpj = C.iauEpj(C.double(dj1), C.double(dj2))
	return float64(cEpj)
}

//  GoEpj Julian Date to Julian Epoch.
func GoEpj(dj1, dj2 float64) (epj float64) {

	epj = 2000.0 + ((dj1-DJ00)+dj2)/DJY
	return
}
