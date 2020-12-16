package sofa

// #include "sofa.h"
import "C"

//  CgoLtp Long-term precession matrix.
//
//  - - - -
//   L t p
//  - - - -
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
//     rp      double[3][3]   precession matrix, J2000.0 to date
//
//  Notes:
//
//  1) The matrix is in the sense
//
//        P_date = rp x P_J2000,
//
//     where P_J2000 is a vector with respect to the J2000.0 mean
//     equator and equinox and P_date is the same vector with respect to
//     the equator and equinox of epoch epj.
//
//  2) The Vondrak et al. (2011, 2012) 400 millennia precession model
//     agrees with the IAU 2006 precession at J2000.0 and stays within
//     100 microarcseconds during the 20th and 21st centuries.  It is
//     accurate to a few arcseconds throughout the historical period,
//     worsening to a few tenths of a degree at the end of the
//     +/- 200,000 year time span.
//
//  Called:
//     iauLtpequ    equator pole, long term
//     iauLtpecl    ecliptic pole, long term
//     iauPxp       vector product
//     iauPn        normalize vector
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
//  This revision:  2015 December 6
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoLtp Long-term precession matrix.
func CgoLtp(epj float64) (rp [3][3]float64) {
	var cRp [3][3]C.double
	C.iauLtp(C.double(epj), &cRp[0])
	return v3tC2Go(cRp)
}

//  GoLtp Long-term precession matrix.
func GoLtp(epj float64) (rp [3][3]float64) {
	var i int
	var peqr, pecl, v, eqx [3]float64

	// Equator pole (bottom row of matrix).
	peqr = GoLtpequ(epj)

	// Ecliptic pole.
	pecl = GoLtpecl(epj)

	// Equinox (top row of matrix).
	v = GoPxp(peqr, pecl)
	_, eqx = GoPn(v)

	// Middle row of matrix.
	v = GoPxp(peqr, eqx)

	// Assemble the matrix.
	for i = 0; i < 3; i++ {
		rp[0][i] = eqx[i]
		rp[1][i] = v[i]
		rp[2][i] = peqr[i]
	}
	return
}
