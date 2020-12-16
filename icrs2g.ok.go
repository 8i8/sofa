package sofa

// #include "sofa.h"
import "C"

//  CgoIcrs2g Transformation from ICRS to Galactic Coordinates.
//
//  - - - - - - -
//   I c r s 2 g
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     dr     double      ICRS right ascension (radians)
//     dd     double      ICRS declination (radians)
//
//  Returned:
//     dl     double      galactic longitude (radians)
//     db     double      galactic latitude (radians)
//
//  Notes:
//
//  1) The IAU 1958 system of Galactic coordinates was defined with
//     respect to the now obsolete reference system FK4 B1950.0.  When
//     interpreting the system in a modern context, several factors have
//     to be taken into account:
//
//     . The inclusion in FK4 positions of the E-terms of aberration.
//
//     . The distortion of the FK4 proper motion system by differential
//       Galactic rotation.
//
//     . The use of the B1950.0 equinox rather than the now-standard
//       J2000.0.
//
//     . The frame bias between ICRS and the J2000.0 mean place system.
//
//     The Hipparcos Catalogue (Perryman & ESA 1997) provides a rotation
//     matrix that transforms directly between ICRS and Galactic
//     coordinates with the above factors taken into account.  The
//     matrix is derived from three angles, namely the ICRS coordinates
//     of the Galactic pole and the longitude of the ascending node of
//     the galactic equator on the ICRS equator.  They are given in
//     degrees to five decimal places and for canonical purposes are
//     regarded as exact.  In the Hipparcos Catalogue the matrix
//     elements are given to 10 decimal places (about 20 microarcsec).
//     In the present SOFA function the matrix elements have been
//     recomputed from the canonical three angles and are given to 30
//     decimal places.
//
//  2) The inverse transformation is performed by the function iauG2icrs.
//
//  Called:
//     iauAnp       normalize angle into range 0 to 2pi
//     iauAnpm      normalize angle into range +/- pi
//     iauS2c       spherical coordinates to unit vector
//     iauRxp       product of r-matrix and p-vector
//     iauC2s       p-vector to spherical
//
//  Reference:
//     Perryman M.A.C. & ESA, 1997, ESA SP-1200, The Hipparcos and Tycho
//     catalogues.  Astrometric and photometric star catalogues
//     derived from the ESA Hipparcos Space Astrometry Mission.  ESA
//     Publications Division, Noordwijk, Netherlands.
//
//  This revision:   2018 January 2
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoIcrs2g Transformation from ICRS to Galactic Coordinates.
// void iauIcrs2g ( double dr, double dd, double *dl, double *db )
func CgoIcrs2g(dr, dd float64) (dl, db float64) {
	var cDl, cDb C.double
	C.iauIcrs2g(C.double(dr), C.double(dd),
		&cDl, &cDb)
	return float64(cDl), float64(cDb)
}

//  GoIcrs2g Transformation from ICRS to Galactic Coordinates.
// void iauIcrs2g ( double dr, double dd, double *dl, double *db )
func GoIcrs2g(dr, dd float64) (dl, db float64) {
	var v1, v2 [3]float64

	//
	//  L2,B2 system of galactic coordinates in the form presented
	//  in the Hipparcos Catalogue.  In degrees:
	//
	//  P = 192.85948    right ascension of the Galactic north pole
	//  in ICRS Q =  27.12825    declination of the Galactic north
	//  pole in ICRS R =  32.93192    longitude of the ascending
	//  node of the Galactic plane on the ICRS equator
	//
	//  ICRS to galactic rotation matrix, obtained by computing
	//  R_3(-R) R_1(pi/2-Q) R_3(pi/2+P) to the full precision shown:

	var r = [3][3]float64{{-0.054875560416215368492398900454,
		-0.873437090234885048760383168409,
		-0.483835015548713226831774175116},
		{+0.494109427875583673525222371358,
			-0.444829629960011178146614061616,
			+0.746982244497218890527388004556},
		{-0.867666149019004701181616534570,
			-0.198076373431201528180486091412,
			+0.455983776175066922272100478348}}

	// Spherical to Cartesian.
	v1 = GoS2c(dr, dd)

	// ICRS to Galactic.
	v2 = GoRxp(r, v1)

	// Cartesian to spherical.
	dl, db = GoC2s(v2)

	// Express in conventional ranges.
	dl = GoAnp(dl)
	db = GoAnpm(db)
	return
}
