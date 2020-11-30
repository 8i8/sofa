package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t P r 0 0
//  - - - - - - - - -
//
//  Test Pr00 function.
//
//  Called:  iauPr00, vvd
//
//  This revision:  2013 August 7
//
func TestPr00(t *testing.T) {
	const fname = "Pr00"
	var dpsipr, depspr float64
	dpsipr, depspr = Pr00(2400000.5, 53736)

	vvd(t, dpsipr, -0.8716465172668347629e-7, 1e-22,
		fname, "dpsipr")
	vvd(t, depspr, -0.7342018386722813087e-8, 1e-22,
		fname, "depspr")
}
