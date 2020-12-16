package sofa

// #include "sofa.h"
import "C"

//  CgoHfk5z Transform a Hipparcos star position into FK5 J2000.0,
//  assuming zero Hipparcos proper motion.
//
//  - - - - - -
//   H f k 5 z
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     rh            double    Hipparcos RA (radians)
//     dh            double    Hipparcos Dec (radians)
//     date1,date2   double    TDB date (Note 1)
//
//  Returned (all FK5, equinox J2000.0, date date1+date2):
//     r5            double    RA (radians)
//     d5            double    Dec (radians)
//     dr5           double    FK5 RA proper motion (rad/year, Note 4)
//     dd5           double    Dec proper motion (rad/year, Note 4)
//
//  Notes:
//
//  1) The TT date date1+date2 is a Julian Date, apportioned in any
//     convenient way between the two arguments.  For example,
//     JD(TT)=2450123.7 could be expressed in any of these ways,
//     among others:
//
//            date1          date2
//
//         2450123.7           0.0       (JD method)
//         2451545.0       -1421.3       (J2000 method)
//         2400000.5       50123.2       (MJD method)
//         2450123.5           0.2       (date & time method)
//
//     The JD method is the most natural and convenient to use in
//     cases where the loss of several decimal digits of resolution
//     is acceptable.  The J2000 method is best matched to the way
//     the argument is handled internally and will deliver the
//     optimum resolution.  The MJD method and the date & time methods
//     are both good compromises between resolution and convenience.
//
//  2) The proper motion in RA is dRA/dt rather than cos(Dec)*dRA/dt.
//
//  3) The FK5 to Hipparcos transformation is modeled as a pure rotation
//     and spin;  zonal errors in the FK5 catalogue are not taken into
//     account.
//
//  4) It was the intention that Hipparcos should be a close
//     approximation to an inertial frame, so that distant objects have
//     zero proper motion;  such objects have (in general) non-zero
//     proper motion in FK5, and this function returns those fictitious
//     proper motions.
//
//  5) The position returned by this function is in the FK5 J2000.0
//     reference system but at date date1+date2.
//
//  6) See also iauFk52h, iauH2fk5, iauFk5zhz.
//
//  Called:
//     iauS2c       spherical coordinates to unit vector
//     iauFk5hip    FK5 to Hipparcos rotation and spin
//     iauRxp       product of r-matrix and p-vector
//     iauSxp       multiply p-vector by scalar
//     iauRxr       product of two r-matrices
//     iauTrxp      product of transpose of r-matrix and p-vector
//     iauPxp       vector product of two p-vectors
//     iauPv2s      pv-vector to spherical
//     iauAnp       normalize angle into range 0 to 2pi
//
//  Reference:
//
//     F.Mignard & M.Froeschle, 2000, Astron.Astrophys. 354, 732-739.
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoHfk5z Transform a Hipparcos star position into FK5 J2000.0,
//  assuming zero Hipparcos proper motion.
func CgoHfk5z(rh, dh, date1, date2 float64) (
	r5, d5, dr5, dd5 float64) {
	var cR5, cD5, cDr5, cDd5 C.double
	C.iauHfk5z(C.double(rh), C.double(dh), C.double(date1),
		C.double(date2), &cR5, &cD5, &cDr5, &cDd5)
	return float64(cR5), float64(cD5), float64(cDr5), float64(cDd5)
}

//  GoHfk5z Transform a Hipparcos star position into FK5 J2000.0,
//  assuming zero Hipparcos proper motion.
func GoHfk5z(rh, dh, date1, date2 float64) (
	r5, d5, dr5, dd5 float64) {

	var t, w float64
	var ph, s5h, sh, vst, vv [3]float64
	var pv5e [2][3]float64
	var r5h, rst, r5ht [3][3]float64

	// Time interval from fundamental epoch J2000.0 to given date
	// (JY).
	t = ((date1 - DJ00) + date2) / DJY

	// Hipparcos barycentric position vector (normalized).
	ph = GoS2c(rh, dh)

	// FK5 to Hipparcos orientation matrix and spin vector.
	r5h, s5h = GoFk5hip()

	// Rotate the spin into the Hipparcos system.
	sh = GoRxp(r5h, s5h)

	// Accumulated Hipparcos wrt FK5 spin over that interval.
	vst = GoSxp(t, s5h)

	// Express the accumulated spin as a rotation matrix.
	rst = GoRv2m(vst)

	// Rotation matrix:  accumulated spin, then FK5 to Hipparcos.
	r5ht = GoRxr(r5h, rst)

	// De-orient & de-spin the Hipparcos position into FK5 J2000.0.
	pv5e[0] = GoTrxp(r5ht, ph)

	// Apply spin to the position giving a space motion.
	vv = GoPxp(sh, ph)

	// De-orient & de-spin the Hipparcos space motion into FK5
	// J2000.0.
	pv5e[1] = GoTrxp(r5ht, vv)

	// FK5 position/velocity pv-vector to spherical.
	w, d5, _, dr5, dd5, _ = GoPv2s(pv5e)
	r5 = GoAnp(w)
	return
}
