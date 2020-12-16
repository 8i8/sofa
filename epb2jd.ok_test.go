package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t E p b 2 j d
//  - - - - - - - - - - -
//
//  Test Epb2jd function.
//
//  Called:  Epb2jd, vvd
//
//  This revision:  2013 August 7
//
func TestEpb2jd(t *testing.T) {
	const fname = "Epb2jd"
	var epb, djm0, djm float64

	epb = 1957.3

	tests := []struct {
		ref string
		fn  func(float64) (c1, c2 float64)
	}{
		{"cgo", CgoEpb2jd},
		{"go", GoEpb2jd},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		djm0, djm = test.fn(epb)

		vvd(t, djm0, 2400000.5, 1e-9, tname, "djm0")
		vvd(t, djm, 35948.1915101513, 1e-9, tname, "mjd")
	}
}

func BenchmarkEpb2jd(b *testing.B) {
	var epb float64

	epb = 1957.3

	tests := []struct {
		ref string
		fn  func(float64) (c1, c2 float64)
	}{
		{"cgo", CgoEpb2jd},
		{"go", GoEpb2jd},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(epb)
			}
		})
	}
}
