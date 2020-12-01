package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t B i 0 0
//  - - - - - - - - -
//
//  Test Bi00 function.
//
//  Called:  iauBi00, vvd
//
//  This revision:  2013 August 7
//
func TestBi00(t *testing.T) {
	const fname = "Bi00"
	var dpsibi, depsbi, dra float64

	tests := []struct {
		ref string
		fn  func(a, b, c float64) (d, e, f float64)
	}{
		{"cgo", CgoBi00},
		{"go", GoBi00},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		a, b, c := test.fn(dpsibi, depsbi, dra)

		vvd(t, a, -0.2025309152835086613e-6, 1e-12, tname, "dpsibi")
		vvd(t, b, -0.3306041454222147847e-7, 1e-12, tname, "depsbi")
		vvd(t, c, -0.7078279744199225506e-7, 1e-12, tname, "dra")
	}
}

func BenchmarkBi00(b *testing.B) {
	var dpsibi, depsbi, dra float64

	tests := []struct {
		ref string
		fn  func(a, b, c float64) (d, e, f float64)
	}{
		{"cgo", CgoBi00},
		{"go", GoBi00},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = test.fn(dpsibi, depsbi, dra)
			}
		})
	}
}
