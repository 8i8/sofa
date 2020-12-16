package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoLtpecl Long-term precession of the ecliptic.
//
//  - - - - - - -
//   L t p e c l
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     epj     double         Julian epoch (TT)
//
//  Returned:
//     vec     double[3]      ecliptic pole unit vector
//
//  Notes:
//
//  1) The returned vector is with respect to the J2000.0 mean equator
//     and equinox.
//
//  2) The Vondrak et al. (2011, 2012) 400 millennia precession model
//     agrees with the IAU 2006 precession at J2000.0 and stays within
//     100 microarcseconds during the 20th and 21st centuries.  It is
//     accurate to a few arcseconds throughout the historical period,
//     worsening to a few tenths of a degree at the end of the
//     +/- 200,000 year time span.
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
//  CgoLtpecl Long-term precession of the ecliptic.
func CgoLtpecl(epj float64) (vec [3]float64) {
	var cVec [3]C.double
	C.iauLtpecl(C.double(epj), &cVec[0])
	return v3sC2Go(cVec)
}

//  GoLtpecl Long-term precession of the ecliptic.
func GoLtpecl(epj float64) (vec [3]float64) {

	// Obliquity at J2000.0 (radians).
	const eps0 = 84381.406 * DAS2R

	// Polynomial coefficients
	const NPOL = 4
	var pqpol = [2][NPOL]float64{
		{5851.607687,
			-0.1189000,
			-0.00028913,
			0.000000101},
		{-1600.886300,
			1.1689818,
			-0.00000020,
			-0.000000437},
	}

	// Periodic coefficients
	var pqper = [...][5]float64{
		{708.15, -5486.751211, -684.661560, 667.666730, -5523.863691},
		{2309.00, -17.127623, 2446.283880, -2354.886252, -549.747450},
		{1620.00, -617.517403, 399.671049, -428.152441, -310.998056},
		{492.20, 413.442940, -356.652376, 376.202861, 421.535876},
		{1183.00, 78.614193, -186.387003, 184.778874, -36.776172},
		{622.00, -180.732815, -316.800070, 335.321713, -145.278396},
		{882.00, -87.676083, 198.296701, -185.138669, -34.744450},
		{547.00, 46.140315, 101.135679, -120.972830, 22.885731},
	}
	const NPER = len(pqper)

	// Miscellaneous
	var i int
	var t, p, q, w, a, s, c float64

	// Centuries since J2000.
	t = (epj - 2000.0) / 100.0

	// Initialize P_A and Q_A accumulators.
	p = 0.0
	q = 0.0

	// Periodic terms.
	w = D2PI * t
	for i = 0; i < NPER; i++ {
		a = w / pqper[i][0]
		s = math.Sin(a)
		c = math.Cos(a)
		p += c*pqper[i][1] + s*pqper[i][3]
		q += c*pqper[i][2] + s*pqper[i][4]
	}

	// Polynomial terms.
	w = 1.0
	for i = 0; i < NPOL; i++ {
		p += pqpol[0][i] * w
		q += pqpol[1][i] * w
		w *= t
	}

	// P_A and Q_A (radians).
	p *= DAS2R
	q *= DAS2R

	// Form the ecliptic pole vector.
	w = 1.0 - p*p - q*q
	if w >= 0.0 {
		w = math.Sqrt(w)
	}
	s = math.Sin(eps0)
	c = math.Cos(eps0)
	vec[0] = p
	vec[1] = -q*c - w*s
	vec[2] = -q*s + w*c
	return
}
