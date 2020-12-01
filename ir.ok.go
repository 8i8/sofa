package sofa

// #include "sofa.h"
import "C"

//  CgoIr Initialize an r-matrix to the identity matrix.
//
//  - - -
//   I r
//  - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards Of Fundamental Astronomy) software collection.
//
//  Status:  vector/matrix support function.
//
//  Returned:
//     r       double[3][3]    r-matrix
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
//  CgoIr Initialize an r-matrix to the identity matrix.
func CgoIr() (idm [3][3]float64) {
	var cIdm [3][3]C.double
	C.iauIr(&cIdm[0])
	return v3tC2Go(cIdm)
}

//  GoIr Initialize an r-matrix to the identity matrix.
func GoIr() (idmatrix [3][3]float64) {
	idmatrix[0][0] = 1.0
	idmatrix[0][1] = 0.0
	idmatrix[0][2] = 0.0
	idmatrix[1][0] = 0.0
	idmatrix[1][1] = 1.0
	idmatrix[1][2] = 0.0
	idmatrix[2][0] = 0.0
	idmatrix[2][1] = 0.0
	idmatrix[2][2] = 1.0
	return
}
