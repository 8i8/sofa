package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoFad03 Mean elongation of the Moon from the Sun (IERS Conventions
//  2003).
//
//  - - - - - -
//   F a d 0 3
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
//           double    D, radians (Note 2)
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
//  CgoFad03 Mean elongation of the Moon from the Sun (IERS Conventions
//  2003).
func CgoFad03(t float64) float64 {
	var cF C.double
	cF = C.iauFad03(C.double(t))
	return float64(cF)
}

// GoFad03 Fundamental argument, IERS Conventions (2003): mean elongation
// of the Moon from the Sun.
func GoFad03(t float64) float64 {
	// Mean elongation of the Moon from the Sun (IERS Conventions 2003).
	return math.Mod(1072260.703692+
		t*(1602961601.2090+
			t*(-6.3706+
				t*(0.006593+
					t*(-0.00003169)))), TURNAS) * DAS2R
}
