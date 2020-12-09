package sofa

// #include "sofa.h"
import "C"

//  CgoAtciq Quick ICRS, epoch J2000.0, to CIRS transformation, given
//  precomputed star-independent astrometry parameters.
//
//  - - - - - -
//   A t c i q
//  - - - - - -
//
//  Use of this function is appropriate when efficiency is important and
//  where many star positions are to be transformed for one date.  The
//  star-independent parameters can be obtained by calling one of the
//  functions iauApci[13], iauApcg[13], iauApco[13] or iauApcs[13].
//
//  If the parallax and proper motions are zero the iauAtciqz function
//  can be used instead.
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     rc,dc  double     ICRS RA,Dec at J2000.0 (radians)
//     pr     double     RA proper motion (radians/year; Note 3)
//     pd     double     Dec proper motion (radians/year)
//     px     double     parallax (arcsec)
//     rv     double     radial velocity (km/s, +ve if receding)
//     astrom iauASTROM* star-independent astrometry parameters:
//      pmt    double       PM time interval (SSB, Julian years)
//      eb     double[3]    SSB to observer (vector, au)
//      eh     double[3]    Sun to observer (unit vector)
//      em     double       distance from Sun to observer (au)
//      v      double[3]    barycentric observer velocity (vector, c)
//      bm1    double       sqrt(1-|v|^2): reciprocal of Lorenz factor
//      bpn    double[3][3] bias-precession-nutation matrix
//      along  double       longitude + s' (radians)
//      xpl    double       polar motion xp wrt local meridian (radians)
//      ypl    double       polar motion yp wrt local meridian (radians)
//      sphi   double       sine of geodetic latitude
//      cphi   double       cosine of geodetic latitude
//      diurab double       magnitude of diurnal aberration vector
//      eral   double       "local" Earth rotation angle (radians)
//      refa   double       refraction constant A (radians)
//      refb   double       refraction constant B (radians)
//
//  Returned:
//     ri,di   double    CIRS RA,Dec (radians)
//
//  Notes:
//
//  1) All the vectors are with respect to BCRS axes.
//
//  2) Star data for an epoch other than J2000.0 (for example from the
//     Hipparcos catalog, which has an epoch of J1991.25) will require a
//     preliminary call to iauPmsafe before use.
//
//  3) The proper motion in RA is dRA/dt rather than cos(Dec)*dRA/dt.
//
//  Called:
//     iauPmpx      proper motion and parallax
//     iauLdsun     light deflection by the Sun
//     iauAb        stellar aberration
//     iauRxp       product of r-matrix and pv-vector
//     iauC2s       p-vector to spherical
//     iauAnp       normalize angle into range 0 to 2pi
//
//  This revision:   2013 October 9
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoAtciq Quick ICRS, epoch J2000.0, to CIRS transformation, given
//  precomputed star-independent astrometry parameters.
// void iauAtciq(double rc, double dc,
//               double pr, double pd, double px, double rv,
//               iauASTROM *astrom, double *ri, double *di)
func CgoAtciq(rc, dc, pr, pd, px, rv float64,
	astrom ASTROM) (ri, di float64) {
	var cRi, cDi C.double
	cAstrom := astrGo2C(astrom)
	C.iauAtciq(C.double(rc), C.double(dc), C.double(pr),
		C.double(pd), C.double(px), C.double(rv),
		&cAstrom, &cRi, &cDi)
	return float64(cRi), float64(cDi)
}

//  GoAtciq Quick ICRS, epoch J2000.0, to CIRS transformation, given
//  precomputed star-independent astrometry parameters.
func GoAtciq(rc, dc, pr, pd, px, rv float64,
	astrom ASTROM) (ri, di float64) {
	var pco, pnat, ppr, pi [3]float64
	var w float64

	// Proper motion and parallax, giving BCRS coordinate direction.
	pco = GoPmpx(rc, dc, pr, pd, px, rv, astrom.pmt, astrom.eb)

	// Light deflection by the Sun, giving BCRS natural direction.
	pnat = GoLdsun(pco, astrom.eh, astrom.em)

	// Aberration, giving GCRS proper direction.
	ppr = GoAb(pnat, astrom.v, astrom.em, astrom.bm1)

	// Bias-precession-nutation, giving CIRS proper direction.
	pi =  GoRxp(astrom.bpn, ppr)

	// CIRS RA,Dec.
	w, di = GoC2s(pi)
	ri = GoAnp(w)
	return
}
