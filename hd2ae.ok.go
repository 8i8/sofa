package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoHd2ae Equatorial to horizon coordinates:  transform hour angle
//  and declination to azimuth and altitude.
//
//  - - - - - -
//   H d 2 a e
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     ha       double       hour angle (local)
//     dec      double       declination
//     phi      double       site latitude
//
//  Returned:
//     *az      double       azimuth
//     *el      double       altitude (informally, elevation)
//
//  Notes:
//
//  1)  All the arguments are angles in radians.
//
//  2)  Azimuth is returned in the range 0-2pi;  north is zero, and east
//      is +pi/2.  Altitude is returned in the range +/- pi/2.
//
//  3)  The latitude phi is pi/2 minus the angle between the Earth's
//      rotation axis and the adopted zenith.  In many applications it
//      will be sufficient to use the published geodetic latitude of the
//      site.  In very precise (sub-arcsecond) applications, phi can be
//      corrected for polar motion.
//
//  4)  The returned azimuth az is with respect to the rotational north
//      pole, as opposed to the ITRS pole, and for sub-arcsecond
//      accuracy will need to be adjusted for polar motion if it is to
//      be with respect to north on a map of the Earth's surface.
//
//  5)  Should the user wish to work with respect to the astronomical
//      zenith rather than the geodetic zenith, phi will need to be
//      adjusted for deflection of the vertical (often tens of
//      arcseconds), and the zero point of the hour angle ha will also
//      be affected.
//
//  6)  The transformation is the same as Vh = Rz(pi)*Ry(pi/2-phi)*Ve,
//      where Vh and Ve are lefthanded unit vectors in the (az,el) and
//      (ha,dec) systems respectively and Ry and Rz are rotations about
//      first the y-axis and then the z-axis.  (n.b. Rz(pi) simply
//      reverses the signs of the x and y components.)  For efficiency,
//      the algorithm is written out rather than calling other utility
//      functions.  For applications that require even greater
//      efficiency, additional savings are possible if constant terms
//      such as functions of latitude are computed once and for all.
//
//  7)  Again for efficiency, no range checking of arguments is carried
//      out.
//
//  Last revision:   2017 September 12
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoHd2ae Equatorial to horizon coordinates:  transform hour angle
//  and declination to azimuth and altitude.
func CgoHd2ae(ha, dec, phi float64) (az, el float64) {
	var cAz, cEl C.double
	C.iauHd2ae(C.double(ha), C.double(dec), C.double(phi),
		&cAz, &cEl)
	return float64(cAz), float64(cEl)
}

//  GoHd2ae Equatorial to horizon coordinates:  transform hour angle
//  and declination to azimuth and altitude.
func GoHd2ae(ha, dec, phi float64) (az, el float64) {

	var sh, ch, sd, cd, sp, cp, x, y, z, r, a float64

	// Useful trig functions.
	sh = math.Sin(ha)
	ch = math.Cos(ha)
	sd = math.Sin(dec)
	cd = math.Cos(dec)
	sp = math.Sin(phi)
	cp = math.Cos(phi)

	// Az,Alt unit vector.
	x = -ch*cd*sp + sd*cp
	y = -sh * cd
	z = ch*cd*cp + sd*sp

	// To spherical.
	r = math.Sqrt(x*x + y*y)
	if r != 0.0 {
		a = math.Atan2(y, x)
	}
	if a < 0.0 {
		az = a + D2PI
	} else {
		az = a
	}
	el = math.Atan2(z, r)
	return
}
