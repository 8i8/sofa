package sofa

import "testing"

//
//  - - - - - -
//   t _ a n p
//  - - - - - -
//
//  Test iauAnp function.
//
//  Returned:
//     status    int         FALSE = success, TRUE = fail
//
//  Called:  iauAnp, vvd
//
//  This revision:  2013 August 7
//
func TestAnp(t *testing.T) {
	const fname = "Anp"
	vvd(t, Anp(-0.1), 6.183185307179586477, 1e-12, fname, "")
}

func TestGoAnp(t *testing.T) {
	const fname = "Anp"
	vvd(t, goAnp(-0.1), 6.183185307179586477, 1e-12, fname, "")
}

func BenchmarkAnp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Anp(-0.1)
	}
}

func BenchmarkGoAnp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = goAnp(-0.1)
	}
}
