package sofa

// #include "sofa.h"
import "C"

//  Apci13 For a terrestrial observer, prepare star-independent astrometry
//  parameters for transformations between ICRS and geocentric CIRS
//  coordinates.  The caller supplies the date, and SOFA models are used
//  to predict the Earth ephemeris and CIP/CIO.
//
//  - - - - - - -
//   A p c i 1 3
//  - - - - - - -
//
//  The parameters produced by this function are required in the
//  parallax, light deflection, aberration, and bias-precession-nutation
//  parts of the astrometric transformation chain.
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     date1  double      TDB as a 2-part...
//     date2  double      ...Julian Date (Note 1)
//
//  Returned:
//     astrom iauASTROM*  star-independent astrometry parameters:
//      pmt    double       PM time interval (SSB, Julian years)
//      eb     double[3]    SSB to observer (vector, au)
//      eh     double[3]    Sun to observer (unit vector)
//      em     double       distance from Sun to observer (au)
//      v      double[3]    barycentric observer velocity (vector, c)
//      bm1    double       sqrt(1-|v|^2): reciprocal of Lorenz factor
//      bpn    double[3][3] bias-precession-nutation matrix
//      along  double       unchanged
//      xpl    double       unchanged
//      ypl    double       unchanged
//      sphi   double       unchanged
//      cphi   double       unchanged
//      diurab double       unchanged
//      eral   double       unchanged
//      refa   double       unchanged
//      refb   double       unchanged
//     eo     double*     equation of the origins (ERA-GST)
//
//  Notes:
//
//  1) The TDB date date1+date2 is a Julian Date, apportioned in any
//     convenient way between the two arguments.  For example,
//     JD(TDB)=2450123.7 could be expressed in any of these ways, among
//     others:
//
//            date1          date2
//
//         2450123.7           0.0       (JD method)
//         2451545.0       -1421.3       (J2000 method)
//         2400000.5       50123.2       (MJD method)
//         2450123.5           0.2       (date & time method)
//
//     The JD method is the most natural and convenient to use in cases
//     where the loss of several decimal digits of resolution is
//     acceptable.  The J2000 method is best matched to the way the
//     argument is handled internally and will deliver the optimum
//     resolution.  The MJD method and the date & time methods are both
//     good compromises between resolution and convenience.  For most
//     applications of this function the choice will not be at all
//     critical.
//
//     TT can be used instead of TDB without any significant impact on
//     accuracy.
//
//  2) All the vectors are with respect to BCRS axes.
//
//  3) In cases where the caller wishes to supply his own Earth
//     ephemeris and CIP/CIO, the function iauApci can be used instead
//     of the present function.
//
//  4) This is one of several functions that inserts into the astrom
//     structure star-independent parameters needed for the chain of
//     astrometric transformations ICRS <-> GCRS <-> CIRS <-> observed.
//
//     The various functions support different classes of observer and
//     portions of the transformation chain:
//
//          functions         observer        transformation
//
//       iauApcg iauApcg13    geocentric      ICRS <-> GCRS
//       iauApci iauApci13    terrestrial     ICRS <-> CIRS
//       iauApco iauApco13    terrestrial     ICRS <-> observed
//       iauApcs iauApcs13    space           ICRS <-> GCRS
//       iauAper iauAper13    terrestrial     update Earth rotation
//       iauApio iauApio13    terrestrial     CIRS <-> observed
//
//     Those with names ending in "13" use contemporary SOFA models to
//     compute the various ephemerides.  The others accept ephemerides
//     supplied by the caller.
//
//     The transformation from ICRS to GCRS covers space motion,
//     parallax, light deflection, and aberration.  From GCRS to CIRS
//     comprises frame bias and precession-nutation.  From CIRS to
//     observed takes account of Earth rotation, polar motion, diurnal
//     aberration and parallax (unless subsumed into the ICRS <-> GCRS
//     transformation), and atmospheric refraction.
//
//  5) The context structure astrom produced by this function is used by
//     iauAtciq* and iauAticq*.
//
//  Called:
//     iauEpv00     Earth position and velocity
//     iauPnm06a    classical NPB matrix, IAU 2006/2000A
//     iauBpn2xy    extract CIP X,Y coordinates from NPB matrix
//     iauS06       the CIO locator s, given X,Y, IAU 2006
//     iauApci      astrometry parameters, ICRS-CIRS
//     iauEors      equation of the origins, given NPB matrix and s
//
//  This revision:   2013 October 9
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
func Apci13(date1, date2 float64) (astrom ASTROM, eo float64) {
	var cR [3][3]C.double
	var cEhpv, cEbpv [2][3]C.double
	var x, y, s C.double
	var cAstrom C.iauASTROM

	/* Earth barycentric & heliocentric position/velocity (au, au/d). */
	C.iauEpv00(C.double(date1), C.double(date2), &cEhpv[0], &cEbpv[0])

	/* Form the equinox based BPN matrix, IAU 2006/2000A. */
	C.iauPnm06a(C.double(date1), C.double(date2), &cR[0])

	/* Extract CIP X,Y. */
	C.iauBpn2xy(&cR[0], &x, &y)

	/* Obtain CIO locator s. */
	s = C.iauS06(C.double(date1), C.double(date2), x, y)

	/* Compute the star-independent astrometry parameters. */
	C.iauApci(C.double(date1), C.double(date2), &cEbpv[0], &cEhpv[0][0], x, y, s, &cAstrom)

	/* Equation of the origins. */
	eo = float64(C.iauEors(&cR[0], s))
	astrom = astrC2Go(cAstrom)
	return
}

// func goApci13() {
//     var r [3][3]float64
//     var ehpv, ebpv [2][3]float64
//     var x, y, s float64

// /* Earth barycentric & heliocentric position/velocity (au, au/d). */
//    (void) iauEpv00(date1, date2, ehpv, ebpv);

// /* Form the equinox based BPN matrix, IAU 2006/2000A. */
//    iauPnm06a(date1, date2, r);

// /* Extract CIP X,Y. */
//    iauBpn2xy(r, &x, &y);

// /* Obtain CIO locator s. */
//    s = iauS06(date1, date2, x, y);

// /* Compute the star-independent astrometry parameters. */
//    iauApci(date1, date2, ebpv, ehpv[0], x, y, s, astrom);

// /* Equation of the origins. */
//    *eo = iauEors(r, s);

// /* Finished. */
// }
