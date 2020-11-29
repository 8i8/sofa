package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t A n p m
//  - - - - - - - - -
//
//  Test Anpm function.
//
//  Called:  Anpm, vvd
//
//  This revision:  2013 August 7
//
func TestAnpm(t *testing.T) {
	const fname = "Anpm"
	vvd(t, Anpm(-4.0), 2.283185307179586477, 1e-12, fname, "")
}

func TestGoAnpm(t *testing.T) {
	const fname = "Anpm"
	vvd(t, goAnpm(-4.0), 2.283185307179586477, 1e-12, fname, "")
}
