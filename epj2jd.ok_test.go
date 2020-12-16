package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t E p j 2 j d
//  - - - - - - - - -
//
//  Test Epj2jd function.
//
//  Called:  Epj2jd, vvd
//
//  This revision:  2013 August 7
//
func TestEpj2jd(t *testing.T) {
	const fname = "Epj2jd"
	var epj, djm0, djm float64

	epj = 1996.8

	tests := []struct {
		ref string
		fn  func(float64) (c1, c2 float64)
	}{
		{"cgo", CgoEpj2jd},
		{"go", GoEpj2jd},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		djm0, djm = test.fn(epj)

		vvd(t, djm0, 2400000.5, 1e-9, tname, "djm0")
		vvd(t, djm, 50375.7, 1e-9, tname, "mjd")
	}
}

func BenchmarkEpj2jd(b *testing.B) {
	var epj float64

	epj = 1996.8

	tests := []struct {
		ref string
		fn  func(float64) (c1, c2 float64)
	}{
		{"cgo", CgoEpj2jd},
		{"go", GoEpj2jd},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(epj)
			}
		})
	}
}
