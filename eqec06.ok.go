package sofa

// #include "sofa.h"
import "C"

//  CgoEqec06 Transformation from ICRS equatorial coordinates to
//  ecliptic coordinates (mean equinox and ecliptic of date) using IAU
//  2006 precession model.
//
//  - - - - - - -
//   E q e c 0 6
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     date1,date2 double TT as a 2-part Julian date (Note 1)
//     dr,dd       double ICRS right ascension and declination (radians)
//
//  Returned:
//     dl,db       double ecliptic longitude and latitude (radians)
//
//  1) The TT date date1+date2 is a Julian Date, apportioned in any
//     convenient way between the two arguments.  For example,
//     JD(TT)=2450123.7 could be expressed in any of these ways,
//     among others:
//
//            date1          date2
//
//         2450123.7           0.0       (JD method)
//         2451545.0       -1421.3       (J2000 method)
//         2400000.5       50123.2       (MJD method)
//         2450123.5           0.2       (date & time method)
//
//     The JD method is the most natural and convenient to use in
//     cases where the loss of several decimal digits of resolution
//     is acceptable.  The J2000 method is best matched to the way
//     the argument is handled internally and will deliver the
//     optimum resolution.  The MJD method and the date & time methods
//     are both good compromises between resolution and convenience.
//
//  2) No assumptions are made about whether the coordinates represent
//     starlight and embody astrometric effects such as parallax or
//     aberration.
//
//  3) The transformation is approximately that from mean J2000.0 right
//     ascension and declination to ecliptic longitude and latitude
//     (mean equinox and ecliptic of date), with only frame bias (always
//     less than 25 mas) to disturb this classical picture.
//
//  Called:
//     iauS2c       spherical coordinates to unit vector
//     iauEcm06     J2000.0 to ecliptic rotation matrix, IAU 2006
//     iauRxp       product of r-matrix and p-vector
//     iauC2s       unit vector to spherical coordinates
//     iauAnp       normalize angle into range 0 to 2pi
//     iauAnpm      normalize angle into range +/- pi
//
//  This revision:  2016 February 9
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoEqec06 Transformation from ICRS equatorial coordinates to
//  ecliptic coordinates (mean equinox and ecliptic of date) using IAU
//  2006 precession model.
func CgoEqec06(date1, date2, dr, dd float64) (dl, db float64) {
	var cDl, cDb C.double
	C.iauEqec06(C.double(date1), C.double(date2), C.double(dr),
		C.double(dd), &cDl, &cDb)
	return float64(cDl), float64(cDb)
}

//  GoEqec06 Transformation from ICRS equatorial coordinates to
//  ecliptic coordinates (mean equinox and ecliptic of date) using IAU
//  2006 precession model.
func GoEqec06(date1, date2, dr, dd float64) (dl, db float64) {
	var rm [3][3]float64
	var v1, v2 [3]float64
	var a, b float64

	// Spherical to Cartesian.
	v1 = GoS2c(dr, dd)

	// Rotation matrix, ICRS equatorial to ecliptic.
	rm = GoEcm06(date1, date2)

	// The transformation from ICRS to ecliptic.
	v2 = GoRxp(rm, v1)

	// Cartesian to spherical.
	a, b = GoC2s(v2)

	// Express in conventional ranges.
	dl = GoAnp(a)
	db = GoAnpm(b)
	return
}
