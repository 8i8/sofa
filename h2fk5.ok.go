package sofa

// #include "sofa.h"
import "C"

//  CgoH2fk5 Transform Hipparcos star data into the FK5 (J2000.0)
//  system.
//
//  - - - - - -
//   H 2 f k 5
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given (all Hipparcos, epoch J2000.0):
//     rh      double    RA (radians)
//     dh      double    Dec (radians)
//     drh     double    proper motion in RA (dRA/dt, rad/Jyear)
//     ddh     double    proper motion in Dec (dDec/dt, rad/Jyear)
//     pxh     double    parallax (arcsec)
//     rvh     double    radial velocity (km/s, positive = receding)
//
//  Returned (all FK5, equinox J2000.0, epoch J2000.0):
//     r5      double    RA (radians)
//     d5      double    Dec (radians)
//     dr5     double    proper motion in RA (dRA/dt, rad/Jyear)
//     dd5     double    proper motion in Dec (dDec/dt, rad/Jyear)
//     px5     double    parallax (arcsec)
//     rv5     double    radial velocity (km/s, positive = receding)
//
//  Notes:
//
//  1) This function transforms Hipparcos star positions and proper
//     motions into FK5 J2000.0.
//
//  2) The proper motions in RA are dRA/dt rather than
//     cos(Dec)*dRA/dt, and are per year rather than per century.
//
//  3) The FK5 to Hipparcos transformation is modeled as a pure
//     rotation and spin;  zonal errors in the FK5 catalog are not
//     taken into account.
//
//  4) See also iauFk52h, iauFk5hz, iauHfk5z.
//
//  Called:
//     iauStarpv    star catalog data to space motion pv-vector
//     iauFk5hip    FK5 to Hipparcos rotation and spin
//     iauRv2m      r-vector to r-matrix
//     iauRxp       product of r-matrix and p-vector
//     iauTrxp      product of transpose of r-matrix and p-vector
//     iauPxp       vector product of two p-vectors
//     iauPmp       p-vector minus p-vector
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
//  CgoH2fk5 Transform Hipparcos star data into the FK5 (J2000.0)
//  system.
func CgoH2fk5(rh, dh, drh, ddh, pxh, rvh float64) (
	r5, d5, dr5, dd5, px5, rv5 float64) {
	var cR5, cD5, cDr5, cDd5, cPx5, cRv5 C.double
	C.iauH2fk5(C.double(rh), C.double(dh), C.double(drh),
		C.double(ddh), C.double(pxh), C.double(rvh),
		&cR5, &cD5, &cDr5, &cDd5, &cPx5, &cRv5)
	return float64(cR5), float64(cD5), float64(cDr5), float64(cDd5),
		float64(cPx5), float64(cRv5)
}

//  GoH2fk5 Transform Hipparcos star data into the FK5 (J2000.0)
//  system.
func GoH2fk5(rh, dh, drh, ddh, pxh, rvh float64) (
	r5, d5, dr5, dd5, px5, rv5 float64) {

	var i int

	var s5h, sh, wxp, vv [3]float64
	var r5h [3][3]float64
	var pvh, pv5 [2][3]float64

	// Hipparcos barycentric position/velocity pv-vector
	// (normalized).
	pvh, _ = GoStarpv(rh, dh, drh, ddh, pxh, rvh)

	// FK5 to Hipparcos orientation matrix and spin vector.
	r5h, s5h = GoFk5hip()

	// Make spin units per day instead of per year.
	for i = 0; i < 3; i++ {
		s5h[i] /= 365.25
	}

	// Orient the spin into the Hipparcos system.
	sh = GoRxp(r5h, s5h)

	// De-orient the Hipparcos position into the FK5 system.
	pv5[0] = GoTrxp(r5h, pvh[0])

	// Apply spin to the position giving an extra space motion
	// component.
	wxp = GoPxp(pvh[0], sh)

	// Subtract this component from the Hipparcos space motion.
	vv = GoPmp(pvh[1], wxp)

	// De-orient the Hipparcos space motion into the FK5 system.
	pv5[1] = GoTrxp(r5h, vv)

	// FK5 pv-vector to spherical.
	r5, d5, dr5, dd5, px5, rv5, _ = GoPvstar(pv5)
	return
}
