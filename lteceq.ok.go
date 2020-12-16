package sofa

// #include "sofa.h"
import "C"

//  CgoLteceq Transformation from ecliptic coordinates (mean equinox and
//  ecliptic of date) to ICRS RA,Dec, using a long-term precession
//  model.
//
//  - - - - - - -
//   L t e c e q
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     epj     double     Julian epoch (TT)
//     dl,db   double     ecliptic longitude and latitude (radians)
//
//  Returned:
//     dr,dd   double     ICRS right ascension and declination (radians)
//
//  1) No assumptions are made about whether the coordinates represent
//     starlight and embody astrometric effects such as parallax or
//     aberration.
//
//  2) The transformation is approximately that from ecliptic longitude
//     and latitude (mean equinox and ecliptic of date) to mean J2000.0
//     right ascension and declination, with only frame bias (always
//     less than 25 mas) to disturb this classical picture.
//
//  3) The Vondrak et al. (2011, 2012) 400 millennia precession model
//     agrees with the IAU 2006 precession at J2000.0 and stays within
//     100 microarcseconds during the 20th and 21st centuries.  It is
//     accurate to a few arcseconds throughout the historical period,
//     worsening to a few tenths of a degree at the end of the
//     +/- 200,000 year time span.
//
//  Called:
//     iauS2c       spherical coordinates to unit vector
//     iauLtecm     J2000.0 to ecliptic rotation matrix, long term
//     iauTrxp      product of transpose of r-matrix and p-vector
//     iauC2s       unit vector to spherical coordinates
//     iauAnp       normalize angle into range 0 to 2pi
//     iauAnpm      normalize angle into range +/- pi
//
//  References:
//
//    Vondrak, J., Capitaine, N. and Wallace, P., 2011, New precession
//    expressions, valid for long time intervals, Astron.Astrophys. 534,
//    A22
//
//    Vondrak, J., Capitaine, N. and Wallace, P., 2012, New precession
//    expressions, valid for long time intervals (Corrigendum),
//    Astron.Astrophys. 541, C1
//
//  This revision:  2016 February 9
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoLteceq Transformation from ecliptic coordinates (mean equinox and
//  ecliptic of date) to ICRS RA,Dec, using a long-term precession
//  model.
func CgoLteceq(epj, dl, db float64) (dr, dd float64) {
	var cDr, cDd C.double
	C.iauLteceq(C.double(epj), C.double(dl), C.double(db),
		&cDr, &cDd)
	return float64(cDr), float64(cDd)
}

//  GoLteceq Transformation from ecliptic coordinates (mean equinox and
//  ecliptic of date) to ICRS RA,Dec, using a long-term precession
//  model.
func GoLteceq(epj, dl, db float64) (dr, dd float64) {
	var rm [3][3]float64
	var v1, v2 [3]float64
	var a, b float64

	// Spherical to Cartesian.
	v1 = GoS2c(dl, db)

	// Rotation matrix, ICRS equatorial to ecliptic.
	rm = GoLtecm(epj)

	// The transformation from ecliptic to ICRS.
	v2 = GoTrxp(rm, v1)

	// Cartesian to spherical.
	a, b = GoC2s(v2)

	// Express in conventional ranges.
	dr = GoAnp(a)
	dd = GoAnpm(b)
	return
}
