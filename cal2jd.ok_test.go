package sofa

import (
	"log"
	"testing"
)

/*
**  - - - - - - - - -
**   t _ c a l 2 j d
**  - - - - - - - - -
**
**  Test iauCal2jd function.
**
**  Returned:
**     status    int         FALSE = success, TRUE = fail
**
**  Called:  iauCal2jd, vvd, viv
**
**  This revision:  2014 August 7
 */
func TestCal2jd(t *testing.T) {
	const fname = "Cal2jd"

	// Test date 2003, 06, 01
	djm0, djm, err := Cal2jd(2003, 06, 01)
	if err != nil {
		t.Errorf("%s failed: error %s", fname, err)
	} else if verbose {
		log.Printf("%s passed: error %s", fname, err)
	}

	// expect 2400000.5, 0.0
	vvd(t, djm0, 2400000.5, 0.0, fname, "djm0")
	// expect 52791.0, 0.0
	vvd(t, djm, 52791.0, 0.0, fname, "djm")
}
