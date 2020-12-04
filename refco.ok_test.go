package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t R e f c o
//  - - - - - - - - - -
//
//  Test Refco function.
//
//  Called:  Refco, vvd
//
//  This revision:  2013 October 2
//
func TestRefco(t *testing.T) {
	const fname = "Refco"
	var phpa, tc, rh, wl float64

	phpa = 800.0
	tc = 10.0
	rh = 0.9
	wl = 0.4

	tests := []struct {
		ref string
		fn  func(a1,a2,a3,a4 float64) (a5, a6 float64)
	}{
		{"cgo", CgoRefco},
		{"go", GoRefco},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		refa, refb := test.fn(phpa, tc, rh, wl)

		vvd(t, refa, 0.2264949956241415009e-3, 1e-15,
			tname, "refa")
		vvd(t, refb, -0.2598658261729343970e-6, 1e-18,
			tname, "refb")
	}
}

func BenchmarkRefco(b *testing.B) {
	var phpa, tc, rh, wl float64

	phpa = 800.0
	tc = 10.0
	rh = 0.9
	wl = 0.4

	tests := []struct {
		ref string
		fn  func(a1,a2,a3,a4 float64) (a5, a6 float64)
	}{
		{"cgo", CgoRefco},
		{"go", GoRefco},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(phpa, tc, rh, wl)
			}
		})
	}
}
