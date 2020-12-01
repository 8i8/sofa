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
	const fname = "A2tf"
	var ihmsf [4]int
	var s byte
	tests := []struct {
		ref string
		fn  func(int, float64) (byte, [4]int)
	}{
		{"cgo", CgoA2tf},
		{"go", GoA2tf},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		s, ihmsf = test.fn(4, -3.01234)

		viv(t, int(s), '-', tname, "s")
		viv(t, ihmsf[0], 11, tname, "0")
		viv(t, ihmsf[1], 30, tname, "1")
		viv(t, ihmsf[2], 22, tname, "2")
		viv(t, ihmsf[3], 6484, tname, "3")
	}
}

func BenchmarkA2tf(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(int, float64) (byte, [4]int)
	}{
		{"cgo", CgoA2tf},
		{"go", GoA2tf},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(4, -3.01234)
			}
		})
	}
}
