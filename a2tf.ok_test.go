package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t A 2 t f
//  - - - - - - - - -
//
//  Test A2tf function.
//
//  Called:  A2tf, viv
//
//  This revision:  2013 August 7
//
func TestA2tf(t *testing.T) {
	var ihmsf [4]int
	var s byte

	s, ihmsf = A2tf(4, -3.01234)

	viv(t, int(s), '-', "iauA2tf", "s")

	viv(t, ihmsf[0], 11, "iauA2tf", "0")
	viv(t, ihmsf[1], 30, "iauA2tf", "1")
	viv(t, ihmsf[2], 22, "iauA2tf", "2")
	viv(t, ihmsf[3], 6484, "iauA2tf", "3")
}

func TestGoA2tf(t *testing.T) {
	var ihmsf [4]int
	var s byte

	s, ihmsf = goA2tf(4, -3.01234)

	viv(t, int(s), '-', "iauA2tf", "s")

	viv(t, ihmsf[0], 11, "iauA2tf", "0")
	viv(t, ihmsf[1], 30, "iauA2tf", "1")
	viv(t, ihmsf[2], 22, "iauA2tf", "2")
	viv(t, ihmsf[3], 6484, "iauA2tf", "3")
}

func BenchmarkA2tf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = A2tf(4, -3.01234)
	}
}

func BenchmarkGoA2tf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = goA2tf(4, -3.01234)
	}
}
