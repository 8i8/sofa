package sofa

import "testing"

//
//  - - - - - - -
//   t _ d 2 t f
//  - - - - - - -
//
//  Test D2tf function.
//
//  Returned:
//     status    int         FALSE = success, TRUE = fail
//
//  Called:  D2tf, viv
//
//  This revision:  2013 August 7
//
func TestD2tf(t *testing.T) {
	const fname = "D2tf"
	sign, idmsf := D2tf(4, -0.987654321)
	require := byte('-')
	if sign != require {
		t.Errorf("%s: expected %q received %q", fname, require, sign)
	}
	viv(t, idmsf[0], 23, fname, "0")
	viv(t, idmsf[1], 42, fname, "1")
	viv(t, idmsf[2], 13, fname, "2")
	viv(t, idmsf[3], 3333, fname, "3")
}
