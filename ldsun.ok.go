package sofa

// #include "sofa.h"
import "C"

//  CgoLdsun Deflection of starlight by the Sun.
//
//  - - - - - -
//   L d s u n
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     p      double[3]  direction from observer to star (unit vector)
//     e      double[3]  direction from Sun to observer (unit vector)
//     em     double     distance from Sun to observer (au)
//
//  Returned:
//     p1     double[3]  observer to deflected star (unit vector)
//
//  Notes:
//
//  1) The source is presumed to be sufficiently distant that its
//     directions seen from the Sun and the observer are essentially
//     the same.
//
//  2) The deflection is restrained when the angle between the star and
//     the center of the Sun is less than a threshold value, falling to
//     zero deflection for zero separation.  The chosen threshold value
//     is within the solar limb for all solar-system applications, and
//     is about 5 arcminutes for the case of a terrestrial observer.
//
//  3) The arguments p and p1 can be the same array.
//
//  Called:
//     iauLd        light deflection by a solar-system body
//
//  This revision:   2016 June 16
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoLdsun Deflection of starlight by the Sun.
func CgoLdsun(p, e [3]float64, em float64) (p1 [3]float64) {
	var cP1 [3]C.double
	cP, cE := v3sGo2C(p), v3sGo2C(e)
	C.iauLdsun(&cP[0], &cE[0], C.double(em), &cP1[0])
	return v3sC2Go(cP1)
}

//  GoLdsun Deflection of starlight by the Sun.
func GoLdsun(p, e [3]float64, em float64) (p1 [3]float64) {

	var em2, dlim float64

	// Deflection limiter (smaller for distant observers).
	em2 = em * em
	if em2 < 1.0 {
		em2 = 1.0
	}

	if em2 > 1.0 {
		dlim = 1e-6 / em2
	} else {
		dlim = 1e-6 / 1.0
	}

	// Apply the deflection.
	p1 = GoLd(1.0, p, p, e, em, dlim)

	return
}
