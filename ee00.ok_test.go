package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t E e 0 0
//  - - - - - - - - -
//
//  Test Ee00 function.
//
//  Called:  Ee00, vvd
//
//  This revision:  2013 August 7
//
func TestEe00(t *testing.T) {
	const fname = "Ee00"
	var epsa, dpsi, ee float64

	epsa = 0.4090789763356509900
	dpsi = -0.9630909107115582393e-5

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) float64
	}{
		{"cgo", CgoEe00},
		{"go", GoEe00},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		ee = test.fn(2400000.5, 53736.0, epsa, dpsi)

		vvd(t, ee, -0.8834193235367965479e-5, 1e-18, tname, "")
	}
}

func BenchmarkEe00(b *testing.B) {
	var epsa, dpsi float64

	epsa = 0.4090789763356509900
	dpsi = -0.9630909107115582393e-5

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) float64
	}{
		{"cgo", CgoEe00},
		{"go", GoEe00},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736.0, epsa, dpsi)
			}
		})
	}
}
