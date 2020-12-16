package sofa

// #include "sofa.h"
import "C"

//  CgoFk524 Convert J2000.0 FK5 star catalog data to B1950.0 FK4.
//
//  - - - - - -
//   F k 5 2 4
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given: (all J2000.0, FK5)
//     r2000,d2000    double   J2000.0 RA,Dec (rad)
//     dr2000,dd2000  double   J2000.0 proper motions (rad/Jul.yr)
//     p2000          double   parallax (arcsec)
//     v2000          double   radial velocity (km/s, +ve = moving away)
//
//  Returned: (all B1950.0, FK4)
//     r1950,d1950    double   B1950.0 RA,Dec (rad)
//     dr1950,dd1950  double   B1950.0 proper motions (rad/trop.yr)
//     p1950          double   parallax (arcsec)
//     v1950          double   radial velocity (km/s, +ve = moving away)
//
//  Notes:
//
//  1) The proper motions in RA are dRA/dt rather than cos(Dec)*dRA/dt,
//     and are per year rather than per century.
//
//  2) The conversion is somewhat complicated, for several reasons:
//
//     . Change of standard epoch from J2000.0 to B1950.0.
//
//     . An intermediate transition date of 1984 January 1.0 TT.
//
//     . A change of precession model.
//
//     . Change of time unit for proper motion (Julian to tropical).
//
//     . FK4 positions include the E-terms of aberration, to simplify
//       the hand computation of annual aberration.  FK5 positions
//       assume a rigorous aberration computation based on the Earth's
//       barycentric velocity.
//
//     . The E-terms also affect proper motions, and in particular cause
//       objects at large distances to exhibit fictitious proper
//       motions.
//
//     The algorithm is based on Smith et al. (1989) and Yallop et al.
//     (1989), which presented a matrix method due to Standish (1982) as
//     developed by Aoki et al. (1983), using Kinoshita's development of
//     Andoyer's post-Newcomb precession.  The numerical constants from
//     Seidelmann (1992) are used canonically.
//
//  4) In the FK4 catalog the proper motions of stars within 10 degrees
//     of the poles do not embody differential E-terms effects and
//     should, strictly speaking, be handled in a different manner from
//     stars outside these regions.  However, given the general lack of
//     homogeneity of the star data available for routine astrometry,
//     the difficulties of handling positions that may have been
//     determined from astrometric fields spanning the polar and non-
//     polar regions, the likelihood that the differential E-terms
//     effect was not taken into account when allowing for proper motion
//     in past astrometry, and the undesirability of a discontinuity in
//     the algorithm, the decision has been made in this SOFA algorithm
//     to include the effects of differential E-terms on the proper
//     motions for all stars, whether polar or not.  At epoch J2000.0,
//     and measuring "on the sky" rather than in terms of RA change, the
//     errors resulting from this simplification are less than
//     1 milliarcsecond in position and 1 milliarcsecond per century in
//     proper motion.
//
//  Called:
//     iauAnp       normalize angle into range 0 to 2pi
//     iauPdp       scalar product of two p-vectors
//     iauPm        modulus of p-vector
//     iauPmp       p-vector minus p-vector
//     iauPpp       p-vector pluus p-vector
//     iauPv2s      pv-vector to spherical coordinates
//     iauS2pv      spherical coordinates to pv-vector
//     iauSxp       multiply p-vector by scalar
//
//  References:
//
//     Aoki, S. et al., 1983, "Conversion matrix of epoch B1950.0
//     FK4-based positions of stars to epoch J2000.0 positions in
//     accordance with the new IAU resolutions".  Astron.Astrophys.
//     128, 263-267.
//
//     Seidelmann, P.K. (ed), 1992, "Explanatory Supplement to the
//     Astronomical Almanac", ISBN 0-935702-68-7.
//
//     Smith, C.A. et al., 1989, "The transformation of astrometric
//     catalog systems to the equinox J2000.0".  Astron.J. 97, 265.
//
//     Standish, E.M., 1982, "Conversion of positions and proper motions
//     from B1950.0 to the IAU system at J2000.0".  Astron.Astrophys.,
//     115, 1, 20-22.
//
//     Yallop, B.D. et al., 1989, "Transformation of mean star places
//     from FK4 B1950.0 to FK5 J2000.0 using matrices in 6-space".
//     Astron.J. 97, 274.
//
//  This revision:   2019 October 3
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoFk524 Convert J2000.0 FK5 star catalog data to B1950.0 FK4.
func CgoFk524(r2000, d2000, dr2000, dd2000, p2000, v2000 float64) (
	r1950, d1950, dr1950, dd1950, p1950, v1950 float64) {
	var cR1950, cD1950, cDr1950, cDd1950, cP1950, cV1950 C.double
	C.iauFk524(C.double(r2000), C.double(d2000), C.double(dr2000),
		C.double(dd2000), C.double(p2000), C.double(v2000),
		&cR1950, &cD1950, &cDr1950, &cDd1950, &cP1950, &cV1950)
	return float64(cR1950), float64(cD1950), float64(cDr1950),
		float64(cDd1950), float64(cP1950), float64(cV1950)
}

//  GoFk524 Convert J2000.0 FK5 star catalog data to B1950.0 FK4.
func GoFk524(r2000, d2000, dr2000, dd2000, p2000, v2000 float64) (
	r1950, d1950, dr1950, dd1950, p1950, v1950 float64) {
	// Radians per year to arcsec per century
	const PMF = 100.0 * DR2AS

	// Small number to avoid arithmetic problems
	const TINY = 1e-30

	// Miscellaneous
	var r, d, ur, ud, px, rv, pxvf, w, rd float64
	var i, j, k, l int

	// Vectors, p and pv
	var r0, r1, pv [2][3]float64
	var p1, p2 [3]float64

	// CANONICAL CONSTANTS (Seidelmann 1992)

	// Km per sec to AU per tropical century
	// = 86400 * 36524.2198782 / 149597870.7
	const VF = 21.095

	// Constant pv-vector (cf. Seidelmann 3.591-2, vectors A and
	// Adot)
	var a = [2][3]float64{
		{-1.62557e-6, -0.31919e-6, -0.13843e-6},
		{+1.245e-3, -1.580e-3, -0.659e-3},
	}

	// 3x2 matrix of pv-vectors (cf. Seidelmann 3.592-1, matrix
	// M^-1)
	var em = [2][3][2][3]float64{

		{{{+0.9999256795, +0.0111814828, +0.0048590039},
			{-0.00000242389840, -0.00000002710544, -0.00000001177742}},

			{{-0.0111814828, +0.9999374849, -0.0000271771},
				{+0.00000002710544, -0.00000242392702, +0.00000000006585}},

			{{-0.0048590040, -0.0000271557, +0.9999881946},
				{+0.00000001177742, +0.00000000006585, -0.00000242404995}}},

		{{{-0.000551, +0.238509, -0.435614},
			{+0.99990432, +0.01118145, +0.00485852}},

			{{-0.238560, -0.002667, +0.012254},
				{-0.01118145, +0.99991613, -0.00002717}},

			{{+0.435730, -0.008541, +0.002117},
				{-0.00485852, -0.00002716, +0.99996684}}},
	}

	//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

	// The FK5 data (units radians and arcsec per Julian century).
	r = r2000
	d = d2000
	ur = dr2000 * PMF
	ud = dd2000 * PMF
	px = p2000
	rv = v2000

	// Express as a pv-vector.
	pxvf = px * VF
	w = rv * pxvf
	r0 = GoS2pv(r, d, 1.0, ur, ud, w)

	// Convert pv-vector to Bessel-Newcomb system (cf. Seidelmann
	// 3.592-1).
	for i = 0; i < 2; i++ {
		for j = 0; j < 3; j++ {
			w = 0.0
			for k = 0; k < 2; k++ {
				for l = 0; l < 3; l++ {
					w += em[i][j][k][l] * r0[k][l]
				}
			}
			r1[i][j] = w
		}
	}

	// Apply E-terms (equivalent to Seidelmann 3.592-3, one
	// iteration).

	// Direction.
	w = GoPm(r1[0])
	p1 = GoSxp(GoPdp(r1[0], a[0]), r1[0])
	p2 = GoSxp(w, a[0])
	p1 = GoPmp(p2, p1)
	p1 = GoPpp(r1[0], p1)

	// Recompute length.
	w = GoPm(p1)

	// Direction.
	p1 = GoSxp(GoPdp(r1[0], a[0]), r1[0])
	p2 = GoSxp(w, a[0])
	p1 = GoPmp(p2, p1)
	pv[0] = GoPpp(r1[0], p1)

	// Derivative.
	p1 = GoSxp(GoPdp(r1[0], a[1]), pv[0])
	p2 = GoSxp(w, a[1])
	p1 = GoPmp(p2, p1)
	pv[1] = GoPpp(r1[1], p1)

	// Revert to catalog form.
	r, d, w, ur, ud, rd = GoPv2s(pv)
	if px > TINY {
		rv = rd / pxvf
		px = px / w
	}

	// Return the results.
	r1950 = GoAnp(r)
	d1950 = d
	dr1950 = ur / PMF
	dd1950 = ud / PMF
	p1950 = px
	v1950 = rv

	// Finished.
	return
}
