package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoFae03 Mean longitude of Earth (IERS Conventions 2003).
//
//  - - - - - -
//   F a e 0 3
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
//           double    mean longitude of Earth, radians (Note 2)
//
//  Notes:
//
//  1) Though t is strictly TDB, it is usually more convenient to use
//     TT, which makes no significant difference.
//
//  2) The expression used is as adopted in IERS Conventions (2003) and
//     comes from Souchay et al. (1999) after Simon et al. (1994).
//
//  References:
//
//     McCarthy, D. D., Petit, G. (eds.), IERS Conventions (2003),
//     IERS Technical Note No. 32, BKG (2004)
//
//     Simon, J.-L., Bretagnon, P., Chapront, J., Chapront-Touze, M.,
//     Francou, G., Laskar, J. 1994, Astron.Astrophys. 282, 663-683
//
//     Souchay, J., Loysel, B., Kinoshita, H., Folgueira, M. 1999,
//     Astron.Astrophys.Supp.Ser. 135, 111
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoFae03 Mean longitude of Earth (IERS Conventions 2003).
func CgoFae03(t float64) float64 {
	var cF C.double
	cF = C.iauFae03(C.double(t))
	return float64(cF)
}

// GoFae03 Fundamental argument, IERS Conventions (2003): mean longitude
// of Earth.
func GoFae03(t float64) float64 {
	// Mean longitude of Earth (IERS Conventions 2003).
	return math.Mod(1.753470314+628.3075849991*t, D2PI)
}
