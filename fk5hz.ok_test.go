package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t F k 5 h z
//  - - - - - - - - - -
//
//  Test Fk5hz function.
//
//  Called:  Fk5hz, vvd
//
//  This revision:  2013 August 7
//
func TestFk5hz(t *testing.T) {
	const fname = "Fk5hz"
	var r5, d5, rh, dh float64

	r5 = 1.76779433
	d5 = -0.2917517103

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (c1, c2 float64)
	}{
		{"cgo", CgoFk5hz},
		{"go", GoFk5hz},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rh, dh = test.fn(r5, d5, 2400000.5, 54479.0)

		vvd(t, rh, 1.767794191464423978, 1e-12, tname, "ra")
		vvd(t, dh, -0.2917516001679884419, 1e-12, tname, "dec")
	}
}

func BenchmarkFk5hz(b *testing.B) {
	var r5, d5 float64

	r5 = 1.76779433
	d5 = -0.2917517103

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (c1, c2 float64)
	}{
		{"cgo", CgoFk5hz},
		{"go", GoFk5hz},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(r5, d5, 2400000.5, 54479.0)
			}
		})
	}
}
