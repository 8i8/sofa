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

	a, err = Af2a('-', 45, 13, 27.2)
	if err != nil {
		t.Errorf("%s failed: %s", fname, err)
	} else if verbose {
		log.Printf("%s passed: %s", fname, err)
	}

	vvd(t, a, -0.7893115794313644842, 1e-12, fname, "a")
}

func TestGoAf2a(t *testing.T) {
	const fname = "Af2a"
	var a float64
	var err error

	a, err = goAf2a('-', 45, 13, 27.2)
	if err != nil {
		t.Errorf("%s failed: %s", fname, err)
	} else if verbose {
		log.Printf("%s passed: %s", fname, err)
	}

	vvd(t, a, -0.7893115794313644842, 1e-12, fname, "a")
}

func BenchmarkAf2a(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Af2a('-', 45, 13, 27.2)
	}
}

func BenchmarkGoAf2a(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = goAf2a('-', 45, 13, 27.2)
	}
}
