package sofa

// #include "sofa.h"
import "C"

//  CgoSeps Angular separation between two sets of spherical
//  coordinates.
//
//  - - - - -
//   S e p s
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     al     double       first longitude (radians)
//     ap     double       first latitude (radians)
//     bl     double       second longitude (radians)
//     bp     double       second latitude (radians)
//
//  Returned (function value):
//            double       angular separation (radians)
//
//  Called:
//     iauS2c       spherical coordinates to unit vector
//     iauSepp      angular separation between two p-vectors
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoSeps Angular separation between two sets of spherical
//  coordinates.
func CgoSeps(al, ap, bl, bp float64) float64 {
	var cF C.double
	cF = C.iauSeps(C.double(al), C.double(ap), C.double(bl),
	C.double(bp))
	return float64(cF)
}

//  GoSeps Angular separation between two sets of spherical
//  coordinates.
func GoSeps(al, ap, bl, bp float64) float64 {

   var ac, bc[3]float64
   var s float64

// Spherical to Cartesian. 
   ac = GoS2c(al, ap);
   bc = GoS2c(bl, bp);

// Angle between the vectors. 
   s = GoSepp(ac, bc);

   return s;

   }
