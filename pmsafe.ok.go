package sofa

// #include "sofa.h"
import "C"

import "github.com/8i8/sofa/en"

var errPmsafe = en.New(1, "Pmsafe", []string{
	"system error (should not occur)",
	"",
	"distance overridden (Note 6)",
	"excessive velocity (Note 7)",
	"solution didn't converge (Note 8)",
	"binary logical OR of the above warnings",
})

//  CgoPmsafe Star proper motion:  update star catalog data for space
//  motion, with special handling to handle the zero parallax case.
//
//  - - - - - - -
//   P m s a f e
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     ra1    double      right ascension (radians), before
//     dec1   double      declination (radians), before
//     pmr1   double      RA proper motion (radians/year), before
//     pmd1   double      Dec proper motion (radians/year), before
//     px1    double      parallax (arcseconds), before
//     rv1    double      radial velocity (km/s, +ve = receding), before
//     ep1a   double      "before" epoch, part A (Note 1)
//     ep1b   double      "before" epoch, part B (Note 1)
//     ep2a   double      "after" epoch, part A (Note 1)
//     ep2b   double      "after" epoch, part B (Note 1)
//
//  Returned:
//     ra2    double      right ascension (radians), after
//     dec2   double      declination (radians), after
//     pmr2   double      RA proper motion (radians/year), after
//     pmd2   double      Dec proper motion (radians/year), after
//     px2    double      parallax (arcseconds), after
//     rv2    double      radial velocity (km/s, +ve = receding), after
//
//  Returned (function value):
//            int         status:
//                         -1 = system error (should not occur)
//                          0 = no warnings or errors
//                          1 = distance overridden (Note 6)
//                          2 = excessive velocity (Note 7)
//                          4 = solution didn't converge (Note 8)
//                       else = binary logical OR of the above warnings
//
//  Notes:
//
//  1) The starting and ending TDB epochs ep1a+ep1b and ep2a+ep2b are
//     Julian Dates, apportioned in any convenient way between the two
//     parts (A and B).  For example, JD(TDB)=2450123.7 could be
//     expressed in any of these ways, among others:
//
//            epNa            epNb
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
//     good compromises between resolution and convenience.
//
//  2) In accordance with normal star-catalog conventions, the object's
//     right ascension and declination are freed from the effects of
//     secular aberration.  The frame, which is aligned to the catalog
//     equator and equinox, is Lorentzian and centered on the SSB.
//
//     The proper motions are the rate of change of the right ascension
//     and declination at the catalog epoch and are in radians per TDB
//     Julian year.
//
//     The parallax and radial velocity are in the same frame.
//
//  3) Care is needed with units.  The star coordinates are in radians
//     and the proper motions in radians per Julian year, but the
//     parallax is in arcseconds.
//
//  4) The RA proper motion is in terms of coordinate angle, not true
//     angle.  If the catalog uses arcseconds for both RA and Dec proper
//     motions, the RA proper motion will need to be divided by cos(Dec)
//     before use.
//
//  5) Straight-line motion at constant speed, in the inertial frame, is
//     assumed.
//
//  6) An extremely small (or zero or negative) parallax is overridden
//     to ensure that the object is at a finite but very large distance,
//     but not so large that the proper motion is equivalent to a large
//     but safe speed (about 0.1c using the chosen constant).  A warning
//     status of 1 is added to the status if this action has been taken.
//
//  7) If the space velocity is a significant fraction of c (see the
//     constant VMAX in the function iauStarpv), it is arbitrarily set
//     to zero.  When this action occurs, 2 is added to the status.
//
//  8) The relativistic adjustment carried out in the iauStarpv function
//     involves an iterative calculation.  If the process fails to
//     converge within a set number of iterations, 4 is added to the
//     status.
//
//  Called:
//     iauSeps      angle between two points
//     iauStarpm    update star catalog data for space motion
//
//  This revision:   2014 July 1
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPmsafe Star proper motion:  update star catalog data for space
//  motion, with special handling to handle the zero parallax case.
func CgoPmsafe(ra1, dec1, pmr1, pmd1, px1,
	rv1, ep1a, ep1b, ep2a, ep2b float64) (
	ra2, dec2, pmr2, pmd2, px2, rv2 float64, err en.ErrNum) {
	var cRa2, cDec2, cPmr2, cPmd2, cPx2, cRv2 C.double
	cI := C.iauPmsafe(C.double(ra1), C.double(dec1), C.double(pmr1),
		C.double(pmd1), C.double(px1), C.double(rv1),
		C.double(ep1a), C.double(ep1b), C.double(ep2a),
		C.double(ep2b), &cRa2, &cDec2, &cPmr2, &cPmd2, &cPx2,
		&cRv2)
	if n := int(cI); n != 0 {
		err = errPmsafe.Set(n)
	}
	return float64(cRa2), float64(cDec2), float64(cPmr2),
		float64(cPmd2), float64(cPx2), float64(cRv2), err
}

//  GoPmsafe Star proper motion:  update star catalog data for space
//  motion, with special handling to handle the zero parallax case.
func GoPmsafe(ra1, dec1, pmr1, pmd1, px1,
	rv1, ep1a, ep1b, ep2a, ep2b float64) (
	ra2, dec2, pmr2, pmd2, px2, rv2 float64, err en.ErrNum) {
	// Minimum allowed parallax (arcsec)
	const PXMIN = 5e-7

	// Factor giving maximum allowed transverse speed of about 1% c
	const F = 326.0

	//var jpx, j int
	var err1 en.ErrNum
	var pm, px1a float64

	// Proper motion in one year (radians).
	pm = GoSeps(ra1, dec1, ra1+pmr1, dec1+pmd1)

	// Override the parallax to reduce the chances of a warning
	// status.
	//jpx = 0
	px1a = px1
	pm *= F
	if px1a < pm {
		//jpx = 1
		err = errPmsafe.Set(1)
		px1a = pm
	}
	if px1a < PXMIN {
		//jpx = 1
		err = errPmsafe.Set(1)
		px1a = PXMIN
	}

	// Carry out the transformation using the modified parallax.
	ra2, dec2, pmr2, pmd2, px2, rv2, err = GoStarpm(
		ra1, dec1, pmr1, pmd1, px1a, rv1,
		ep1a, ep1b, ep2a, ep2b)

	// Revise and return the status.
	if err != nil && (err.Is()%2) == 0 {
		if err1 != nil {
			err = err.Add(err, err1.Is())
			//j += jpx
		}
	}
	return
}
