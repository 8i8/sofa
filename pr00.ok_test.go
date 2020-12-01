package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t P r 0 0
//  - - - - - - - - -
//
//  Test Pr00 function.
//
//  Called:  iauPr00, vvd
//
//  This revision:  2013 August 7
//
func TestPr00(t *testing.T) {
	const fname = "Pr00"
	var dpsipr, depspr float64
	tests := []struct {
		ref string
		fn  func(d1, d2 float64) (d3, d4 float64)
	}{
		{"cgo", CgoPr00},
		{"go", GoPr00},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		dpsipr, depspr = GoPr00(2400000.5, 53736)

		vvd(t, dpsipr, -0.8716465172668347629e-7, 1e-22,
			tname, "dpsipr")
		vvd(t, depspr, -0.7342018386722813087e-8, 1e-22,
			tname, "depspr")
	}
}

func BenchmarkPr00(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(d1, d2 float64) (d3, d4 float64)
	}{
		{"cgo", CgoPr00},
		{"go", GoPr00},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(2400000.5, 53736)
			}
		})
	}
}
