package sofa

// #include "sofa.h"
import "C"
import (
	"math"

	"github.com/8i8/sofa/en"
)

var errGc2gde = en.New(2, "Gc2gde", []string{
	"illegal a",
	"illegal f",
	"",
})

//  CgoGc2gde Transform geocentric coordinates to geodetic for a
//  reference ellipsoid of specified form.
//
//  - - - - - - -
//   G c 2 g d e
//  - - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     a       double     equatorial radius (Notes 2,4)
//     f       double     flattening (Note 3)
//     xyz     double[3]  geocentric vector (Note 4)
//
//  Returned:
//     elong   double     longitude (radians, east +ve)
//     phi     double     latitude (geodetic, radians)
//     height  double     height above ellipsoid (geodetic, Note 4)
//
//  Returned (function value):
//             int        status:  0 = OK
//                                -1 = illegal f
//                                -2 = illegal a
//
//  Notes:
//
//  1) This function is based on the GCONV2H Fortran subroutine by
//     Toshio Fukushima (see reference).
//
//  2) The equatorial radius, a, can be in any units, but meters is
//     the conventional choice.
//
//  3) The flattening, f, is (for the Earth) a value around 0.00335,
//     i.e. around 1/298.
//
//  4) The equatorial radius, a, and the geocentric vector, xyz,
//     must be given in the same units, and determine the units of
//     the returned height, height.
//
//  5) If an error occurs (status < 0), elong, phi and height are
//     unchanged.
//
//  6) The inverse transformation is performed in the function
//     iauGd2gce.
//
//  7) The transformation for a standard ellipsoid (such as WGS84) can
//     more conveniently be performed by calling iauGc2gd, which uses a
//     numerical code to identify the required A and F values.
//
//  Reference:
//
//     Fukushima, T., "Transformation from Cartesian to geodetic
//     coordinates accelerated by Halley's method", J.Geodesy (2006)
//     79: 689-693
//
//  This revision:  2014 November 7
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoGc2gde Transform geocentric coordinates to geodetic for a
//  reference ellipsoid of specified form.
func CgoGc2gde(a, f float64, xyz [3]float64) (
	elong, phi, height float64, err en.ErrNum) {

	var cElong, cPhi, cHeight C.double
	cXyz := v3sGo2C(xyz)
	cI := C.iauGc2gde(C.double(a), C.double(f), &cXyz[0], &cElong, &cPhi,
		&cHeight)
	if n := int(cI); n != 0 {
		err = errGc2gde.Set(n)
	}
	return float64(cElong), float64(cPhi), float64(cHeight), err
}

//  GoGc2gde Transform geocentric coordinates to geodetic for a
//  reference ellipsoid of specified form.
func GoGc2gde(a, f float64, xyz [3]float64) (
	elong, phi, height float64, err en.ErrNum) {

	var aeps2, e2, e4t, ec2, ec, b, x, y, z, p2, absz, p, s0, pn,
		zc, c0, c02, c03, s02, s03, a02, a0, a03, d0, f0, b0,
		s1, cc, s12, cc2 float64

	// -------------
	// Preliminaries
	// -------------

	// Validate ellipsoid parameters.
	if f < 0.0 || f >= 1.0 {
		err = errGc2gde.Set(-1)
		return
	}
	if a <= 0.0 {
		err = errGc2gde.Set(-2)
		return
	}

	// Functions of ellipsoid parameters (with further validation of
	// f).
	aeps2 = a * a * 1e-32
	e2 = (2.0 - f) * f
	e4t = e2 * e2 * 1.5
	ec2 = 1.0 - e2
	if ec2 <= 0.0 {
		err = errGc2gde.Set(-1)
		return
	}
	ec = math.Sqrt(ec2)
	b = a * ec

	// Cartesian components.
	x = xyz[0]
	y = xyz[1]
	z = xyz[2]

	// Distance from polar axis squared.
	p2 = x*x + y*y

	// Longitude.
	if p2 > 0.0 {
		elong = math.Atan2(y, x)
	}

	// Unsigned z-coordinate.
	absz = math.Abs(z)

	// Proceed unless polar case.
	if p2 > aeps2 {

		// Distance from polar axis.
		p = math.Sqrt(p2)

		// Normalization.
		s0 = absz / a
		pn = p / a
		zc = ec * s0

		// Prepare Newton correction factors.
		c0 = ec * pn
		c02 = c0 * c0
		c03 = c02 * c0
		s02 = s0 * s0
		s03 = s02 * s0
		a02 = c02 + s02
		a0 = math.Sqrt(a02)
		a03 = a02 * a0
		d0 = zc*a03 + e2*s03
		f0 = pn*a03 - e2*c03

		// Prepare Halley correction factor.
		b0 = e4t * s02 * c02 * pn * (a0 - ec)
		s1 = d0*f0 - b0*s0
		cc = ec * (f0*f0 - b0*c0)

		// Evaluate latitude and height.
		phi = math.Atan(s1 / cc)
		s12 = s1 * s1
		cc2 = cc * cc
		height = (p*cc + absz*s1 - a*math.Sqrt(ec2*s12+cc2)) /
			math.Sqrt(s12+cc2)
	} else {

		// Exception: pole.
		phi = DPI / 2.0
		height = absz - b
	}

	// Restore sign of latitude.
	if z < 0 {
		phi = -phi
	}

	// OK status.
	return
}
