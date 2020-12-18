package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t X y 0 6
//  - - - - - - - - -
//
//  Test Xy06 function.
//
//  Called:  Xy06, vvd
//
//  This revision:  2013 August 7
//
func TestXy06(t *testing.T) {
	const fname = "Xy06"
	var x, y float64

	tests := []struct {
		ref string
		fn  func(a1,a2 float64) (c1,c2 float64) 
	}{
		{"cgo", CgoXy06},
		{"go", GoXy06},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		x, y = test.fn(2400000.5, 53736.0)

		vvd(t, x, 0.5791308486706010975e-3, 1e-15, tname, "x")
		vvd(t, y, 0.4020579816732958141e-4, 1e-16, tname, "y")
	}
}

func BenchmarkXy06(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1,a2 float64) (c1,c2 float64) 
	}{
		{"cgo", CgoXy06},
		{"go", GoXy06},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(2400000.5, 53736.0)
			}
		})
	}
}
