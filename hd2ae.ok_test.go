package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t H d 2 a e
//  - - - - - - - - - -
//
//  Test Hd2ae function.
//
//  Called:  Hd2ae and vvd
//
//  This revision:  2017 October 21
//
func TestHd2ae(t *testing.T) {
	const fname = "Hd2ae"
	var h, d, p, a, e float64

	h = 1.1
	d = 1.2
	p = 0.3

	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (c1, c2 float64)
	}{
		{"cgo", CgoHd2ae},
		{"go", GoHd2ae},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		a, e = test.fn(h, d, p)

		vvd(t, a, 5.916889243730066194, 1e-13, tname, "a")
		vvd(t, e, 0.4472186304990486228, 1e-14, tname, "e")
	}
}

func BenchmarkHd2ae(b *testing.B) {
	var h, d, p float64

	h = 1.1
	d = 1.2
	p = 0.3

	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (c1, c2 float64)
	}{
		{"cgo", CgoHd2ae},
		{"go", GoHd2ae},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(h, d, p)
			}
		})
	}
}
