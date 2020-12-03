package sofa

import (
	"flag"
	"fmt"
	"log"
	"math"
	"testing"
)

// verbose shows the output of all tests.
var verbose = flag.Bool("V", false, "show tests output")

func init() {
	log.SetFlags(log.Lshortfile)
}

//
//  - - - -
//   v i v
//  - - - -
//
//  Validate an integer result.
//
//  Internal function used by t_sofa_c program.
//
//  Given:
//     t        *testing.T   go test struct
//     ival     int          value computed by function under test
//     ivalok   int          correct value
//     func     char[]       name of function under test
//     test     char[]       name of individual test
//
//  This revision:  2020 December 3
//
func viv(t *testing.T, ival, ivalok int, fname, test string) {
	if ival != ivalok {
		log.Output(2, fmt.Sprintf(
			"%s failed: %s: want %d got %d",
			fname, test, ival, ivalok))
		t.Fail()
	} else if *verbose {
		log.Output(2, fmt.Sprintf(
			"%s passed: %s: want %d got %d",
			fname, test, ival, ivalok))
	}
}

//
//  - - - -
//   v v d
//  - - - -
//
//  Validate a double result.
//
//  Internal function used by test program.
//
//  Given:
//     t        *testing.T   go test struct
//     val      double       value computed by function under test
//     valok    double       expected value
//     dval     double       maximum allowable error
//     fname    string       name of function under test
//     test     string       name of individual test
//
//  This revision:  2020 December 3
//
func vvd(t *testing.T, val, valok, dval float64, fname, test string) {
	var a, f float64 // Absolute and fractional error.
	a = val - valok
	if a != 0.0 && math.Abs(a) > math.Abs(dval) {
		f = math.Abs(valok / a)
		log.Output(2, fmt.Sprintf(
			"%s failed: %s want %.20f got %.20f (1/%.3f)",
			fname, test, valok, val, f))
		t.Fail()
	} else if *verbose {
		f := math.Abs(valok / a)
		log.Output(2, fmt.Sprintf(
			"%s passed: %s want %.20f got %.20f (1/%.3f)",
			fname, test, valok, val, f))
	}
}

//
//  - - - - -
//   e r r T
//  - - - - -
//
//  Validate an error that has been returned by a library function.
//
//  Internal function used by test program.
//
//  Given:
//     t        *testing.T   go test struct
//     want     error        expected error
//     err      error        error produced by called function
//     tname    string       name of individual test
//
//  This revision:  2020 December 3
//
func errT(t *testing.T, want, err error, tname string) {
	if err != want {
		log.Output(2, fmt.Sprintf("%s failed: want %q got %q",
			tname, want, err))
		t.Fail()
	} else if *verbose {
		log.Output(2, fmt.Sprintf("%s passed: want %q got %q",
			tname, want, err))
	}
}
