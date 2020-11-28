package sofa

//  Pdp p-vector inner (=scalar=dot) product.
//
//  - - - -
//   P d p
//  - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Given:
//     a      double[3]     first p-vector
//     b      double[3]     second p-vector
//
//  Returned (function value):
//            double        a . b
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
// double iauPdp(double a[3], double b[3])
func Pdp(a, b [3]float64) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}
