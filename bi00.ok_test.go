package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t B i 0 0
//  - - - - - - - - -
//
//  Test Bi00 function.
//
//  Called:  iauBi00, vvd
//
//  This revision:  2013 August 7
//
func TestBi00(t *testing.T) {
	const fname = "Bi00"
	var dpsibi, depsbi, dra float64

	a, b, c := Bi00(dpsibi, depsbi, dra)

	vvd(t, a, -0.2025309152835086613e-6, 1e-12, fname, "dpsibi")
	vvd(t, b, -0.3306041454222147847e-7, 1e-12, fname, "depsbi")
	vvd(t, c, -0.7078279744199225506e-7, 1e-12, fname, "dra")
}
