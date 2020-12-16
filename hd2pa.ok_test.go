package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t H d 2 p a
//  - - - - - - - - - -
//
//  Test Hd2pa function.
//
//  Called:  Hd2pa and vvd
//
//  This revision:  2017 October 21
//
func TestHd2pa(t *testing.T) {
	const fname = "Hd2pa"
	var h, d, p, q float64

	h = 1.1
	d = 1.2
	p = 0.3

	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) float64
	}{
		{"cgo", CgoHd2pa},
		{"go", GoHd2pa},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		q = test.fn(h, d, p)

		vvd(t, q, 1.906227428001995580, 1e-13, tname, "q")
	}
}

func BenchmarkHd2pa(b *testing.B) {
	var h, d, p float64

	h = 1.1
	d = 1.2
	p = 0.3

	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) float64
	}{
		{"cgo", CgoHd2pa},
		{"go", GoHd2pa},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(h, d, p)
			}
		})
	}
}
