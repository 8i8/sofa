package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoPmpx Proper motion and parallax.
//
//  - - - - -
//   P m p x
//  - - - - -
//
//  Proper motion and parallax.
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     rc,dc  double     ICRS RA,Dec at catalog epoch (radians)
//     pr     double     RA proper motion (radians/year; Note 1)
//     pd     double     Dec proper motion (radians/year)
//     px     double     parallax (arcsec)
//     rv     double     radial velocity (km/s, +ve if receding)
//     pmt    double     proper motion time interval (SSB, Julian years)
//     pob    double[3]  SSB to observer vector (au)
//
//  Returned:
//     pco    double[3]  coordinate direction (BCRS unit vector)
//
//  Notes:
//
//  1) The proper motion in RA is dRA/dt rather than cos(Dec)*dRA/dt.
//
//  2) The proper motion time interval is for when the starlight
//     reaches the solar system barycenter.
//
//  3) To avoid the need for iteration, the Roemer effect (i.e. the
//     small annual modulation of the proper motion coming from the
//     changing light time) is applied approximately, using the
//     direction of the star at the catalog epoch.
//
//  References:
//
//     1984 Astronomical Almanac, pp B39-B41.
//
//     Urban, S. & Seidelmann, P. K. (eds), Explanatory Supplement to
//     the Astronomical Almanac, 3rd ed., University Science Books
//     (2013), Section 7.2.
//
//  Called:
//     iauPdp       scalar product of two p-vectors
//     iauPn        decompose p-vector into modulus and direction
//
//  This revision:   2013 October 9
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPmpx Proper motion and parallax.
func CgoPmpx(rc, dc, pr, pd, px, rv, pmt float64,
	pob [3]float64) (pco [3]float64) {

	var cPco [3]C.double
	cPob := v3sGo2C(pob)

	C.iauPmpx(C.double(rc), C.double(dc), C.double(pr),
		C.double(pd), C.double(px), C.double(rv), C.double(pmt),
		&cPob[0], &cPco[0])
	return v3sC2Go(cPco)
}

//  GoPmpx Proper motion and parallax.
func GoPmpx(rc, dc, pr, pd, px, rv, pmt float64,
	pob [3]float64) (pco [3]float64) {

	// Km/s to au/year
	const VF = DAYSEC * DJM / DAU

	// Light time for 1 au, Julian years
	const AULTY = AULT / DAYSEC / DJY

	var sr, cr, sd, cd, x, y, z, dt, pxr, w, pdz float64
	var p, pm [3]float64

	// Spherical coordinates to unit vector (and useful functions).
	sr = math.Sin(rc)
	cr = math.Cos(rc)
	sd = math.Sin(dc)
	cd = math.Cos(dc)
	x = cr * cd
	p[0] = x
	y = sr * cd
	p[1] = y
	z = sd
	p[2] = z

	// Proper motion time interval (y) including Roemer effect.
	dt = pmt + GoPdp(p, pob)*AULTY

	// Space motion (radians per year).
	pxr = px * DAS2R
	w = VF * rv * pxr
	pdz = pd * z
	pm[0] = -pr*y - pdz*cr + w*x
	pm[1] = pr*x - pdz*sr + w*y
	pm[2] = pd*cd + w*z

	// Coordinate direction of star (unit vector, BCRS).
	for i := 0; i < 3; i++ {
		p[i] += dt*pm[i] - pxr*pob[i]
	}
	_, pco = GoPn(p)

	return
}
