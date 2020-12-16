package sofa

// #include "sofa.h"
import "C"

//  CgoFk52h Transform FK5 (J2000.0) star data into the Hipparcos
//  system.
//
//  - - - - - -
//   F k 5 2 h
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given (all FK5, equinox J2000.0, epoch J2000.0):
//     r5      double    RA (radians)
//     d5      double    Dec (radians)
//     dr5     double    proper motion in RA (dRA/dt, rad/Jyear)
//     dd5     double    proper motion in Dec (dDec/dt, rad/Jyear)
//     px5     double    parallax (arcsec)
//     rv5     double    radial velocity (km/s, positive = receding)
//
//  Returned (all Hipparcos, epoch J2000.0):
//     rh      double    RA (radians)
//     dh      double    Dec (radians)
//     drh     double    proper motion in RA (dRA/dt, rad/Jyear)
//     ddh     double    proper motion in Dec (dDec/dt, rad/Jyear)
//     pxh     double    parallax (arcsec)
//     rvh     double    radial velocity (km/s, positive = receding)
//
//  Notes:
//
//  1) This function transforms FK5 star positions and proper motions
//     into the system of the Hipparcos catalog.
//
//  2) The proper motions in RA are dRA/dt rather than
//     cos(Dec)*dRA/dt, and are per year rather than per century.
//
//  3) The FK5 to Hipparcos transformation is modeled as a pure
//     rotation and spin;  zonal errors in the FK5 catalog are not
//     taken into account.
//
//  4) See also iauH2fk5, iauFk5hz, iauHfk5z.
//
//  Called:
//     iauStarpv    star catalog data to space motion pv-vector
//     iauFk5hip    FK5 to Hipparcos rotation and spin
//     iauRxp       product of r-matrix and p-vector
//     iauPxp       vector product of two p-vectors
//     iauPpp       p-vector plus p-vector
//     iauPvstar    space motion pv-vector to star catalog data
//
//  Reference:
//
//     F.Mignard & M.Froeschle, Astron.Astrophys., 354, 732-739 (2000).
//
//  This revision:  2017 October 12
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoFk52h Transform FK5 (J2000.0) star data into the Hipparcos
//  system.
func CgoFk52h(r5, d5, dr5, dd5, px5, rv5 float64) (
	rh, dh, drh, ddh, pxh, rvh float64) {
	var cRh, cDh, cDrh, cDdh, cPxh, cRvh C.double
	C.iauFk52h(C.double(r5), C.double(d5), C.double(dr5),
		C.double(dd5), C.double(px5), C.double(rv5),
		&cRh, &cDh, &cDrh, &cDdh, &cPxh, &cRvh)
	return float64(cRh), float64(cDh), float64(cDrh),
		float64(cDdh), float64(cPxh), float64(cRvh)
}

//  GoFk52h Transform FK5 (J2000.0) star data into the Hipparcos
//  system.
func GoFk52h(r5, d5, dr5, dd5, px5, rv5 float64) (
	rh, dh, drh, ddh, pxh, rvh float64) {

	var i int
	var r5h [3][3]float64
	var pv5, pvh [2][3]float64
	var s5h, wxp, vv [3]float64

	// FK5 barycentric position/velocity pv-vector (normalized).
	pv5, _ = GoStarpv(r5, d5, dr5, dd5, px5, rv5)

	// FK5 to Hipparcos orientation matrix and spin vector.
	r5h, s5h = GoFk5hip()

	// Make spin units per day instead of per year.
	for i = 0; i < 3; i++ {
		s5h[i] /= 365.25
	}

	// Orient the FK5 position into the Hipparcos system.
	pvh[0] = GoRxp(r5h, pv5[0])

	// Apply spin to the position giving an extra space motion
	// component.
	wxp = GoPxp(pv5[0], s5h)

	// Add this component to the FK5 space motion.
	vv = GoPpp(wxp, pv5[1])

	// Orient the FK5 space motion into the Hipparcos system.
	pvh[1] = GoRxp(r5h, vv)

	// Hipparcos pv-vector to spherical.
	rh, dh, drh, ddh, pxh, rvh, _ = GoPvstar(pvh)
	return
}
