package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t E e c t 0 0
//  - - - - - - - - - - -
//
//  Test Eect00 function.
//
//  Called:  Eect00, vvd
//
//  This revision:  2013 August 7
//
func TestEect00(t *testing.T) {
	const fname = "Eect00"
	var eect float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoEect00},
		{"go", GoEect00},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		eect = test.fn(2400000.5, 53736.0)

		vvd(t, eect, 0.2046085004885125264e-8, 1e-20, tname, "")
	}
}

func BenchmarkEect00(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoEect00},
		{"go", GoEect00},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736.0)
			}
		})
	}
}
