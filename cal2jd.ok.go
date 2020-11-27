package sofa

// #include <sofa.h>
import "C"

//  Cal2jd returns the MJD, Modified Julian Date of the given date.
//
//  - - - - - - - - - -
//   i a u C a l 2 j d
//  - - - - - - - - - -
//
//  Gregorian Calendar to Julian Date.
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     iy,im,id  int     year, month, day in Gregorian calendar (Note 1)
//
//  Returned:
//     djm0      double  MJD zero-point: always 2400000.5
//     djm       double  Modified Julian Date for 0 hrs
//
//  Returned (function value):
//               int     status:
//                           0 = OK
//                          -1 = bad year   (Note 3: JD not computed)
//                          -2 = bad month  (JD not computed)
//                          -3 = bad day    (JD computed)
//
//  Notes:
//
//  1) The algorithm used is valid from -4800 March 1, but this
//     implementation rejects dates before -4799 January 1.
//
//  2) The Julian Date is returned in two pieces, in the usual SOFA
//     manner, which is designed to preserve time resolution.  The
//     Julian Date is available as a single number by adding djm0 and
//     djm.
//
//  3) In early eras the conversion is from the "Proleptic Gregorian
//     Calendar";  no account is taken of the date(s) of adoption of
//     the Gregorian Calendar, nor is the AD/BC numbering convention
//     observed.
//
//  Reference:
//
//     Explanatory Supplement to the Astronomical Almanac,
//     P. Kenneth Seidelmann (ed), University Science Books (1992),
//     Section 12.92 (p604).
//
//  This revision:  2013 August 7
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
func Cal2jd(iy, im, id int) (djm0, djm float64, err error) {
	var d, e C.double
	i := C.iauCal2jd(C.int(iy), C.int(im), C.int(id), &d, &e)
	djm0 = float64(d)
	djm = float64(e)
	switch i {
	case -1:
		err = ErrYear
	case -2:
		err = ErrMonth
	case -3:
		err = ErrDay
	}
	return
}
