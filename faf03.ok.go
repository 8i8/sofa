package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoFaf03 Mean longitude of the Moon minus that of the ascending node
//  (IERS Conventions 2003).
//
//  - - - - - -
//   F a f 0 3
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  canonical model.
//
//  Given:
//     t     double    TDB, Julian centuries since J2000.0 (Note 1)
//
//  Returned (function value):
//           double    F, radians (Note 2)
//
//  Notes:
//
//  1) Though t is strictly TDB, it is usually more convenient to use
//     TT, which makes no significant difference.
//
//  2) The expression used is as adopted in IERS Conventions (2003) and
//     is from Simon et al. (1994).
//
//  References:
//
//     McCarthy, D. D., Petit, G. (eds.), IERS Conventions (2003),
//     IERS Technical Note No. 32, BKG (2004)
//
//     Simon, J.-L., Bretagnon, P., Chapront, J., Chapront-Touze, M.,
//     Francou, G., Laskar, J. 1994, Astron.Astrophys. 282, 663-683
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  Mean longitude of the Moon minus that of the ascending node (IERS
//  Conventions 2003).
func CgoFaf03(t float64) float64 {
	var cF C.double
	cF = C.iauFaf03(C.double(t))
	return float64(cF)
}

//  GoFaf03 mean longitude of the Moon minus mean longitude of the
//  ascending node.  Fundamental argument, IERS Conventions (2003)
func GoFaf03(t float64) float64 {
	// Mean longitude of the Moon minus that of the ascending node
	// (IERS Conventions 2003).
	return math.Mod(335779.526232+
		t*(1739527262.8478+
			t*(-12.7512+
				t*(-0.001037+
					t*(0.00000417)))), TURNAS) * DAS2R
}
