package sofa

import (
	"log"
	"testing"
)

//
//  - - - - - - - - - - -
//   T e s t C a l 2 j d
//  - - - - - - - - - - -
//
//  Test Cal2jd function.
//
//  Called:  Cal2jd, vvd, viv
//
//  This revision:  2014 August 7
//
func TestCal2jd(t *testing.T) {
	const fname = "Cal2jd"
	tests := []struct {
		ref string
		fn  func(a, b, c int) (d, e float64, err error)
	}{
		{"cgo", CgoCal2jd},
		{"go", GoCal2jd},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		// Test date 2003, 06, 01
		djm0, djm, err := test.fn(2003, 06, 01)
		if err != nil {
			t.Errorf("%s failed: error %s", tname, err)
		} else if *verbose {
			log.Printf("%s passed: error %s", tname, err)
		}

		// expect 2400000.5, 0.0
		vvd(t, djm0, 2400000.5, 0.0, tname, "djm0")
		// expect 52791.0, 0.0
		vvd(t, djm, 52791.0, 0.0, tname, "djm")
	}
}

func BenchmarkCal2jd(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a, b, c int) (d, e float64, err error)
	}{
		{"cgo", CgoCal2jd},
		{"go", GoCal2jd},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = test.fn(2003, 06, 01)
			}
		})
	}
}
