package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t E e 0 0 b
//  - - - - - - - - - -
//
//  Test Ee00b function.
//
//  Called:  Ee00b, vvd
//
//  This revision:  2013 August 7
//
func TestEe00b(t *testing.T) {
	const fname = "Ee00b"
	var ee float64

	tests := []struct {
		ref string
		fn  func(a1,a2 float64) float64	
	}{
		{"cgo", CgoEe00b},
		{"go", GoEe00b},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		ee = test.fn(2400000.5, 53736.0)

		vvd(t, ee, -0.8835700060003032831e-5, 1e-18, tname, "")
	}
}

func BenchmarkEe00b(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1,a2 float64) float64	
	}{
		{"cgo", CgoEe00b},
		{"go", GoEe00b},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736.0)
			}
		})
	}
}
