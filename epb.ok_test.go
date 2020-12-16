package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t E p b
//  - - - - - - - -
//
//  Test Epb function.
//
//  Called:  Epb, vvd
//
//  This revision:  2013 August 7
//
func TestEpb(t *testing.T) {
	const fname = "Epb"
	var epb float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoEpb},
		{"go", GoEpb},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref

		epb = test.fn(2415019.8135, 30103.18648)

		vvd(t, epb, 1982.418424159278580, 1e-12, tname, "")
	}
}

func BenchmarkEpb(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoEpb},
		{"go", GoEpb},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2415019.8135, 30103.18648)
			}
		})
	}
}
