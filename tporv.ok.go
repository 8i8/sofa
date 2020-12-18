package sofa

// #include "sofa.h"
import "C"
import (
	"math"

	"github.com/8i8/sofa/en"
)

var errTporv = en.New(0, "Tporv", []string{
	"no solutions returned (Note 4)",
	"only the first solution is useful (Note 5)",
	"both solutions are useful (Note 5)",
})

//  CgoTporv In the tangent plane projection, given the rectangular
//  coordinates of a star and its direction cosines, determine the
//  direction cosines of the tangent point.
//
//  - - - - - -
//   T p o r v
//  - - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     xi,eta   double    rectangular coordinates of star image (Note 2)
//     v        double[3] star's direction cosines (Note 3)
//
//  Returned:
//     v01      double[3] tangent point's direction cosines, Solution 1
//     v02      double[3] tangent point's direction cosines, Solution 2
//
//  Returned (function value):
//                int     number of solutions:
//                        0 = no solutions returned (Note 4)
//                        1 = only the first solution is useful (Note 5)
//                        2 = both solutions are useful (Note 5)
//
//  Notes:
//
//  1) The tangent plane projection is also called the "gnomonic
//     projection" and the "central projection".
//
//  2) The eta axis points due north in the adopted coordinate system.
//     If the direction cosines represent observed (RA,Dec), the tangent
//     plane coordinates (xi,eta) are conventionally called the
//     "standard coordinates".  If the direction cosines are with
//     respect to a right-handed triad, (xi,eta) are also right-handed.
//     The units of (xi,eta) are, effectively, radians at the tangent
//     point.
//
//  3) The vector v must be of unit length or the result will be wrong.
//
//  4) Cases where there is no solution can arise only near the poles.
//     For example, it is clearly impossible for a star at the pole
//     itself to have a non-zero xi value, and hence it is meaningless
//     to ask where the tangent point would have to be.
//
//  5) Also near the poles, cases can arise where there are two useful
//     solutions.  The return value indicates whether the second of the
//     two solutions returned is useful;  1 indicates only one useful
//     solution, the usual case.
//
//  6) The basis of the algorithm is to solve the spherical triangle
//     PSC, where P is the north celestial pole, S is the star and C is
//     the tangent point.  Calling the celestial spherical coordinates
//     of the star and tangent point (a,b) and (a0,b0) respectively, and
//     writing rho^2 = (xi^2+eta^2) and r^2 = (1+rho^2), and
//     transforming the vector v into (a,b) in the normal way, side c is
//     then (pi/2-b), side p is sqrt(xi^2+eta^2) and side s (to be
//     found) is (pi/2-b0), while angle C is given by sin(C) = xi/rho
//     and cos(C) = eta/rho;  angle P (to be found) is (a-a0).  After
//     solving the spherical triangle, the result (a0,b0) can be
//     expressed in vector form as v0.
//
//  7) This function is a member of the following set:
//
//         spherical      vector         solve for
//
//         iauTpxes      iauTpxev         xi,eta
//         iauTpsts      iauTpstv          star
//         iauTpors    > iauTporv <       origin
//
//  References:
//
//     Calabretta M.R. & Greisen, E.W., 2002, "Representations of
//     celestial coordinates in FITS", Astron.Astrophys. 395, 1077
//
//     Green, R.M., "Spherical Astronomy", Cambridge University Press,
//     1987, Chapter 13.
//
//  This revision:   2018 January 2
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoTporv In the tangent plane projection, given the rectangular
//  coordinates of a star and its direction cosines, determine the
//  direction cosines of the tangent point.
func CgoTporv(xi, eta float64, v [3]float64) (
	v01, v02 [3]float64, err en.ErrNum) {
	var cV01, cV02 [3]C.double
	cV := v3sGo2C(v)
	cI := C.iauTporv(C.double(xi), C.double(eta), &cV[0], &cV01[0],
		&cV02[0])
	err = errTporv.Set(int(cI)) // A value is always returned.
	return v3sC2Go(cV01), v3sC2Go(cV02), err
}

//  GoTporv In the tangent plane projection, given the rectangular
//  coordinates of a star and its direction cosines, determine the
//  direction cosines of the tangent point.
func GoTporv(xi, eta float64, v [3]float64) (
	v01, v02 [3]float64, err en.ErrNum) {
	var x, y, z, rxy2, xi2, eta2p1, r, rsb, rcb, w2, w, c float64

	x = v[0]
	y = v[1]
	z = v[2]
	rxy2 = x*x + y*y
	xi2 = xi * xi
	eta2p1 = eta*eta + 1.0
	r = math.Sqrt(xi2 + eta2p1)
	rsb = r * z
	rcb = r * math.Sqrt(x*x+y*y)
	w2 = rcb*rcb - xi2
	if w2 > 0.0 {
		w = math.Sqrt(w2)
		c = (rsb*eta + w) / (eta2p1 * math.Sqrt(rxy2*(w2+xi2)))
		v01[0] = c * (x*w + y*xi)
		v01[1] = c * (y*w - x*xi)
		v01[2] = (rsb - eta*w) / eta2p1
		w = -w
		c = (rsb*eta + w) / (eta2p1 * math.Sqrt(rxy2*(w2+xi2)))
		v02[0] = c * (x*w + y*xi)
		v02[1] = c * (y*w - x*xi)
		v02[2] = (rsb - eta*w) / eta2p1
		if math.Abs(rsb) < 1.0 {
			err = errTporv.Set(1)
		} else {
			err = errTporv.Set(2)
		}
	} else {
		err = errTporv.Set(0)
	}
	return // An error is always raised.
}
