package sofa

// #include "sofa.h"
import "C"
import "github.com/8i8/sofa/en"

//            int         status:  0 = OK
//                                -1 = illegal identifier (Note 3)
//                                -2 = internal error (Note 3)
var errGc2gd = en.New(2, "Gc2gd", []string{
	"internal error (Note 3)",
	"illegal identifier (Note 3)",
	"",
})

//  CgoGc2gd Transform geocentric coordinates to geodetic using the
//  specified reference ellipsoid.
//
//  - - - - - -
//   G c 2 g d
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  canonical transformation.
//
//  Given:
//     n       int        ellipsoid identifier (Note 1)
//     xyz     double[3]  geocentric vector (Note 2)
//
//  Returned:
//     elong   double     longitude (radians, east +ve, Note 3)
//     phi     double     latitude (geodetic, radians, Note 3)
//     height  double     height above ellipsoid (geodetic, Notes 2,3)
//
//  Returned (function value):
//            int         status:  0 = OK
//                                -1 = illegal identifier (Note 3)
//                                -2 = internal error (Note 3)
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
//  2) The geocentric vector (xyz, given) and height (height, returned)
//     are in meters.
//
//  3) An error status -1 means that the identifier n is illegal.  An
//     error status -2 is theoretically impossible.  In all error cases,
//     all three results are set to -1e9.
//
//  4) The inverse transformation is performed in the function iauGd2gc.
//
//  Called:
//     iauEform     Earth reference ellipsoids
//     iauGc2gde    geocentric to geodetic transformation, general
//
//  This revision:  2013 September 1
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoGc2gd Transform geocentric coordinates to geodetic using the
//  specified reference ellipsoid.
func CgoGc2gd(n int, xyz [3]float64) (
	elong, phi, height float64, err en.ErrNum) {
	var cElong, cPhi, cHeight C.double
	cXyz := v3sGo2C(xyz)
	cI := C.iauGc2gd(C.int(n), &cXyz[0], &cElong, &cPhi, &cHeight)
	if int(cI) != 0 {
		err = errGc2gd.Set(int(cI))
	}
	return float64(cElong), float64(cPhi), float64(cHeight), err
}

//  GoGc2gd Transform geocentric coordinates to geodetic using the
//  specified reference ellipsoid.
func GoGc2gd(n int, xyz [3]float64) (
	elong, phi, height float64, err en.ErrNum) {

	var a, f float64

	// Obtain reference ellipsoid parameters.
	a, f, err = GoEform(n)

	// If OK, transform x,y,z to longitude, geodetic latitude, height.
	if err == nil {
		elong, phi, height, err = GoGc2gde(a, f, xyz)
		if err != nil && err.Is() < 0 {
			err = errGc2gd.Wrap(err)
			err = errGc2gd.Add(err, -2)
		}
	}

	// Deal with any errors.
	if err != nil && err.Is() < 0 {
		elong = -1e9
		phi = -1e9
		height = -1e9
	}

	// Return the status.
	return
}
