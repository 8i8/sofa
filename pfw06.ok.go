package sofa

// #include "sofa.h"
import "C"

//  CgoPfw06 Precession angles, IAU 2006 (Fukushima-Williams 4-angle
//  formulation).
//
//  - - - - - -
//   P f w 0 6
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  canonical model.
//
//  Given:
//     date1,date2  double   TT as a 2-part Julian Date (Note 1)
//
//  Returned:
//     gamb         double   F-W angle gamma_bar (radians)
//     phib         double   F-W angle phi_bar (radians)
//     psib         double   F-W angle psi_bar (radians)
//     epsa         double   F-W angle epsilon_A (radians)
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
//  2) Naming the following points:
//
//           e = J2000.0 ecliptic pole,
//           p = GCRS pole,
//           E = mean ecliptic pole of date,
//     and   P = mean pole of date,
//
//     the four Fukushima-Williams angles are as follows:
//
//        gamb = gamma_bar = epE
//        phib = phi_bar = pE
//        psib = psi_bar = pEP
//        epsa = epsilon_A = EP
//
//  3) The matrix representing the combined effects of frame bias and
//     precession is:
//
//        PxB = R_1(-epsa).R_3(-psib).R_1(phib).R_3(gamb)
//
//  4) The matrix representing the combined effects of frame bias,
//     precession and nutation is simply:
//
//        NxPxB = R_1(-epsa-dE).R_3(-psib-dP).R_1(phib).R_3(gamb)
//
//     where dP and dE are the nutation components with respect to the
//     ecliptic of date.
//
//  Reference:
//
//     Hilton, J. et al., 2006, Celest.Mech.Dyn.Astron. 94, 351
//
//  Called:
//     iauObl06     mean obliquity, IAU 2006
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoPfw06 Precession angles, IAU 2006 (Fukushima-Williams 4-angle
//  formulation).
// void iauPfw06(double date1, double date2,
//               double *gamb, double *phib, double *psib, double *epsa)
func CgoPfw06(date1, date2 float64) (gamb, phib, psib, epsa float64) {
	var cF1, cF2, cF3, cF4 C.double
	C.iauPfw06(C.double(date1), C.double(date2),
		&cF1, &cF2, &cF3, &cF4)
	return float64(cF1), float64(cF2), float64(cF3), float64(cF4)
}

func GoPfw06(date1, date2 float64) (gamb, phib, psib, epsa float64) {
	var t float64

	// Interval between fundamental date J2000.0 and given date
	// (JC).
	t = ((date1 - DJ00) + date2) / DJC

	// P03 bias+precession angles.
	gamb = (-0.052928 +
		(10.556378+(0.4932044+
			(-0.00031238+(-0.000002788+
				(0.0000000260)*
					t)*t)*t)*t)*t) * DAS2R
	phib = (84381.412819 +
		(-46.811016+(0.0511268+
			(0.00053289+(-0.000000440+
				(-0.0000000176)*
					t)*t)*t)*t)*t) * DAS2R
	psib = (-0.041775 +
		(5038.481484+(1.5584175+
			(-0.00018522+(-0.000026452+
				(-0.0000000148)*
					t)*t)*t)*t)*t) * DAS2R
	epsa = GoObl06(date1, date2)

	return
}
