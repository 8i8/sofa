package sofa

// #include "sofa.h"
import "C"

//  CgoA2tf Decompose radians into hours, minutes, seconds, fraction.
//
//  - - - - -
//   A 2 t f
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     ndp     int     resolution (Note 1)
//     angle   double  angle in radians
//
//  Returned:
//     sign    char*   '+' or '-'
//     ihmsf   int[4]  hours, minutes, seconds, fraction
//
//  Notes:
//
//  1) The argument ndp is interpreted as follows:
//
//     ndp         resolution
//      :      ...0000 00 00
//     -7         1000 00 00
//     -6          100 00 00
//     -5           10 00 00
//     -4            1 00 00
//     -3            0 10 00
//     -2            0 01 00
//     -1            0 00 10
//      0            0 00 01
//      1            0 00 00.1
//      2            0 00 00.01
//      3            0 00 00.001
//      :            0 00 00.000...
//
//  2) The largest positive useful value for ndp is determined by the
//     size of angle, the format of doubles on the target platform, and
//     the risk of overflowing ihmsf[3].  On a typical platform, for
//     angle up to 2pi, the available floating-point precision might
//     correspond to ndp=12.  However, the practical limit is typically
//     ndp=9, set by the capacity of a 32-bit int, or ndp=4 if int is
//     only 16 bits.
//
//  3) The absolute value of angle may exceed 2pi.  In cases where it
//     does not, it is up to the caller to test for and handle the
//     case where angle is very nearly 2pi and rounds up to 24 hours,
//     by testing for ihmsf[0]=24 and setting ihmsf[0-3] to zero.
//
//  Called:
//     iauD2tf      decompose days to hms
//
//  This revision:  2020 April 1
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoA2tf Decompose radians into hours, minutes, seconds, fraction.
func CgoA2tf(ndp int, angle float64) (sign byte, ihmsf [4]int) {
	var cSign C.char
	var cIhmsf [4]C.int
	C.iauA2tf(C.int(ndp), C.double(angle), &cSign, &cIhmsf[0])
	return byte(cSign), v4sIntC2Go(cIhmsf)
}

// GoA2tf Decompose radians into hours, minutes, seconds, fraction.
func GoA2tf(ndp int, angle float64) (sign byte, ihmsf [4]int) {
	// Scale then use days to h,m,s function.
	sign, ihmsf = GoD2tf(ndp, angle/D2PI)
	return
}
