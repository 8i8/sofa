package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t X y s 0 6 a
//  - - - - - - - - - - -
//
//  Test Xys06a function.
//
//  Called:  Xys06a, vvd
//
//  This revision:  2013 August 7
//
func TestXys06a(t *testing.T) {
	const fname = "Xys06a"
	var x, y, s float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2, c3 float64)
	}{
		{"cgo", CgoXys06a},
		{"go", GoXys06a},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		x, y, s = test.fn(2400000.5, 53736.0)

		vvd(t, x, 0.5791308482835292617e-3, 1e-14, tname, "x")
		vvd(t, y, 0.4020580099454020310e-4, 1e-15, tname, "y")
		vvd(t, s, -0.1220032294164579896e-7, 1e-18, tname, "s")
	}
}

func BenchmarkXys06a(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2, c3 float64)
	}{
		{"cgo", CgoXys06a},
		{"go", GoXys06a},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = test.fn(2400000.5, 53736.0)
			}
		})
	}
}
