package sofa

// #include "sofa.h"
import "C"

//  CgoFw2m Form rotation matrix given the Fukushima-Williams angles.
//
//  - - - - -
//   F w 2 m
//  - - - - -
//
//  Form rotation matrix given the Fukushima-Williams angles.
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     gamb     double         F-W angle gamma_bar (radians)
//     phib     double         F-W angle phi_bar (radians)
//     psi      double         F-W angle psi (radians)
//     eps      double         F-W angle epsilon (radians)
//
//  Returned:
//     r        double[3][3]   rotation matrix
//
//  Notes:
//
//  1) Naming the following points:
//
//           e = J2000.0 ecliptic pole,
//           p = GCRS pole,
//           E = ecliptic pole of date,
//     and   P = CIP,
//
//     the four Fukushima-Williams angles are as follows:
//
//        gamb = gamma = epE
//        phib = phi = pE
//        psi = psi = pEP
//        eps = epsilon = EP
//
//  2) The matrix representing the combined effects of frame bias,
//     precession and nutation is:
//
//        NxPxB = R_1(-eps).R_3(-psi).R_1(phib).R_3(gamb)
//
//  3) The present function can construct three different matrices,
//     depending on which angles are supplied as the arguments gamb,
//     phib, psi and eps:
//
//     o  To obtain the nutation x precession x frame bias matrix,
//        first generate the four precession angles known conventionally
//        as gamma_bar, phi_bar, psi_bar and epsilon_A, then generate
//        the nutation components Dpsi and Depsilon and add them to
//        psi_bar and epsilon_A, and finally call the present function
//        using those four angles as arguments.
//
//     o  To obtain the precession x frame bias matrix, generate the
//        four precession angles and call the present function.
//
//     o  To obtain the frame bias matrix, generate the four precession
//        angles for date J2000.0 and call the present function.
//
//     The nutation-only and precession-only matrices can if necessary
//     be obtained by combining these three appropriately.
//
//  Called:
//     iauIr        initialize r-matrix to identity
//     iauRz        rotate around Z-axis
//     iauRx        rotate around X-axis
//
//  References:
//
//     Capitaine, N. & Wallace, P.T., 2006, Astron.Astrophys. 450, 855
//     Hilton, J. et al., 2006, Celest.Mech.Dyn.Astron. 94, 351
//
//  This revision:  2019 July 26
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoFw2m Form rotation matrix given the Fukushima-Williams angles.
func CgoFw2m(gamb, phib, psi, eps float64) (r [3][3]float64) {
	var cR [3][3]C.double
	C.iauFw2m(C.double(gamb), C.double(phib),
		C.double(psi), C.double(eps), &cR[0])
	return v3tC2Go(cR)
}

// GoFw2m Form rotation matrix given the Fukushima-Williams angles.
func GoFw2m(gamb, phib, psi, eps float64) (r [3][3]float64) {
	// Construct the matrix.
	r = GoIr()
	r = GoRz(gamb, r)
	r = GoRx(phib, r)
	r = GoRz(-psi, r)
	r = GoRx(-eps, r)
	return
}
