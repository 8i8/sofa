package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t E r a 0 0
//  - - - - - - - - - -
//
//  Test Era00 function.
//
//  Called:  Era00, vvd
//
//  This revision:  2013 August 7
//
func TestEra00(t *testing.T) {
	const fname = "Era00"
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoEra00},
		{"go", GoEra00},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		era00 := test.fn(2400000.5, 54388.0)
		vvd(t, era00, 0.4022837240028158102, 1e-12, tname, "")
	}
}

func BenchmarkEra00(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoEra00},
		{"go", GoEra00},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 54388.0)
			}
		})
	}
}
