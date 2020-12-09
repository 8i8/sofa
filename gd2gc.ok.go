package sofa

// #include "sofa.h"
import "C"
import "errors"

var errGd2gcE1 = errors.New("illegal identifier (gd2c documentation note 3)")
var errGd2gcE2 = errors.New("illegal case (gd2c documentation note 3)")

//  CgoGd2gc Transform geodetic coordinates to geocentric using the
//  specified reference ellipsoid.
//
//  - - - - - -
//   G d 2 g c
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  canonical transformation.
//
//  Given:
//     n       int        ellipsoid identifier (Note 1)
//     elong   double     longitude (radians, east +ve)
//     phi     double     latitude (geodetic, radians, Note 3)
//     height  double     height above ellipsoid (geodetic, Notes 2,3)
//
//  Returned:
//     xyz     double[3]  geocentric vector (Note 2)
//
//  Returned (function value):
//     err     error      nil = OK
//                        errGd2gcE1 = illegal identifier (Note 3)
//                        errGd2gcE2 = illegal case (Note 3)
//
//  Notes:
//
//  1) The identifier n is a number that specifies the choice of
//     reference ellipsoid.  The following are supported:
//
//        n    ellipsoid
//
//        1     WGS84
//        2     GRS80
//        3     WGS72
//
//     The n value has no significance outside the SOFA software.  For
//     convenience, symbols WGS84 etc. are defined in sofam.h.
//
//  2) The height (height, given) and the geocentric vector (xyz,
//     returned) are in meters.
//
//  3) No validation is performed on the arguments elong, phi and
//     height.  An error status -1 means that the identifier n is
//     illegal.  An error status -2 protects against cases that would
//     lead to arithmetic exceptions.  In all error cases, xyz is set
//     to zeros.
//
//  4) The inverse transformation is performed in the function iauGc2gd.
//
//  Called:
//     iauEform     Earth reference ellipsoids
//     iauGd2gce    geodetic to geocentric transformation, general
//     iauZp        zero p-vector
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoGd2gc Transform geodetic coordinates to geocentric using the
//  specified reference ellipsoid.
func CgoGd2gc(n int, elong, phi,
	height float64) (xyz [3]float64, err error) {

	var cXyz [3]C.double
	cI := C.iauGd2gc(C.int(n), C.double(elong), C.double(phi),
		C.double(height), &cXyz[0])
	switch int(cI) {
	case 0:
	case -1:
		err = errGd2gcE1
	case -2:
		err = errGd2gcE2
	default:
		err = errAdmin
	}
	return v3sC2Go(cXyz), err
}

func GoGd2gc(n int, elong, phi,
	height float64) (xyz [3]float64, err error) {

	var a, f float64

	// Obtain reference ellipsoid parameters.
	a, f, err = GoEform(n)
	if err != nil {
		err = errGd2gcE1
	}

	// If OK, transform longitude, geodetic latitude, height to x,y,z.
	if err == nil {
		xyz, err = GoGd2gce(a, f, elong, phi, height)
		if err != nil {
			err = errGd2gcE2
		}
	}

	// Deal with any errors.
	if err != nil {
		xyz = [3]float64{}
	}

	return
}
