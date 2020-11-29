package sofa

import (
	"log"
	"testing"
)

//
//  - - - - - - - - -
//   T e s t A 2 a f
//  - - - - - - - - -
//
//  Test A2af function.
//
//  Called:  A2af, viv
//
//  This revision:  2013 August 7
//
func TestA2af(t *testing.T) {
	const fname = "A2af"
	tests := []struct {
		ref string
		fn  func(int, float64) (byte, [4]int)
	}{
		{"cgo", A2af},
		{"go", goA2af},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		sign, idmsf := test.fn(4, 2.345)

		if sign != '+' {
			t.Errorf("%s failed: want \"+\" got %q",
				tname, sign)
		} else if verbose {
			log.Printf("%s passed: want \"+\" got %q",
				tname, sign)
		}
		viv(t, idmsf[0], 134, tname, "0")
		viv(t, idmsf[1], 21, tname, "1")
		viv(t, idmsf[2], 30, tname, "2")
		viv(t, idmsf[3], 9706, tname, "3")
	}
}

func BenchmarkA2af(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(int, float64) (byte, [4]int)
	}{
		{"cgo", A2af},
		{"go", goA2af},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(4, 2.345)
			}
		})
	}
}
