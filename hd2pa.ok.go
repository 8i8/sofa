package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoHd2pa Parallactic angle for a given hour angle and declination.
//
//  - - - - - -
//   H d 2 p a
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     ha     double     hour angle
//     dec    double     declination
//     phi    double     site latitude
//
//  Returned (function value):
//            double     parallactic angle
//
//  Notes:
//
//  1)  All the arguments are angles in radians.
//
//  2)  The parallactic angle at a point in the sky is the position
//      angle of the vertical, i.e. the angle between the directions to
//      the north celestial pole and to the zenith respectively.
//
//  3)  The result is returned in the range -pi to +pi.
//
//  4)  At the pole itself a zero result is returned.
//
//  5)  The latitude phi is pi/2 minus the angle between the Earth's
//      rotation axis and the adopted zenith.  In many applications it
//      will be sufficient to use the published geodetic latitude of the
//      site.  In very precise (sub-arcsecond) applications, phi can be
//      corrected for polar motion.
//
//  6)  Should the user wish to work with respect to the astronomical
//      zenith rather than the geodetic zenith, phi will need to be
//      adjusted for deflection of the vertical (often tens of
//      arcseconds), and the zero point of the hour angle ha will also
//      be affected.
//
//  Reference:
//     Smart, W.M., "Spherical Astronomy", Cambridge University Press,
//     6th edition (Green, 1977), p49.
//
//  Last revision:   2017 September 12
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoHd2pa Parallactic angle for a given hour angle and declination.
func CgoHd2pa(ha, dec, phi float64) (pa float64) {
	var cPa C.double
	cPa = C.iauHd2pa(C.double(ha), C.double(dec), C.double(phi))
	return float64(cPa)
}

//  GoHd2pa Parallactic angle for a given hour angle and declination.
func GoHd2pa(ha, dec, phi float64) (pa float64) {
	var cp, cqsz, sqsz float64

	cp = math.Cos(phi)
	sqsz = cp * math.Sin(ha)
	cqsz = math.Sin(phi)*math.Cos(dec) -
		cp*math.Sin(dec)*math.Cos(ha)
	if sqsz != 0.0 || cqsz != 0.0 {
		pa = math.Atan2(sqsz, cqsz)
	}
	return
}
