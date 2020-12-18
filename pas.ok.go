package sofa

// #include "sofa.h"
import "C"
import "math"

//  CgoPas Position-angle from spherical coordinates.
//
//  - - - -
//   P a s
//  - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     al     double     longitude of point A (e.g. RA) in radians
//     ap     double     latitude of point A (e.g. Dec) in radians
//     bl     double     longitude of point B
//     bp     double     latitude of point B
//
//  Returned (function value):
//            double     position angle of B with respect to A
//
//  Notes:
//
//  1) The result is the bearing (position angle), in radians, of point
//     B with respect to point A.  It is in the range -pi to +pi.  The
//     sense is such that if B is a small distance "east" of point A,
//     the bearing is approximately +pi/2.
//
//  2) Zero is returned if the two points are coincident.
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPas Position-angle from spherical coordinates.
func CgoPas(al, ap, bl, bp float64) (pa float64) {
	var cPa C.double
	cPa = C.iauPas(C.double(al), C.double(ap), C.double(bl), C.double(bp))
	return float64(cPa)
}

//  GoPas Position-angle from spherical coordinates.
func GoPas(al, ap, bl, bp float64) (pa float64) {

	var dl, x, y float64

	dl = bl - al
	y = math.Sin(dl) * math.Cos(bp)
	x = math.Sin(bp)*math.Cos(ap) - math.Cos(bp)*math.Sin(ap)*math.Cos(dl)
	if x != 0.0 || y != 0.0 {
		pa = math.Atan2(y, x)
	}
	return
}
