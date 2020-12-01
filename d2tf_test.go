package sofa

import "testing"

//
//  - - - - - - -
//   t _ d 2 t f
//  - - - - - - -
//
//  Test D2tf function.
//
//  Returned:
//     status    int         FALSE = success, TRUE = fail
//
//  Called:  D2tf, viv
//
//  This revision:  2013 August 7
//
func TestD2tf(t *testing.T) {
	const fname = "D2tf"

	tests := []struct {
		ref string
		fn  func(int, float64) (byte, [4]int)
	}{
		{"cgo", CgoD2tf},
		{"go", GoD2tf},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		sign, idmsf := test.fn(4, -0.987654321)
		want := byte('-')
		if sign != want {
			t.Errorf("%s: expected %q received %q",
				tname, want, sign)
		}
		viv(t, idmsf[0], 23, tname, "0")
		viv(t, idmsf[1], 42, tname, "1")
		viv(t, idmsf[2], 13, tname, "2")
		viv(t, idmsf[3], 3333, tname, "3")
	}
}

func BenchmarkD2tf(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(int, float64) (byte, [4]int)
	}{
		{"cgo", CgoD2tf},
		{"go", GoD2tf},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(4, -0.987654321)
			}
		})
	}
}
