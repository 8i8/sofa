package sofa

import (
	"log"
	"testing"
)

//
//  - - - - - - - - -
//   T e s t A f 2 a
//  - - - - - - - - -
//
//  Test Af2a function.
//
//  Called:  iauAf2a, viv
//
//  This revision:  2013 August 7
//
func TestAf2a(t *testing.T) {
	const fname = "Af2a"
	var a float64
	var err error

	tests := []struct {
		ref string
		fn  func(byte, int, int, float64) (float64, error)
	}{
		{"cgo", Af2a},
		{"go", goAf2a},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		a, err = test.fn('-', 45, 13, 27.2)
		if err != nil {
			t.Errorf("%s failed: %s", tname, err)
		} else if *verbose {
			log.Printf("%s passed: %s", tname, err)
		}

		vvd(t, a, -0.7893115794313644842, 1e-12, tname, "a")
	}
}

func BenchmarkAf2a(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(byte, int, int, float64) (float64, error)
	}{
		{"cgo", Af2a},
		{"go", goAf2a},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn('-', 45, 13, 27.2)
			}
		})
	}
}
