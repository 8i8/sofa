package sofa

import (
	"fmt"
	"log"
	"math"
	"testing"
)

var verbose = false

func init() {
	log.SetFlags(log.Lshortfile)
}

/*
**  - - - -
**   v i v
**  - - - -
**
**  Validate an integer result.
**
**  Internal function used by t_sofa_c program.
**
**  Given:
**     ival     int          value computed by function under test
**     ivalok   int          correct value
**     func     char[]       name of function under test
**     test     char[]       name of individual test
**
**  Given and returned:
**     status   int          set to TRUE if test fails
**
**  This revision:  2013 August 7
**
 */
func viv(t *testing.T, ival, ivalok int, fname, test string) {
	if ival != ivalok {
		log.Output(2, fmt.Sprintf(
			"%s failed: %s: want %d got %d",
			fname, test, ival, ivalok))
		t.Fail()
	} else if verbose {
		log.Output(2, fmt.Sprintf(
			"%s passed: %s: want %d got %d",
			fname, test, ival, ivalok))
	}
}

/*
**  - - - -
**   v v d
**  - - - -
**
**  Validate a double result.
**
**  Internal function used by test program.
**
**  Given:
**     val      double       value computed by function under test
**     valok    double       expected value
**     dval     double       maximum allowable error
**     fname    string       name of function under test
**     test     string       name of individual test
**
**  Given and returned:
**     status   int          set to TRUE if test fails
**
**  This revision:  2016 April 21
 */
func vvd(t *testing.T, val, valok, dval float64, fname, test string) {
	a := val - valok
	if a != 0.0 && math.Abs(a) > math.Abs(dval) {
		f := math.Abs(valok / a)
		log.Output(2, fmt.Sprintf(
			"%s failed: %s want %.20f got %.20f (1/%.3f)",
			fname, test, valok, val, f))
		t.Fail()
	} else if verbose {
		f := math.Abs(valok / a)
		log.Output(2, fmt.Sprintf(
			"%s passed: %s want %.20f got %.20f (1/%.3f)",
			fname, test, valok, val, f))
	}
}
